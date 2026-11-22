// Package embedquant shrinks cache embeddings: int8 scalar quantization, 1-bit binary
// quantization and Matryoshka-style truncation. Every function maps a unit-length float32
// vector to a compact form, plus a similarity function that works on that form.
package embedquant

import (
	"math"
	"math/bits"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/vector"
)

// Truncate keeps the first d dimensions and re-normalizes (valid for Matryoshka-trained models).
func Truncate(v []float32, d int) []float32 {
	return vector.Normalize(v[:min(d, len(v))])
}

// Int8 maps each component of a unit vector to [-127,127] using a fixed scale.
// scale should be the largest absolute component expected (see Scale).
func Int8(v []float32, scale float32) []int8 {
	out := make([]int8, len(v))
	for i, x := range v {
		out[i] = int8(math.Round(float64(max(-1, min(1, x/scale))) * 127))
	}
	return out
}

// Scale returns the 99.9th percentile-like bound: mean absolute value times 4, a simple
// calibration that needs one pass over a sample of vectors.
func Scale(sample [][]float32) float32 {
	var sum float64
	n := 0
	for _, v := range sample {
		for _, x := range v {
			sum += math.Abs(float64(x))
			n++
		}
	}
	return float32(4 * sum / float64(max(n, 1)))
}

// DotInt8 is the integer dot product.
func DotInt8(a, b []int8) int32 {
	var s int32
	b = b[:len(a)]
	for i := range a {
		s += int32(a[i]) * int32(b[i])
	}
	return s
}

// CosineInt8 is the cosine similarity of two int8 vectors.
func CosineInt8(a, b []int8) float32 {
	return float32(float64(DotInt8(a, b)) / math.Sqrt(float64(DotInt8(a, a))*float64(DotInt8(b, b))))
}

// Binary keeps only the sign of each component, packed 64 per word.
func Binary(v []float32) []uint64 {
	out := make([]uint64, (len(v)+63)/64)
	for i, x := range v {
		if x > 0 {
			out[i/64] |= 1 << (i % 64)
		}
	}
	return out
}

// CosineBinary is the cosine similarity of two sign vectors of dimension d: 1 - 2*hamming/d.
func CosineBinary(a, b []uint64, d int) float32 {
	h := 0
	for i := range a {
		h += bits.OnesCount64(a[i] ^ b[i])
	}
	return 1 - 2*float32(h)/float32(d)
}

// Decode* turn the compact forms back into unit float32 vectors whose dot product equals the
// compact similarity. The cache simulator uses them, so every variant runs through the same code.

// DecodeInt8 returns the unit float32 vector equivalent to an int8 vector.
func DecodeInt8(q []int8) []float32 {
	out := make([]float32, len(q))
	for i, x := range q {
		out[i] = float32(x)
	}
	return vector.Normalize(out)
}

// DecodeBinary returns the unit float32 vector equivalent to a sign vector of dimension d.
func DecodeBinary(q []uint64, d int) []float32 {
	out := make([]float32, d)
	for i := range out {
		out[i] = -1
		if q[i/64]>>(i%64)&1 == 1 {
			out[i] = 1
		}
	}
	return vector.Normalize(out)
}
