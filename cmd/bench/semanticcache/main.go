// Benchmark: GPTCache-style semantic cache with a static similarity threshold.
// Replays a vCache benchmark (precomputed embeddings + responses) for a sweep of thresholds.
//
//	go run ./cmd/bench/semanticcache -dataset lmarena
package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/dataset"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/output"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/semanticcache"
)

type run struct {
	Dataset   string  `json:"dataset"`
	Policy    string  `json:"policy"`
	Threshold float32 `json:"threshold"`
	Seconds   float64 `json:"seconds"`
	semanticcache.Result
}

func main() {
	name := flag.String("dataset", "lmarena", "lmarena | classification | search")
	n := flag.Int("n", 0, "use only the first n requests (0 = all)")
	flag.Parse()

	queries, err := dataset.VCache(*name)
	if err != nil {
		log.Fatal(err)
	}
	if *n > 0 && *n < len(queries) {
		queries = queries[:*n]
	}
	dir := "output/semanticcache/" + *name + "/"
	var runs []run
	for _, t := range []float32{0.70, 0.75, 0.80, 0.84, 0.86, 0.88, 0.90, 0.92, 0.94, 0.95, 0.96, 0.97, 0.98, 0.99} {
		start := time.Now()
		decide, learn := semanticcache.Static(t)
		steps := semanticcache.Simulate(queries, decide, learn)
		r := run{*name, "static", t, time.Since(start).Seconds(), semanticcache.Summarize(steps)}
		runs = append(runs, r)
		fmt.Printf("threshold %.2f  hit %.4f  error %.4f  precision %.4f  cache %d  (%.0fs)\n", t, r.HitRate, r.ErrorRate, r.Precision, r.CacheSize, r.Seconds)
		if err := output.JSONL(fmt.Sprintf("%ssteps_threshold_%.2f.jsonl", dir, t), steps); err != nil {
			log.Fatal(err)
		}
	}
	if err := output.JSON(dir+"summary.json", runs); err != nil {
		log.Fatal(err)
	}
}
