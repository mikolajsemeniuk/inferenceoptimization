// Package tracecompress shrinks a model's reasoning trace so it can be stored (e.g. in a cache)
// and reused later as a hint. All methods are training-free.
package tracecompress

import (
	"context"
	"fmt"
	"strings"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/llm"
)

// Head keeps the first n words of the trace.
func Head(trace string, n int) string {
	w := strings.Fields(trace)
	return strings.Join(w[:min(n, len(w))], " ")
}

// Tail keeps the last n words of the trace (where the conclusion usually is).
func Tail(trace string, n int) string {
	w := strings.Fields(trace)
	return strings.Join(w[max(0, len(w)-n):], " ")
}

// Summarize asks the model at base to rewrite the trace as its key steps in at most n words.
func Summarize(ctx context.Context, base, trace string, n int) (string, llm.Response, error) {
	prompt := fmt.Sprintf("Compress the reasoning below into its essential steps and the result, in at most %d words. Keep every number that matters. Output only the compressed reasoning.\n\nReasoning:\n%s", n, trace)
	res, err := llm.Chat(ctx, base, llm.Request{Messages: llm.User(prompt), MaxTokens: 3 * n, Extra: llm.Think(false)})
	return res.Content, res, err
}

// ChainOfDraft wraps a question with the Chain-of-Draft instruction: reason in minimal drafts
// instead of full sentences (Xu et al., "Chain of Draft: Thinking Faster by Writing Less").
func ChainOfDraft(question string) string {
	return "Think step by step, but only keep a minimum draft for each thinking step, with 5 words at most. Then give the final answer.\n\n" + question
}

// WithTrace builds the prompt that lets a model answer the question using a stored trace.
func WithTrace(question, trace string) string {
	return "Reasoning notes for the question below (they may be abbreviated):\n" + trace + "\n\nUse the notes to answer.\n\n" + question
}
