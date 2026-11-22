// Benchmark: quantized cache embeddings (float32 vs int8 vs binary vs Matryoshka truncation).
// For every variant: the static-threshold cache curve, how often the nearest neighbour agrees
// with float32, bytes per vector and the time of one brute-force lookup.
//
//	go run ./cmd/bench/embedquant -n 20000
package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/dataset"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/embedquant"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/output"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/semanticcache"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/vector"
)

type run struct {
	Dataset       string  `json:"dataset"`
	Variant       string  `json:"variant"`
	Dimensions    int     `json:"dimensions"`
	BytesPerEntry int     `json:"bytes_per_entry"`
	Threshold     float32 `json:"threshold"`
	NNAgreement   float64 `json:"nn_agreement"` // share of requests with the same nearest neighbour as float32 (static cache, same threshold index)
	LookupMicros  float64 `json:"lookup_us"`    // one scan over the whole cache, single core, native representation
	semanticcache.Result
}

// lookup times one single-core scan of n entries with the given similarity function.
func lookup(n int, sim func(i int) float32) float64 {
	start := time.Now()
	const rounds = 20
	var best float32
	for range rounds {
		for i := range n {
			best = max(best, sim(i))
		}
	}
	_ = best
	return float64(time.Since(start).Microseconds()) / rounds
}

func main() {
	name := flag.String("dataset", "lmarena-openai", "vCache benchmark with Matryoshka embeddings")
	n := flag.Int("n", 20000, "use only the first n requests (0 = all)")
	flag.Parse()
	queries, err := dataset.VCache(*name)
	if err != nil {
		log.Fatal(err)
	}
	if *n > 0 && *n < len(queries) {
		queries = queries[:*n]
	}
	full := make([][]float32, len(queries))
	for i, q := range queries {
		full[i] = q.Vec
	}
	dim := len(full[0])
	scale := embedquant.Scale(full[:min(1000, len(full))])

	type variant struct {
		name   string
		dim    int
		bytes  int
		vecs   [][]float32 // decoded vectors for the simulator
		micros float64
	}
	convert := func(f func([]float32) []float32) [][]float32 {
		out := make([][]float32, len(full))
		for i, v := range full {
			out[i] = f(v)
		}
		return out
	}
	int8s := make([][]int8, len(full))
	bins := make([][]uint64, len(full))
	for i, v := range full {
		int8s[i], bins[i] = embedquant.Int8(v, scale), embedquant.Binary(v)
	}
	variants := []variant{
		{"float32", dim, dim * 4, full, lookup(len(full), func(i int) float32 { return vector.Dot(full[0], full[i]) })},
		{"int8", dim, dim, convert(func(v []float32) []float32 { return embedquant.DecodeInt8(embedquant.Int8(v, scale)) }),
			lookup(len(full), func(i int) float32 { return float32(embedquant.DotInt8(int8s[0], int8s[i])) })},
		{"binary", dim, dim / 8, convert(func(v []float32) []float32 { return embedquant.DecodeBinary(embedquant.Binary(v), dim) }),
			lookup(len(full), func(i int) float32 { return embedquant.CosineBinary(bins[0], bins[i], dim) })},
	}
	for _, d := range []int{768, 512, 256, 128, 64} {
		t := convert(func(v []float32) []float32 { return embedquant.Truncate(v, d) })
		variants = append(variants, variant{fmt.Sprintf("truncate-%d", d), d, d * 4, t, lookup(len(t), func(i int) float32 { return vector.Dot(t[0], t[i]) })})
		q := convert(func(v []float32) []float32 {
			return embedquant.DecodeInt8(embedquant.Int8(embedquant.Truncate(v, d), scale*2))
		})
		variants = append(variants, variant{fmt.Sprintf("truncate-%d-int8", d), d, d, q, 0})
	}

	dir := "output/embedquant/" + *name + "/"
	thresholds := []float32{0.50, 0.55, 0.60, 0.65, 0.70, 0.75, 0.80, 0.84, 0.88, 0.90, 0.92, 0.94, 0.96, 0.98}
	reference := map[float32][]semanticcache.Step{}
	var runs []run
	for _, v := range variants {
		qs := append([]semanticcache.Query(nil), queries...)
		for i := range qs {
			qs[i].Vec = v.vecs[i]
		}
		for _, t := range thresholds {
			decide, learn := semanticcache.Static(t)
			steps := semanticcache.Simulate(qs, decide, learn)
			if v.name == "float32" {
				reference[t] = steps
			}
			same := 0
			for i, s := range steps {
				if s.Neighbor == reference[t][i].Neighbor {
					same++
				}
			}
			r := run{*name, v.name, v.dim, v.bytes, t, float64(same) / float64(len(steps)), v.micros, semanticcache.Summarize(steps)}
			runs = append(runs, r)
			fmt.Printf("%-18s bytes %5d  threshold %.2f  hit %.4f  error %.4f  nn-agree %.3f  lookup %.0fus\n", v.name, v.bytes, t, r.HitRate, r.ErrorRate, r.NNAgreement, v.micros)
			if err := output.JSONL(fmt.Sprintf("%ssteps_%s_%.2f.jsonl", dir, v.name, t), steps); err != nil {
				log.Fatal(err)
			}
		}
	}
	if err := output.JSON(dir+"summary.json", runs); err != nil {
		log.Fatal(err)
	}
}
