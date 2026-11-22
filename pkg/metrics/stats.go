package metrics

import (
	"math"
	"sort"
)

// Mean of the values (0 for empty input).
func Mean(v []float64) float64 {
	if len(v) == 0 {
		return 0
	}
	return Sum(v) / float64(len(v))
}

// Sum of the values.
func Sum(v []float64) float64 {
	s := 0.0
	for _, x := range v {
		s += x
	}
	return s
}

// Percentile with linear interpolation, p in [0,100].
func Percentile(v []float64, p float64) float64 {
	if len(v) == 0 {
		return 0
	}
	s := append([]float64(nil), v...)
	sort.Float64s(s)
	rank := p / 100 * float64(len(s)-1)
	lo, hi := int(math.Floor(rank)), int(math.Ceil(rank))
	return s[lo] + (s[hi]-s[lo])*(rank-float64(lo))
}

// Summary is the usual set of distribution statistics.
type Summary struct {
	N    int     `json:"n"`
	Mean float64 `json:"mean"`
	P50  float64 `json:"p50"`
	P90  float64 `json:"p90"`
	P99  float64 `json:"p99"`
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
	Sum  float64 `json:"sum"`
}

// Summarize computes Summary for the values.
func Summarize(v []float64) Summary {
	return Summary{len(v), Mean(v), Percentile(v, 50), Percentile(v, 90), Percentile(v, 99), Percentile(v, 0), Percentile(v, 100), Sum(v)}
}
