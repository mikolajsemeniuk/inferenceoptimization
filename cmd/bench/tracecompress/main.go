// Benchmark: reasoning-trace compression. Traces logged by the baseline (thinking) run are
// compressed in several ways; a model with thinking switched off must then answer using only
// the compressed trace. Shows how much of a trace must be stored to keep its value.
//
//	none        no trace (thinking off)            cod         Chain-of-Draft prompting, no trace
//	full        the whole trace                    head-N/tail-N   first / last N words
//	summary-N   small-model summary in <= N words
//
//	make bench-tracecompress
package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/bench"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/dataset"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/tracecompress"
)

func main() {
	small := flag.String("small-url", "http://localhost:8102", "small model that writes the summaries")
	from := flag.String("traces", "output/baseline/Q8_0/think/", "baseline run whose traces are compressed")
	c := bench.Flags("tracecompress", "gsm8k,math500,mmlupro", false)
	tasks := c.Tasks

	traces := map[string]string{} // example id -> full trace
	for _, task := range tasks {
		rows, _ := dataset.JSONL[bench.Row](*from + task + ".jsonl")
		for _, r := range rows {
			if r.Error == "" && r.FinishReason == "stop" {
				traces[r.ID] = r.Reasoning
			}
		}
	}
	fmt.Println("traces loaded:", len(traces))
	c.Keep = func(e dataset.Example) bool { _, ok := traces[e.ID]; return ok } // baseline truncated: no usable trace

	type method func(ctx context.Context, e dataset.Example) (prompt string, seconds float64, tokens int)
	keep := func(f func(string) string) method {
		return func(_ context.Context, e dataset.Example) (string, float64, int) {
			return tracecompress.WithTrace(e.Prompt, f(traces[e.ID])), 0, 0
		}
	}
	summary := func(n int) method {
		return func(ctx context.Context, e dataset.Example) (string, float64, int) {
			s, res, _ := tracecompress.Summarize(ctx, *small, traces[e.ID], n)
			return tracecompress.WithTrace(e.Prompt, s), res.Latency, res.PromptTokens + res.CompletionTokens
		}
	}
	methods := []struct {
		name string
		f    method
	}{
		{"none", func(_ context.Context, e dataset.Example) (string, float64, int) { return e.Prompt, 0, 0 }},
		{"cod", func(_ context.Context, e dataset.Example) (string, float64, int) {
			return tracecompress.ChainOfDraft(e.Prompt), 0, 0
		}},
		{"full", keep(func(t string) string { return t })},
		{"head-50", keep(func(t string) string { return tracecompress.Head(t, 50) })},
		{"head-150", keep(func(t string) string { return tracecompress.Head(t, 150) })},
		{"tail-50", keep(func(t string) string { return tracecompress.Tail(t, 50) })},
		{"tail-150", keep(func(t string) string { return tracecompress.Tail(t, 150) })},
		{"summary-30", summary(30)},
		{"summary-80", summary(80)},
	}
	for _, m := range methods {
		c.Tag = m.name
		c.Ask = func(ctx context.Context, c bench.Config, e dataset.Example) bench.Row {
			prompt, seconds, tokens := m.f(ctx, e)
			words := len(strings.Fields(prompt)) - len(strings.Fields(e.Prompt))
			e.Prompt = prompt
			row := bench.Ask(ctx, c, e)
			row.OverheadSeconds, row.OverheadTokens, row.Note = seconds, tokens, fmt.Sprintf("trace_words=%d full_words=%d", max(words, 0), len(strings.Fields(traces[e.ID])))
			return row
		}
		bench.Run(c)
	}
}
