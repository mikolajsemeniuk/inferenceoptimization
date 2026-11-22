package metrics

import (
	"math"
	"strings"
)

// BM25 scores every document against the query (k1=1.5, b=0.75).
func BM25(query string, documents []string) []float64 {
	const k1, b = 1.5, 0.75
	docs := make([][]string, len(documents))
	df := map[string]int{}
	total := 0
	for i, d := range documents {
		docs[i] = strings.Fields(Normalize(d))
		total += len(docs[i])
		seen := map[string]bool{}
		for _, t := range docs[i] {
			if !seen[t] {
				seen[t] = true
				df[t]++
			}
		}
	}
	scores := make([]float64, len(documents))
	if total == 0 {
		return scores
	}
	avg, n := float64(total)/float64(len(docs)), float64(len(docs))
	for i, d := range docs {
		tf := map[string]int{}
		for _, t := range d {
			tf[t]++
		}
		for _, t := range strings.Fields(Normalize(query)) {
			f := float64(tf[t])
			if f == 0 {
				continue
			}
			idf := math.Log(1 + (n-float64(df[t])+0.5)/(float64(df[t])+0.5))
			scores[i] += idf * f * (k1 + 1) / (f + k1*(1-b+b*float64(len(d))/avg))
		}
	}
	return scores
}
