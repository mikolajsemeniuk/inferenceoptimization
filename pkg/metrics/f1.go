package metrics

import "strings"

// F1 is the token-level F1 between prediction and reference (SQuAD / HotpotQA style).
func F1(prediction, reference string) float64 {
	p, r := strings.Fields(Normalize(prediction)), strings.Fields(Normalize(reference))
	if len(p) == 0 || len(r) == 0 {
		if len(p) == len(r) {
			return 1
		}
		return 0
	}
	left := map[string]int{}
	for _, t := range r {
		left[t]++
	}
	same := 0
	for _, t := range p {
		if left[t] > 0 {
			left[t]--
			same++
		}
	}
	if same == 0 {
		return 0
	}
	precision, recall := float64(same)/float64(len(p)), float64(same)/float64(len(r))
	return 2 * precision * recall / (precision + recall)
}

// BestF1 is the maximum F1 over several acceptable references.
func BestF1(prediction string, references []string) float64 {
	best := 0.0
	for _, r := range references {
		best = max(best, F1(prediction, r))
	}
	return best
}

// PrecisionRecall for binary decisions (e.g. cache hits vs. should-have-hit).
func PrecisionRecall(tp, fp, fn int) (precision, recall, f1 float64) {
	precision, recall = Rate(tp, tp+fp), Rate(tp, tp+fn)
	if precision+recall > 0 {
		f1 = 2 * precision * recall / (precision + recall)
	}
	return
}
