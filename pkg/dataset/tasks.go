package dataset

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/metrics"
)

// Example is one benchmark question in a model-ready form.
type Example struct {
	ID      string   `json:"id"`
	Task    string   `json:"task"`    // dataset name
	Kind    string   `json:"kind"`    // number | choice | math | span | summary | none
	Context string   `json:"-"`       // long shared part of the prompt (documents); goes first so it can be prefix-cached
	Prompt  string   `json:"-"`       // the question with instructions
	Answers []string `json:"answers"` // acceptable gold answers
}

// Text is the full user message.
func (e Example) Text() string {
	if e.Context == "" {
		return e.Prompt
	}
	return e.Context + "\n\n" + e.Prompt
}

var boxed = regexp.MustCompile(`\\boxed\s*\{`)

// Boxed returns the content of the last \boxed{...} ("" if none).
func Boxed(s string) string {
	loc := boxed.FindAllStringIndex(s, -1)
	if len(loc) == 0 {
		return ""
	}
	start, depth := loc[len(loc)-1][1], 1
	for i := start; i < len(s); i++ {
		switch s[i] {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return s[start:i]
			}
		}
	}
	return ""
}

var mathNoise = strings.NewReplacer(" ", "", "\\!", "", "\\,", "", "\\left", "", "\\right", "", "dfrac", "frac", "tfrac", "frac",
	"\\$", "", "$", "", "^\\circ", "", "^{\\circ}", "", "\\%", "", "%", "", "\\text{", "{", "\\mbox{", "{")

// MathEqual compares two LaTeX answers after light normalization.
func MathEqual(a, b string) bool {
	clean := func(s string) string { return strings.TrimSuffix(mathNoise.Replace(strings.TrimSpace(s)), ".") }
	return clean(a) != "" && clean(a) == clean(b)
}

// Extract pulls the short final answer out of the model output (used to compare two outputs).
func Extract(e Example, output string) string {
	switch e.Kind {
	case "number":
		return metrics.LastNumber(output)
	case "choice":
		return metrics.Choice(output)
	case "math":
		return mathNoise.Replace(Boxed(output))
	}
	return metrics.Normalize(output)
}

// Grade scores the model output for the example in [0,1].
func Grade(e Example, output string) float64 {
	hit := func(ok bool) float64 {
		if ok {
			return 1
		}
		return 0
	}
	switch e.Kind {
	case "number":
		return hit(metrics.LastNumber(output) == metrics.LastNumber(e.Answers[0]))
	case "choice":
		return hit(metrics.Choice(output) == e.Answers[0])
	case "math":
		return hit(MathEqual(Boxed(output), e.Answers[0]))
	case "span":
		return metrics.BestF1(output, e.Answers)
	case "summary":
		return metrics.RougeL(output, e.Answers[0])
	}
	return 0
}

const letters = "ABCDEFGHIJ"

func options(choices []string) string {
	var b strings.Builder
	for i, c := range choices {
		fmt.Fprintf(&b, "%c. %s\n", letters[i], c)
	}
	return b.String()
}

// GSM8K: grade-school math word problems, numeric answer.
func GSM8K() ([]Example, error) {
	rows, err := Parquet[struct {
		Question string `parquet:"question"`
		Answer   string `parquet:"answer"`
	}](Root + "openai/gsm8k/main/test-00000-of-00001.parquet")
	out := make([]Example, len(rows))
	for i, r := range rows {
		_, gold, _ := strings.Cut(r.Answer, "####")
		out[i] = Example{ID: fmt.Sprint("gsm8k-", i), Task: "gsm8k", Kind: "number", Answers: []string{strings.TrimSpace(gold)},
			Prompt: r.Question + "\n\nSolve the problem. Finish with a line of the form \"Answer: <number>\"."}
	}
	return out, err
}

// MATH500: competition math, LaTeX answer in \boxed{}.
func MATH500() ([]Example, error) {
	rows, err := JSONL[struct {
		Problem string `json:"problem"`
		Answer  string `json:"answer"`
		ID      string `json:"unique_id"`
	}](Root + "HuggingFaceH4/MATH-500/test.jsonl")
	out := make([]Example, len(rows))
	for i, r := range rows {
		out[i] = Example{ID: "math500-" + r.ID, Task: "math500", Kind: "math", Answers: []string{r.Answer},
			Prompt: r.Problem + "\n\nSolve the problem. Put the final answer in \\boxed{}."}
	}
	return out, err
}

// MMLU: 4-way multiple choice over 57 subjects.
func MMLU() ([]Example, error) {
	rows, err := Parquet[struct {
		Question string   `parquet:"question"`
		Subject  string   `parquet:"subject"`
		Choices  []string `parquet:"choices,list"`
		Answer   int64    `parquet:"answer"`
	}](Root + "cais/mmlu/all/test-00000-of-00001.parquet")
	out := make([]Example, len(rows))
	for i, r := range rows {
		out[i] = Example{ID: fmt.Sprint("mmlu-", i), Task: "mmlu", Kind: "choice", Answers: []string{string(letters[r.Answer])},
			Prompt: r.Question + "\n" + options(r.Choices) + "\nFinish with a line of the form \"Answer: <letter>\"."}
	}
	return out, err
}

// MMLUPro: harder 10-way multiple choice.
func MMLUPro() ([]Example, error) {
	rows, err := Parquet[struct {
		ID       int64    `parquet:"question_id"`
		Question string   `parquet:"question"`
		Options  []string `parquet:"options,list"`
		Answer   string   `parquet:"answer"`
	}](Root + "TIGER-Lab/MMLU-Pro/data/test-00000-of-00001.parquet")
	out := make([]Example, len(rows))
	for i, r := range rows {
		out[i] = Example{ID: fmt.Sprint("mmlupro-", r.ID), Task: "mmlupro", Kind: "choice", Answers: []string{r.Answer},
			Prompt: r.Question + "\n" + options(r.Options) + "\nFinish with a line of the form \"Answer: <letter>\"."}
	}
	return out, err
}

// CNNDailyMail: news summarization, scored with ROUGE-L.
func CNNDailyMail() ([]Example, error) {
	rows, err := Parquet[struct {
		ID         string `parquet:"id"`
		Article    string `parquet:"article"`
		Highlights string `parquet:"highlights"`
	}](Root + "abisee/cnn_dailymail/3.0.0/test-00000-of-00001.parquet")
	out := make([]Example, len(rows))
	for i, r := range rows {
		out[i] = Example{ID: "cnndm-" + r.ID, Task: "cnndm", Kind: "summary", Answers: []string{r.Highlights},
			Context: "Article:\n" + r.Article, Prompt: "Summarize the article in three sentences."}
	}
	return out, err
}

// HumanEval: code completion prompts. Used as a generation-speed workload only (answers are not executed).
func HumanEval() ([]Example, error) {
	rows, err := Parquet[struct {
		ID     string `parquet:"task_id"`
		Prompt string `parquet:"prompt"`
	}](Root + "openai/openai_humaneval/openai_humaneval/test-00000-of-00001.parquet")
	out := make([]Example, len(rows))
	for i, r := range rows {
		out[i] = Example{ID: r.ID, Task: "humaneval", Kind: "none", Answers: []string{""},
			Prompt: "Complete the following Python function. Return the full function in a code block.\n\n" + r.Prompt}
	}
	return out, err
}

// LongBench loads one LongBench subset (e.g. "hotpotqa", "2wikimqa", "musique", "multi_news", "samsum").
// Documents go into Context so that requests sharing a document share a prefix.
// Examples longer than 100k characters (~25k tokens) are dropped so every request fits a 32k context.
func LongBench(subset string) ([]Example, error) {
	rows, err := JSONL[struct {
		ID      string   `json:"_id"`
		Input   string   `json:"input"`
		Context string   `json:"context"`
		Answers []string `json:"answers"`
	}](Root + "zai-org/LongBench/data/" + subset + ".jsonl")
	kind, prompt := "span", "Answer the question based on the passages above. Only give me the answer and do not output any other words.\n\nQuestion: %s\nAnswer:"
	switch subset {
	case "multi_news", "gov_report", "multi_news_e", "gov_report_e":
		kind, prompt = "summary", "Write a one-page summary of the documents above.%s"
	case "samsum", "samsum_e":
		kind, prompt = "summary", "Summarize the last dialogue above in the same style as the examples.\n\n%s\nSummary:"
	}
	var out []Example
	for _, r := range rows {
		if len(r.Context) <= 100_000 {
			out = append(out, Example{ID: "longbench-" + subset + "-" + r.ID, Task: "longbench-" + subset, Kind: kind, Answers: r.Answers,
				Context: r.Context, Prompt: fmt.Sprintf(prompt, r.Input)})
		}
	}
	return out, err
}

// Load returns the named task: gsm8k, math500, mmlu, mmlupro, cnndm, humaneval or longbench-<subset>.
func Load(task string) ([]Example, error) {
	switch task {
	case "gsm8k":
		return GSM8K()
	case "math500":
		return MATH500()
	case "mmlu":
		return MMLU()
	case "mmlupro":
		return MMLUPro()
	case "cnndm":
		return CNNDailyMail()
	case "humaneval":
		return HumanEval()
	}
	if subset, ok := strings.CutPrefix(task, "longbench-"); ok {
		return LongBench(subset)
	}
	return nil, fmt.Errorf("unknown task %q", task)
}
