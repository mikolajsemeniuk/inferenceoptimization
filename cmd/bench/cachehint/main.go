// Benchmark: cached neighbour as a hint. The cache holds the baseline answers (question, answer,
// reasoning). New requests are paraphrases of the cached questions. For each request the nearest
// cache entry is retrieved by embedding and used in different ways:
//
//	nohint-think / nohint-nothink     cache ignored (the two cost extremes)
//	hint-qa-nothink                   neighbour's question + answer in the prompt, thinking off
//	hint-qat-nothink                  ... plus the neighbour's reasoning trace
//	hint-qa-think256                  ... question + answer, thinking capped at 256 tokens
//	wrong-qa-nothink / wrong-qa-think256   same, but with the nearest *different* question:
//	                                  the near-miss case, shows what a misleading hint costs
//
// Run it on a server with n-gram speculative decoding (-tag ngram) to measure the cached answer
// acting as a draft: make bench-cachehint
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/bench"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/cachehint"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/dataset"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/llm"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/output"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/run"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/vector"
)

type paraphrase struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type lookup struct {
	ID          string  `json:"id"`
	Retrieved   string  `json:"retrieved"` // id of the nearest cache entry
	Similarity  float32 `json:"similarity"`
	Correct     bool    `json:"correct"`   // nearest entry is the paraphrased original
	NearMiss    string  `json:"near_miss"` // nearest entry that is a different question
	NearMissSim float32 `json:"near_miss_similarity"`
}

func main() {
	embed := flag.String("embed-url", "http://localhost:8103", "embedding server")
	from := flag.String("cache", "output/baseline/Q8_0/think/", "baseline run that fills the cache")
	only := flag.String("modes", "", "comma-separated subset of modes (default all)")
	c := bench.Flags("cachehint", "gsm8k,mmlupro", false)
	ctx := context.Background()
	server := c.Tag

	for _, task := range c.Tasks {
		// 1. cache content = baseline answers
		logged, err := dataset.JSONL[bench.Row](*from + task + ".jsonl")
		if err != nil {
			log.Fatal(err)
		}
		var entries []cachehint.Entry
		var originals []dataset.Example
		for _, r := range logged {
			if r.Error == "" && r.FinishReason == "stop" {
				entries = append(entries, cachehint.Entry{Question: r.Example.Prompt, Answer: r.Content, Reasoning: r.Reasoning})
				originals = append(originals, r.Example)
			}
		}
		// the logged rows do not carry the prompt text: take it from the dataset
		all, err := dataset.Load(task)
		if err != nil {
			log.Fatal(err)
		}
		text := map[string]dataset.Example{}
		for _, e := range all {
			text[e.ID] = e
		}
		for i := range originals {
			originals[i] = text[originals[i].ID]
			entries[i].Question = originals[i].Prompt
		}
		if c.N < len(originals) {
			originals, entries = originals[:c.N], entries[:c.N]
		}

		// 2. new requests = paraphrases (generated once, then reused by every run)
		file := "output/cachehint/paraphrases-" + task + ".jsonl"
		paraphrases, err := dataset.JSONL[paraphrase](file)
		if err != nil || len(paraphrases) != len(originals) {
			paraphrases = run.Map(originals, c.Workers, func(_ int, e dataset.Example) paraphrase {
				body, instruction, _ := strings.Cut(e.Prompt, "\n\n")
				if task == "mmlupro" { // keep the options, rewrite only the question stem
					stem, rest, _ := strings.Cut(e.Prompt, "\nA. ")
					body, instruction = stem, "A. "+rest
				}
				p, err := cachehint.Paraphrase(ctx, c.URL, body)
				if err != nil {
					log.Fatal(err)
				}
				sep := "\n\n"
				if task == "mmlupro" {
					sep = "\n"
				}
				return paraphrase{e.ID, p + sep + instruction}
			})
			if err := output.JSONL(file, paraphrases); err != nil {
				log.Fatal(err)
			}
		}

		// 3. retrieval
		embedAll := func(texts []string) [][]float32 {
			var out [][]float32
			for start := 0; start < len(texts); start += 16 {
				v, err := llm.Embed(ctx, *embed, texts[start:min(start+16, len(texts))])
				if err != nil {
					log.Fatal(err)
				}
				for _, x := range v {
					out = append(out, vector.Normalize(x))
				}
			}
			return out
		}
		var cacheTexts, queryTexts []string
		for i := range originals {
			cacheTexts, queryTexts = append(cacheTexts, originals[i].Prompt), append(queryTexts, paraphrases[i].Text)
		}
		store, queries := embedAll(cacheTexts), embedAll(queryTexts)
		lookups := make([]lookup, len(originals))
		best, miss := make([]int, len(originals)), make([]int, len(originals))
		hits := 0
		for i := range originals {
			b, bs, o, osim := cachehint.Neighbours(queries[i], store, i)
			best[i], miss[i] = b, o
			lookups[i] = lookup{originals[i].ID, originals[b].ID, bs, b == i, originals[o].ID, osim}
			if b == i {
				hits++
			}
		}
		fmt.Printf("%s: %d cache entries, retrieval accuracy %.3f\n", task, len(originals), float64(hits)/float64(len(originals)))
		if err := output.JSONL("output/cachehint/lookup-"+task+".jsonl", lookups); err != nil {
			log.Fatal(err)
		}

		// 4. answer the paraphrases in every mode
		index := map[string]int{}
		requests := make([]dataset.Example, len(originals))
		for i, e := range originals {
			index[e.ID] = i
			requests[i] = e
			requests[i].Prompt = paraphrases[i].Text
		}
		type mode struct {
			name   string
			think  bool
			budget int
			prompt func(i int, q string) string
		}
		plain := func(_ int, q string) string { return q }
		hint := func(pick []int, trace bool) func(int, string) string {
			return func(i int, q string) string { return cachehint.Prompt(q, entries[pick[i]], trace) }
		}
		modes := []mode{
			{"nohint-think", true, 0, plain},
			{"nohint-nothink", false, 0, plain},
			{"hint-qa-nothink", false, 0, hint(best, false)},
			{"hint-qat-nothink", false, 0, hint(best, true)},
			{"hint-qa-think256", true, 256, hint(best, false)},
			{"wrong-qa-nothink", false, 0, hint(miss, false)},
			{"wrong-qa-think256", true, 256, hint(miss, false)},
		}
		for _, m := range modes {
			if *only != "" && !strings.Contains(","+*only+",", ","+m.name+",") {
				continue
			}
			cfg := c
			cfg.Tasks, cfg.Tag, cfg.Think, cfg.N = []string{task}, server+"/"+m.name, m.think, 0
			if m.budget > 0 {
				cfg.Extra = llm.Budget(m.budget)
			}
			cfg.Ask = func(ctx context.Context, cfg bench.Config, e dataset.Example) bench.Row {
				i, ok := index[e.ID]
				if !ok {
					return bench.Row{Example: e, Tag: cfg.Tag, Error: "not cached"}
				}
				e.Prompt = m.prompt(i, requests[i].Prompt)
				row := bench.Ask(ctx, cfg, e)
				row.Note = fmt.Sprintf("similarity=%.3f retrieved_correct=%v", lookups[i].Similarity, lookups[i].Correct)
				return row
			}
			bench.RunExamples(cfg, task, originals)
		}
	}
}
