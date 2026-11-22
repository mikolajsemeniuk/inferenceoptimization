package metrics

import "strings"

// RougeL is the ROUGE-L F-measure (longest common subsequence of tokens).
func RougeL(prediction, reference string) float64 {
	p, r := strings.Fields(Normalize(prediction)), strings.Fields(Normalize(reference))
	if len(p) == 0 || len(r) == 0 {
		return 0
	}
	prev, cur := make([]int, len(r)+1), make([]int, len(r)+1)
	for i := range p {
		for j := range r {
			if p[i] == r[j] {
				cur[j+1] = prev[j] + 1
			} else {
				cur[j+1] = max(cur[j], prev[j+1])
			}
		}
		prev, cur = cur, prev
	}
	lcs := float64(prev[len(r)])
	if lcs == 0 {
		return 0
	}
	precision, recall := lcs/float64(len(p)), lcs/float64(len(r))
	return 2 * precision * recall / (precision + recall)
}
