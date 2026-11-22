// Package semanticcache simulates a semantic cache over a stream of requests.
// The cache policy is passed in as two functions, so GPTCache-style static thresholds,
// vCache and our own policies all run through the same loop.
package semanticcache

import (
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/metrics"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/vector"
)

// Query is one request: unit-length embedding plus what is needed to judge a cache hit.
type Query struct {
	ID       int
	Prompt   string
	Vec      []float32
	Class    int    // equivalence class of the prompt, -1 if unknown
	Response string // the answer the LLM gives for this prompt
}

// Step records what happened to one request.
type Step struct {
	Index      int     `json:"index"`
	ID         int     `json:"id"`
	Neighbor   int     `json:"neighbor"`   // index of the nearest cached request, -1 if cache empty
	Similarity float32 `json:"similarity"` // cosine similarity to the neighbour
	Hit        bool    `json:"hit"`        // cached answer was served
	Reusable   bool    `json:"reusable"`   // the neighbour's answer is correct for this request
	Added      bool    `json:"added"`      // request was inserted into the cache
	CacheSize  int     `json:"cache_size"` // entries in the cache before this request
	Confidence float64 `json:"confidence"` // policy-specific score (e.g. vCache 1-tau), 0 if unused
}

// Decide returns true to serve the cached answer of entry (hit), plus an optional confidence.
type Decide func(entry int, similarity float32) (hit bool, confidence float64)

// Learn is called on a miss with the ground truth; it returns true to insert the request.
type Learn func(entry int, similarity float32, reusable bool) (add bool)

// Reusable tells whether the answer cached for b is correct for a.
func Reusable(a, b Query) bool {
	if a.Class >= 0 && b.Class >= 0 {
		return a.Class == b.Class
	}
	return metrics.ExactMatch(a.Response, b.Response)
}

// Simulate feeds the queries through the cache one by one.
func Simulate(queries []Query, decide Decide, learn Learn) []Step {
	var vecs [][]float32 // cached embeddings
	var owner []int      // cache entry -> query index
	steps := make([]Step, len(queries))
	for i, q := range queries {
		entry, sim := vector.Nearest(q.Vec, vecs)
		s := Step{Index: i, ID: q.ID, Neighbor: -1, Similarity: sim, CacheSize: len(vecs)}
		if entry < 0 {
			s.Similarity, s.Added = 0, true
		} else {
			s.Neighbor = owner[entry]
			s.Reusable = Reusable(q, queries[owner[entry]])
			s.Hit, s.Confidence = decide(entry, sim)
			s.Added = !s.Hit && learn(entry, sim, s.Reusable)
		}
		if s.Added {
			vecs, owner = append(vecs, q.Vec), append(owner, i)
		}
		steps[i] = s
	}
	return steps
}

// Static is the GPTCache policy: hit when similarity >= threshold, insert every miss.
func Static(threshold float32) (Decide, Learn) {
	return func(_ int, sim float32) (bool, float64) { return sim >= threshold, 0 },
		func(int, float32, bool) bool { return true }
}

// Result summarises a simulation.
type Result struct {
	Requests   int     `json:"requests"`
	Hits       int     `json:"hits"`
	WrongHits  int     `json:"wrong_hits"`
	HitRate    float64 `json:"hit_rate"`    // hits / requests
	ErrorRate  float64 `json:"error_rate"`  // wrong hits / requests (vCache definition)
	Precision  float64 `json:"precision"`   // correct hits / hits
	Recall     float64 `json:"recall"`      // correct hits / requests whose neighbour was reusable
	F1         float64 `json:"f1"`          // of precision and recall
	Reusable   int     `json:"reusable"`    // requests whose nearest neighbour was reusable
	MissedHits int     `json:"missed_hits"` // reusable neighbour but answered by the LLM
	CacheSize  int     `json:"cache_size"`  // entries at the end
	LLMCalls   int     `json:"llm_calls"`
}

// Summarize aggregates steps into a Result.
func Summarize(steps []Step) Result {
	r := Result{Requests: len(steps)}
	for _, s := range steps {
		switch {
		case s.Hit && s.Reusable:
			r.Hits++
		case s.Hit:
			r.Hits++
			r.WrongHits++
		case s.Reusable:
			r.MissedHits++
		}
		if s.Reusable {
			r.Reusable++
		}
		if s.Added {
			r.CacheSize++
		}
	}
	r.LLMCalls = r.Requests - r.Hits
	r.HitRate, r.ErrorRate = metrics.Rate(r.Hits, r.Requests), metrics.Rate(r.WrongHits, r.Requests)
	r.Precision, r.Recall, r.F1 = metrics.PrecisionRecall(r.Hits-r.WrongHits, r.WrongHits, r.MissedHits)
	return r
}
