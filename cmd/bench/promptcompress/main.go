// Benchmark: prompt compression on long-context QA. The context is cut to a share of its
// original size before the reference model sees it.
//
//	full                       no compression
//	first|random|bm25|embedding-R   keep the best ~100-word chunks up to ratio R
//	llm-R                      a small model extracts the relevant sentences
//
//	make bench-promptcompress
package main

import (
	"context"
	"flag"
	"fmt"
	"hash/fnv"
	"strings"
	"time"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/bench"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/dataset"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/promptcompress"
)

func main() {
	small := flag.String("small-url", "http://localhost:8102", "small model for llm compression")
	embed := flag.String("embed-url", "http://localhost:8103", "embedding server")
	c := bench.Flags("promptcompress", "longbench-hotpotqa,longbench-2wikimqa,longbench-musique", false)

	// the question is the last "Question: ..." line of the prompt
	question := func(e dataset.Example) string {
		_, q, _ := strings.Cut(e.Prompt, "Question: ")
		return strings.TrimSuffix(q, "\nAnswer:")
	}
	type method func(ctx context.Context, e dataset.Example, ratio float64) (string, int)
	chunked := func(score func(ctx context.Context, q string, chunks []string, e dataset.Example) []float64) method {
		return func(ctx context.Context, e dataset.Example, ratio float64) (string, int) {
			chunks := promptcompress.Chunks(e.Context, 100)
			return promptcompress.Keep(chunks, score(ctx, question(e), chunks, e), ratio), 0
		}
	}
	methods := map[string]method{
		"first": chunked(func(_ context.Context, _ string, ch []string, _ dataset.Example) []float64 {
			return promptcompress.ScoreFirst(ch)
		}),
		"random": chunked(func(_ context.Context, _ string, ch []string, e dataset.Example) []float64 {
			h := fnv.New64a()
			h.Write([]byte(e.ID))
			return promptcompress.ScoreRandom(ch, h.Sum64())
		}),
		"bm25": chunked(func(_ context.Context, q string, ch []string, _ dataset.Example) []float64 {
			return promptcompress.ScoreBM25(q, ch)
		}),
		"embedding": chunked(func(ctx context.Context, q string, ch []string, _ dataset.Example) []float64 {
			s, _ := promptcompress.ScoreEmbedding(ctx, *embed, q, ch)
			return s
		}),
		"llm": func(ctx context.Context, e dataset.Example, ratio float64) (string, int) {
			words := int(ratio * float64(len(strings.Fields(e.Context))))
			out, res, _ := promptcompress.LLM(ctx, *small, question(e), e.Context, words)
			return out, res.PromptTokens + res.CompletionTokens
		},
	}

	type variant struct {
		name  string
		ratio float64
	}
	variants := []variant{{"full", 1}}
	for _, r := range []float64{0.5, 0.25, 0.1} {
		for _, m := range []string{"first", "random", "bm25", "embedding"} {
			variants = append(variants, variant{m, r})
		}
	}
	variants = append(variants, variant{"llm", 0.1}, variant{"llm", 0.03})

	for _, v := range variants {
		c.Tag = v.name
		if v.name != "full" {
			c.Tag = fmt.Sprintf("%s-%.2f", v.name, v.ratio)
		}
		c.Ask = func(ctx context.Context, c bench.Config, e dataset.Example) bench.Row {
			original := len(e.Context)
			start, tokens := time.Now(), 0
			if v.name != "full" {
				e.Context, tokens = methods[v.name](ctx, e, v.ratio)
			}
			overhead := time.Since(start).Seconds()
			row := bench.Ask(ctx, c, e)
			row.OverheadSeconds, row.OverheadTokens = overhead, tokens
			row.Note = fmt.Sprintf("chars=%d original_chars=%d", len(e.Context), original)
			return row
		}
		bench.Run(c)
	}
}
