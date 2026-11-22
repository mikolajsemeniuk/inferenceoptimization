// Package vcache is a Go port of the vCache verified decision policy
// (vcache/vcache_policy/strategies/verified.py in github.com/vcache-project/vCache).
// Every cache entry keeps (similarity, was-reusable) observations; a logistic curve fitted
// to them gives the probability that a new request at similarity s can reuse the entry,
// and the entry is only exploited when that keeps the error rate below delta.
package vcache

import (
	"math"
	"math/rand/v2"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/semanticcache"
)

// Observation is one explored request against a cache entry.
type Observation struct {
	Similarity float64
	Reusable   bool
}

// variance of the threshold estimate under perfect separation, by number of observations (from the authors' code)
var separatedVariance = map[int]float64{
	6: 0.035445, 7: 0.028285, 8: 0.026436, 9: 0.021349, 10: 0.019371, 11: 0.012615, 12: 0.011433, 13: 0.010228,
	14: 0.009963, 15: 0.009253, 16: 0.011674, 17: 0.013015, 18: 0.010897, 19: 0.011841, 20: 0.013081, 21: 0.010585,
	22: 0.014255, 23: 0.012058, 24: 0.013002, 25: 0.011715, 26: 0.00839, 27: 0.008839, 28: 0.010628, 29: 0.009899,
	30: 0.008033, 31: 0.00457, 32: 0.007335, 33: 0.008932, 34: 0.00729, 35: 0.007445, 36: 0.00761, 37: 0.011423,
	38: 0.011233, 39: 0.006783, 40: 0.005233, 41: 0.00872, 42: 0.010005, 43: 0.01199, 44: 0.00977, 45: 0.01891,
	46: 0.01513, 47: 0.02109, 48: 0.01531,
}

func sigmoid(z float64) float64 { return 1 / (1 + math.Exp(-z)) }

func round(x float64, digits int) float64 {
	p := math.Pow(10, float64(digits))
	return math.Round(x*p) / p
}

// Fit estimates the entry's decision threshold tHat, the steepness gamma and the variance of tHat.
// ok is false when there is too little evidence (fewer than 6 observations or only one label).
func Fit(obs []Observation) (tHat, gamma, varT float64, ok bool) {
	pos, neg := 0, 0
	minPos, maxNeg := math.Inf(1), math.Inf(-1)
	for _, o := range obs {
		if o.Reusable {
			pos++
			minPos = min(minPos, o.Similarity)
		} else {
			neg++
			maxNeg = max(maxNeg, o.Similarity)
		}
	}
	if len(obs) < 6 || pos == 0 || neg == 0 {
		return 0, 0, 0, false
	}
	// unpenalised logistic regression p = sigmoid(b0 + b1*s), damped Newton steps
	var b0, b1 float64
	var h00, h01, h11 float64
	for range 100 {
		var g0, g1 float64
		h00, h01, h11 = 1e-9, 0, 1e-9
		for _, o := range obs {
			p := sigmoid(b0 + b1*o.Similarity)
			y := 0.0
			if o.Reusable {
				y = 1
			}
			w := p * (1 - p)
			g0, g1 = g0+(y-p), g1+(y-p)*o.Similarity
			h00, h01, h11 = h00+w, h01+w*o.Similarity, h11+w*o.Similarity*o.Similarity
		}
		det := h00*h11 - h01*h01
		d0, d1 := (h11*g0-h01*g1)/det, (h00*g1-h01*g0)/det
		if n := math.Hypot(d0, d1); n > 10 { // keeps separated data from overflowing
			d0, d1 = d0*10/n, d1*10/n
		}
		b0, b1 = b0+d0, b1+d1
		if math.Hypot(g0, g1) < 1e-8 {
			break
		}
	}
	g := max(b1, 1e-6)
	tHat = min(max(-b0/g, 0), 1)
	if minPos > maxNeg { // perfect separation: Hessian is degenerate, use the tabulated variance
		varT = separatedVariance[min(len(obs), 48)]
	} else { // delta method on t = -b0/b1
		det := h00*h11 - h01*h01
		c00, c01, c11 := h11/det, -h01/det, h00/det
		a, b := -1/g, b0/(g*g)
		varT = max(0, a*a*c00+2*a*b*c01+b*b*c11)
	}
	return round(tHat, 3), round(g, 3), varT, true
}

// Tau is the exploration probability for a request at similarity s (1 = always ask the LLM).
func Tau(s, tHat, gamma, varT, delta float64) float64 {
	tau := math.Inf(1)
	for i := range 50 {
		eps := 1e-6 + float64(i)*(1-2e-6)/49
		z := math.Sqrt2 * math.Erfinv(2*(1-eps)-1) // normal quantile of 1-eps
		t := min(max(tHat+z*math.Sqrt(varT), 0), 1)
		alpha := (1 - eps) * sigmoid(gamma*(s-t))
		tau = min(tau, 1-delta/(1-alpha))
	}
	return round(tau, 5)
}

// Policy returns the vCache decide/learn pair for semanticcache.Simulate.
// delta is the user's maximum error rate; rng makes runs repeatable.
func Policy(delta float64, rng *rand.Rand) (semanticcache.Decide, semanticcache.Learn) {
	var observations [][]Observation // per cache entry
	entry := func(i int) []Observation {
		for len(observations) <= i {
			observations = append(observations, nil)
		}
		return observations[i]
	}
	decide := func(e int, sim float32) (bool, float64) {
		s := round(float64(sim), 3)
		tHat, gamma, varT, ok := Fit(entry(e))
		if !ok {
			return false, 0
		}
		tau := Tau(s, tHat, gamma, varT, delta)
		return rng.Float64() > tau, 1 - tau
	}
	learn := func(e int, sim float32, reusable bool) bool {
		entry(e)
		observations[e] = append(observations[e], Observation{round(float64(sim), 3), reusable})
		return !reusable // only requests the entry cannot answer become new entries
	}
	return decide, learn
}
