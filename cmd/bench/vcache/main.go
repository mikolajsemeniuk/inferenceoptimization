// Benchmark: vCache verified semantic cache (per-entry learned decision boundary, error bound delta).
// Replays a vCache benchmark (precomputed embeddings + responses) for a sweep of deltas.
//
//	go run ./cmd/bench/vcache -dataset lmarena
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"time"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/dataset"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/output"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/semanticcache"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/vcache"
)

type run struct {
	Dataset string  `json:"dataset"`
	Policy  string  `json:"policy"`
	Delta   float64 `json:"delta"`
	Seconds float64 `json:"seconds"`
	semanticcache.Result
}

func main() {
	name := flag.String("dataset", "lmarena", "lmarena | classification | search")
	n := flag.Int("n", 0, "use only the first n requests (0 = all)")
	seed := flag.Uint64("seed", 1, "random seed of the explore/exploit draw")
	flag.Parse()

	queries, err := dataset.VCache(*name)
	if err != nil {
		log.Fatal(err)
	}
	if *n > 0 && *n < len(queries) {
		queries = queries[:*n]
	}
	dir := "output/vcache/" + *name + "/"
	var runs []run
	for _, delta := range []float64{0.005, 0.01, 0.015, 0.02, 0.03, 0.05, 0.07, 0.10} {
		start := time.Now()
		decide, learn := vcache.Policy(delta, rand.New(rand.NewPCG(*seed, *seed)))
		steps := semanticcache.Simulate(queries, decide, learn)
		r := run{*name, "vcache", delta, time.Since(start).Seconds(), semanticcache.Summarize(steps)}
		runs = append(runs, r)
		fmt.Printf("delta %.3f  hit %.4f  error %.4f  precision %.4f  cache %d  (%.0fs)\n", delta, r.HitRate, r.ErrorRate, r.Precision, r.CacheSize, r.Seconds)
		if err := output.JSONL(fmt.Sprintf("%ssteps_delta_%.3f.jsonl", dir, delta), steps); err != nil {
			log.Fatal(err)
		}
	}
	if err := output.JSON(dir+"summary.json", runs); err != nil {
		log.Fatal(err)
	}
}
