// Package cachehint reuses a cached neighbour that is not trusted enough to be returned
// directly: its question, answer and (optionally) reasoning are given to the model as a hint.
package cachehint

import (
	"context"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/llm"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/vector"
)

// Entry is one cached request.
type Entry struct {
	Question  string
	Answer    string
	Reasoning string
}

// Prompt puts the cached entry in front of the new question. trace=false leaves the reasoning out.
func Prompt(question string, e Entry, trace bool) string {
	hint := "A similar question was answered earlier.\n\nEarlier question:\n" + e.Question
	if trace && e.Reasoning != "" {
		hint += "\n\nEarlier reasoning:\n" + e.Reasoning
	}
	hint += "\n\nEarlier answer:\n" + e.Answer
	return hint + "\n\nIf the earlier answer applies to the new question, reuse it; otherwise solve the new question yourself.\n\nNew question:\n" + question
}

// Paraphrase rewrites a question with different wording but the same content (model at base).
func Paraphrase(ctx context.Context, base, question string) (string, error) {
	prompt := "Rewrite the problem below using different wording and sentence structure. Keep every number, name, quantity and the final question exactly equivalent, so the answer does not change. Output only the rewritten problem.\n\n" + question
	extra := llm.Think(false)
	extra["temperature"] = 0.7
	res, err := llm.Chat(ctx, base, llm.Request{Messages: llm.User(prompt), MaxTokens: 1024, Seed: 1, Extra: extra})
	return res.Content, err
}

// Neighbours returns, for query vector q, the most similar stored vector and the most similar
// one that is not `own` (the entry that truly matches q): the best case and the near-miss case.
func Neighbours(q []float32, store [][]float32, own int) (best int, bestSim float32, other int, otherSim float32) {
	best, other, bestSim, otherSim = -1, -1, -2, -2
	for i, v := range store {
		s := vector.Dot(q, v)
		if s > bestSim {
			best, bestSim = i, s
		}
		if i != own && s > otherSim {
			other, otherSim = i, s
		}
	}
	return
}
