// Package cascade answers with a cheap model first and escalates to a strong model only when
// the cheap answer looks unreliable (FrugalGPT / AutoMix style, without any trained scorer).
package cascade

import (
	"context"
	"strings"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/llm"
)

// Vote returns the most frequent non-empty answer and the share of samples that agree with it.
// This is the self-consistency confidence signal.
func Vote(answers []string) (winner string, agreement float64) {
	count := map[string]int{}
	for _, a := range answers {
		if a != "" {
			count[a]++
		}
	}
	best := 0
	for a, n := range count {
		if n > best || (n == best && a < winner) {
			winner, best = a, n
		}
	}
	return winner, float64(best) / float64(max(len(answers), 1))
}

const verifyPrompt = `Question:
%QUESTION%

Proposed answer:
%ANSWER%

Is the proposed answer correct? Reply with exactly one word: "yes" or "no".`

// Verify asks the model at base whether the answer is correct (AutoMix-style self-verification).
// It returns the verdict and the response with its cost.
func Verify(ctx context.Context, base, question, answer string) (bool, llm.Response, error) {
	prompt := strings.NewReplacer("%QUESTION%", question, "%ANSWER%", answer).Replace(verifyPrompt)
	res, err := llm.Chat(ctx, base, llm.Request{Messages: llm.User(prompt), MaxTokens: 4, Extra: llm.Think(false)})
	return strings.Contains(strings.ToLower(res.Content), "yes"), res, err
}

// Escalate is the cascade decision: keep the cheap answer when confidence reaches the threshold.
func Escalate(confidence, threshold float64) bool { return confidence < threshold }
