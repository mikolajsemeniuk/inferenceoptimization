// Package thinkingbudget decides how much a reasoning model may think for a given request.
// Every strategy returns the request options to use plus the cost of making the decision.
package thinkingbudget

import (
	"context"
	"strconv"

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

// The question is quoted first and the instruction comes last; otherwise the model follows the
// answer-format instruction inside the question and answers it instead of judging it.
func judge(question, instruction string) string {
	return "<question>\n" + question + "\n</question>\n\nDo not answer the question above. " + instruction
}

const estimatePrompt = `Estimate how many tokens of step-by-step reasoning a careful solver needs to answer it correctly: about 100 for a one-step problem, 300-600 for a typical multi-step problem, 1000-2000 for a hard one, 0 if the answer can be given directly.
Reply with a single integer between 0 and 2000 and nothing else.`

// estimate asks the model at base for a reasoning budget; -1 means it gave no usable number.
func estimate(ctx context.Context, base, question string) (int, llm.Response) {
	res, err := llm.Chat(ctx, base, llm.Request{Messages: llm.User(judge(question, estimatePrompt)), MaxTokens: 8, Extra: llm.Think(false)})
	budget, convErr := strconv.Atoi(metrics.LastNumber(res.Content))
	if err != nil || convErr != nil {
		return -1, res
	}
	return budget, res
}

// Estimate is the TALE-style strategy (Token-Budget-Aware LLM Reasoning): the model at base
// first estimates its own reasoning budget without thinking, then gets exactly that budget.
// Without a usable estimate it thinks freely.
func Estimate(ctx context.Context, base, question string) Decision {
	budget, res := estimate(ctx, base, question)
	d := Fixed(min(budget, 4096))
	d.Seconds, d.Tokens = res.Latency, res.PromptTokens+res.CompletionTokens
	return d
}

// Route is the "When to Reason" strategy: thinking is switched fully on or off per request.
// The paper trains a small encoder for this; asking a model "is this hard?" gave the same reply
// for every question, so the training-free stand-in is the model's own budget estimate:
// thinking stays off when the estimate is at most threshold tokens.
func Route(ctx context.Context, base, question string, threshold int) Decision {
	budget, res := estimate(ctx, base, question)
	d := Fixed(-1)
	if budget >= 0 && budget <= threshold {
		d = Fixed(0)
	}
	d.Seconds, d.Tokens = res.Latency, res.PromptTokens+res.CompletionTokens
	return d
}
