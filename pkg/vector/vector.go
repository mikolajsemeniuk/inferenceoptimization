// Package vector has the few float32 vector helpers the caches need.
package vector

import (
	"math"
	"runtime"
	"sync"
)

// Dot product of two equal-length vectors.
func Dot(a, b []float32) float32 {
	var s0, s1, s2, s3 float32
	n := len(a) &^ 3
	b = b[:len(a)]
	for i := 0; i < n; i += 4 {
		s0 += a[i] * b[i]
		s1 += a[i+1] * b[i+1]
		s2 += a[i+2] * b[i+2]
		s3 += a[i+3] * b[i+3]
	}
	for i := n; i < len(a); i++ {
		s0 += a[i] * b[i]
	}
	return s0 + s1 + s2 + s3
}

// Normalize returns v scaled to unit length, so Dot becomes cosine similarity.
func Normalize(v []float32) []float32 {
	norm := float32(math.Sqrt(float64(Dot(v, v))))
	if norm == 0 {
		return v
	}
	out := make([]float32, len(v))
	for i, x := range v {
		out[i] = x / norm
	}
	return out
}

// Nearest returns the index and similarity of the stored vector closest to q (-1 if store is empty).
// Exact brute-force search, split over all CPU cores.
func Nearest(q []float32, store [][]float32) (int, float32) {
	workers := min(runtime.NumCPU(), len(store)/256+1)
	type hit struct {
		idx int
		sim float32
	}
	hits := make([]hit, workers)
	var wg sync.WaitGroup
	for w := range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			best := hit{-1, -2}
			for i := w * len(store) / workers; i < (w+1)*len(store)/workers; i++ {
				if s := Dot(q, store[i]); s > best.sim {
					best = hit{i, s}
				}
			}
			hits[w] = best
		}()
	}
	wg.Wait()
	best := hit{-1, -2}
	for _, h := range hits {
		if h.sim > best.sim {
			best = h
		}
	}
	return best.idx, best.sim
}
