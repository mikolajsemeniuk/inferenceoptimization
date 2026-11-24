// Package bench runs benchmark tasks against a llama-server and stores the results.
// Model-level benchmarks (baseline, quantization, KV-cache quantization, speculative decoding)
// differ only in how the server was started, so they share this code.
package bench

import (
	"context"
	"flag"
	"fmt"
	"log"
	"maps"
	"slices"
	"strings"
	"time"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/dataset"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/llm"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/metrics"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/output"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/run"
)

// Config says what to run and where to store it.
type Config struct {
	Name      string   // benchmark name -> output/<Name>/<Tag>/
	Tag       string   // server configuration being measured, e.g. "Q4_K_M"
	URL       string   // llama-server base URL
	Tasks     []string // dataset.Load names
	N         int      // examples per task
	Workers   int      // concurrent requests
	MaxTokens int
	Think     bool
	Seed      uint64
	Extra     map[string]any // extra request options
	// Ask overrides how one example is answered (nil = a single plain request, see Ask).
	Ask func(context.Context, Config, dataset.Example) Row
	// Keep drops sampled examples a benchmark cannot use (nil = keep all).
	Keep func(dataset.Example) bool
}

// Flags reads the common command line flags, starting from the given defaults.
func Flags(name, tasks string, think bool) Config {
	c := Config{Name: name}
	var list string
	flag.StringVar(&c.Tag, "tag", "default", "label of the server configuration under test")
	flag.StringVar(&c.URL, "url", "http://localhost:8101", "llama-server base URL")
	flag.StringVar(&list, "tasks", tasks, "comma-separated tasks")
	flag.IntVar(&c.N, "n", 200, "examples per task")
	flag.IntVar(&c.Workers, "workers", 8, "concurrent requests (1 for clean latency numbers)")
	flag.IntVar(&c.MaxTokens, "max-tokens", 4096, "generation limit, thinking included")
	flag.BoolVar(&c.Think, "think", think, "enable model thinking")
	flag.Uint64Var(&c.Seed, "seed", 1, "sampling seed for the examples")
	flag.Parse()
	c.Tasks = strings.Split(list, ",")
	return c
}

// Row is one request: the example, the model's answer, its score and all measurements.
type Row struct {
	dataset.Example
	Tag   string  `json:"tag"`
	Think bool    `json:"think"`
	Score float64 `json:"score"`
	Error string  `json:"error,omitempty"`
	// Extra calls made before the answer (router, budget estimate, ...): their cost.
	OverheadSeconds float64 `json:"overhead_s"`
	OverheadTokens  int     `json:"overhead_tokens"`
	Note            string  `json:"note,omitempty"` // method-specific decision, e.g. the chosen budget
	llm.Response
}

// Ask sends one example and grades the answer.
func Ask(ctx context.Context, c Config, e dataset.Example) Row {
	extra := llm.Think(c.Think)
	maps.Copy(extra, c.Extra)
	res, err := llm.Chat(ctx, c.URL, llm.Request{Messages: llm.User(e.Text()), MaxTokens: c.MaxTokens, Extra: extra})
	row := Row{Example: e, Tag: c.Tag, Think: c.Think, Response: res}
	if err != nil {
		row.Error = err.Error()
		return row
	}
	row.ReasoningTokens, _ = llm.Tokens(ctx, c.URL, res.Reasoning)
	row.Score = dataset.Grade(e, res.Content)
	return row
}

// Summary aggregates the rows of one task.
type Summary struct {
	Name             string          `json:"benchmark"`
	Tag              string          `json:"tag"`
	Task             string          `json:"task"`
	Think            bool            `json:"think"`
	Workers          int             `json:"workers"`
	N                int             `json:"n"`
	Errors           int             `json:"errors"`
	Truncated        int             `json:"truncated"` // stopped by max_tokens
	Score            float64         `json:"score"`     // accuracy / F1 / ROUGE-L depending on the task
	WallSeconds      float64         `json:"wall_s"`
	Throughput       float64         `json:"throughput_tokens_per_s"` // all completion tokens / wall time
	Latency          metrics.Summary `json:"latency_s"`
	TTFT             metrics.Summary `json:"ttft_s"`
	PromptSeconds    metrics.Summary `json:"prompt_s"`
	TokensPerSecond  metrics.Summary `json:"tokens_per_s"` // per-request generation speed
	PromptTokens     metrics.Summary `json:"prompt_tokens"`
	CompletionTokens metrics.Summary `json:"completion_tokens"`
	ReasoningTokens  metrics.Summary `json:"reasoning_tokens"`
	CachedTokens     metrics.Summary `json:"cached_tokens"`
	DraftAcceptance  float64         `json:"draft_acceptance"` // accepted / drafted tokens
	OverheadSeconds  metrics.Summary `json:"overhead_s"`       // extra calls before the answer
	OverheadTokens   metrics.Summary `json:"overhead_tokens"`
	TotalSeconds     metrics.Summary `json:"total_s"` // latency + overhead
}

// Summarize builds the Summary of one task's rows.
func Summarize(c Config, task string, rows []Row, wall time.Duration) Summary {
	s := Summary{Name: c.Name, Tag: c.Tag, Task: task, Think: c.Think, Workers: c.Workers, N: len(rows), WallSeconds: wall.Seconds()}
	col := map[string][]float64{}
	drafted, accepted := 0, 0
	for _, r := range rows {
		if r.Error != "" {
			s.Errors++
			continue
		}
		if r.FinishReason == "length" {
			s.Truncated++
		}
		drafted, accepted = drafted+r.DraftTokens, accepted+r.DraftAccepted
		for k, v := range map[string]float64{"score": r.Score, "latency": r.Latency, "ttft": r.TTFT, "prompt_s": r.PromptSeconds, "tps": r.TokensPerSecond,
			"prompt": float64(r.PromptTokens), "completion": float64(r.CompletionTokens), "reasoning": float64(r.ReasoningTokens), "cached": float64(r.CachedTokens),
			"overhead_s": r.OverheadSeconds, "overhead_tokens": float64(r.OverheadTokens), "total_s": r.Latency + r.OverheadSeconds} {
			col[k] = append(col[k], v)
		}
	}
	s.Score = metrics.Sum(col["score"]) / float64(max(len(rows), 1)) // errors count as 0
	s.Throughput = metrics.Sum(col["completion"]) / wall.Seconds()
	s.Latency, s.TTFT, s.PromptSeconds, s.TokensPerSecond = metrics.Summarize(col["latency"]), metrics.Summarize(col["ttft"]), metrics.Summarize(col["prompt_s"]), metrics.Summarize(col["tps"])
	s.PromptTokens, s.CompletionTokens, s.ReasoningTokens, s.CachedTokens = metrics.Summarize(col["prompt"]), metrics.Summarize(col["completion"]), metrics.Summarize(col["reasoning"]), metrics.Summarize(col["cached"])
	s.OverheadSeconds, s.OverheadTokens, s.TotalSeconds = metrics.Summarize(col["overhead_s"]), metrics.Summarize(col["overhead_tokens"]), metrics.Summarize(col["total_s"])
	s.DraftAcceptance = metrics.Rate(accepted, drafted)
	return s
}

// Dir is where the results of this configuration go.
func Dir(c Config) string {
	mode := "nothink"
	if c.Think {
		mode = "think"
	}
	return fmt.Sprintf("output/%s/%s/%s/", c.Name, c.Tag, mode)
}

// RunExamples answers the given examples, writes <task>.jsonl and returns the summary.
func RunExamples(c Config, task string, examples []dataset.Example) Summary {
	ctx := context.Background()
	ask := Ask
	if c.Ask != nil {
		ask = c.Ask
	}
	start := time.Now()
	rows := run.Map(examples, c.Workers, func(_ int, e dataset.Example) Row { return ask(ctx, c, e) })
	s := Summarize(c, task, rows, time.Since(start))
	fmt.Printf("%-12s %-10s %-22s think=%-5v n=%d score=%.3f latency=%.1fs reasoning=%.0f tok/s=%.1f throughput=%.0f errors=%d truncated=%d (%.0fs)\n",
		c.Name, c.Tag, task, c.Think, s.N, s.Score, s.Latency.Mean, s.ReasoningTokens.Mean, s.TokensPerSecond.Mean, s.Throughput, s.Errors, s.Truncated, s.WallSeconds)
	if err := output.JSONL(Dir(c)+task+".jsonl", rows); err != nil {
		log.Fatal(err)
	}
	if err := output.JSON(Dir(c)+task+".summary.json", s); err != nil {
		log.Fatal(err)
	}
	return s
}

// Run executes every task, writes <task>.jsonl and summary.json, and returns the summaries.
func Run(c Config) []Summary {
	var summaries []Summary
	for _, task := range c.Tasks {
		examples, err := dataset.Load(task)
		if err != nil {
			log.Fatal(err)
		}
		examples = dataset.Sample(examples, c.N, c.Seed)
		if c.Keep != nil {
			examples = slices.DeleteFunc(examples, func(e dataset.Example) bool { return !c.Keep(e) })
		}
		summaries = append(summaries, RunExamples(c, task, examples))
		if err := output.JSON(Dir(c)+"summary.json", summaries); err != nil {
			log.Fatal(err)
		}
	}
	return summaries
}
