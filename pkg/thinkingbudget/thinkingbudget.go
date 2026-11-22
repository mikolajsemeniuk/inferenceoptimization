// Package thinkingbudget decides how much a reasoning model may think for a given request.
// Every strategy returns the request options to use plus the cost of making the decision.
package thinkingbudget

import (
	"context"
	"strconv"
	"strings"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/llm"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/metrics"
)

// Decision is the outcome of a strategy for one request.
type Decision struct {
	Extra   map[string]any // request options: thinking on/off and budget
	Think   bool
	Budget  int     // thinking token cap, 0 = no cap (or thinking off)
	Seconds float64 // time spent deciding
	Tokens  int     // tokens spent deciding (prompt + completion)
}

// Fixed caps thinking at budget tokens (0 switches thinking off, negative means unlimited).
func Fixed(budget int) Decision {
	switch {
	case budget == 0:
		return Decision{Extra: llm.Think(false)}
	case budget < 0:
		return Decision{Extra: llm.Think(true), Think: true}
	}
	extra := llm.Think(true)
	extra["thinking_budget_tokens"] = budget
	return Decision{Extra: extra, Think: true, Budget: budget}
}

const estimatePrompt = `Task: analyze the question below and estimate how many tokens of step-by-step reasoning are needed to answer it correctly.
Reply with a single integer between 0 and 2000 and nothing else. Use 0 if the answer can be given directly.

Question:
`

// Estimate is the TALE-style strategy (Token-Budget-Aware LLM Reasoning): the model at base
// first estimates its own reasoning budget without thinking, then gets exactly that budget.
func Estimate(ctx context.Context, base, question string) Decision {
	res, err := llm.Chat(ctx, base, llm.Request{Messages: llm.User(estimatePrompt + question), MaxTokens: 8, Extra: llm.Think(false)})
	budget, convErr := strconv.Atoi(metrics.LastNumber(res.Content))
	if err != nil || convErr != nil {
		budget = -1 // could not estimate: think freely
	}
	d := Fixed(min(budget, 4096))
	d.Seconds, d.Tokens = res.Latency, res.PromptTokens+res.CompletionTokens
	return d
}

const routePrompt = `Decide whether answering the question below requires multi-step reasoning or calculation.
Reply with exactly one word: "reason" if it does, "direct" if it can be answered from knowledge in one step.

Question:
`

// Route is the "When to Reason" strategy: a cheap classifier switches thinking on or off.
// The paper trains a small encoder; here a small LLM at base is asked instead (no training).
func Route(ctx context.Context, base, question string) Decision {
	res, err := llm.Chat(ctx, base, llm.Request{Messages: llm.User(routePrompt + question), MaxTokens: 4, Extra: llm.Think(false)})
	d := Fixed(-1)
	if err == nil && strings.Contains(strings.ToLower(res.Content), "direct") {
		d = Fixed(0)
	}
	d.Seconds, d.Tokens = res.Latency, res.PromptTokens+res.CompletionTokens
	return d
}
