// Package promptcompress shortens a long context before it is sent to the model.
// Training-free, query-aware methods: split the context into chunks, score every chunk
// against the question, keep the best chunks (in their original order) within a budget.
// These are the extractive baselines of the LLMLingua papers, plus compression by a small LLM.
package promptcompress

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sort"
	"strings"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/llm"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/metrics"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/vector"
)

// Chunks splits text into pieces of about n words, breaking at line ends where possible.
func Chunks(text string, n int) []string {
	var out []string
	var cur []string
	words := 0
	for _, line := range strings.Split(text, "\n") {
		w := len(strings.Fields(line))
		if words > 0 && words+w > n {
			out, cur, words = append(out, strings.Join(cur, "\n")), nil, 0
		}
		cur, words = append(cur, line), words+w
	}
	if words > 0 {
		out = append(out, strings.Join(cur, "\n"))
	}
	return out
}

// Keep returns the highest-scoring chunks whose total length stays within ratio of the
// original, in their original order.
func Keep(chunks []string, scores []float64, ratio float64) string {
	total := 0
	for _, c := range chunks {
		total += len(c)
	}
	order := make([]int, len(chunks))
	for i := range order {
		order[i] = i
	}
	sort.SliceStable(order, func(a, b int) bool { return scores[order[a]] > scores[order[b]] })
	budget, keep := int(ratio*float64(total)), map[int]bool{}
	for _, i := range order {
		if len(chunks[i]) <= budget {
			keep[i], budget = true, budget-len(chunks[i])
		}
	}
	var out []string
	for i, c := range chunks {
		if keep[i] {
			out = append(out, c)
		}
	}
	return strings.Join(out, "\n")
}

// ScoreBM25 scores chunks by lexical overlap with the question.
func ScoreBM25(question string, chunks []string) []float64 { return metrics.BM25(question, chunks) }

// ScoreFirst prefers earlier chunks (plain truncation).
func ScoreFirst(chunks []string) []float64 {
	s := make([]float64, len(chunks))
	for i := range s {
		s[i] = -float64(i)
	}
	return s
}

// ScoreRandom is the control: random chunks.
func ScoreRandom(chunks []string, seed uint64) []float64 {
	r := rand.New(rand.NewPCG(seed, seed))
	s := make([]float64, len(chunks))
	for i := range s {
		s[i] = r.Float64()
	}
	return s
}

// ScoreEmbedding scores chunks by cosine similarity to the question, using the embedding server at base.
func ScoreEmbedding(ctx context.Context, base, question string, chunks []string) ([]float64, error) {
	scores := make([]float64, len(chunks))
	q, err := llm.Embed(ctx, base, []string{question})
	if err != nil {
		return scores, err
	}
	qv := vector.Normalize(q[0])
	for start := 0; start < len(chunks); start += 32 {
		batch := chunks[start:min(start+32, len(chunks))]
		vecs, err := llm.Embed(ctx, base, batch)
		if err != nil {
			return scores, err
		}
		for i, v := range vecs {
			scores[start+i] = float64(vector.Dot(qv, vector.Normalize(v)))
		}
	}
	return scores, nil
}

// LLM asks the model at base to extract what is needed to answer the question, in about words words.
func LLM(ctx context.Context, base, question, text string, words int) (string, llm.Response, error) {
	prompt := fmt.Sprintf("%s\n\nCopy from the text above only the sentences needed to answer the question below, at most %d words. Do not answer the question.\n\nQuestion: %s", text, words, question)
	res, err := llm.Chat(ctx, base, llm.Request{Messages: llm.User(prompt), MaxTokens: 2 * words, Extra: llm.Think(false)})
	return res.Content, res, err
}
