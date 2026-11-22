// Benchmark: KV-cache quantization (llama-server --cache-type-k/--cache-type-v f16 | q8_0 | q4_0)
// on long-context tasks, where the KV cache dominates memory.
//
//	make bench-kvcachequant
package main

import "github.com/mikolajsemeniuk/inferenceoptimization/pkg/bench"

func main() {
	bench.Run(bench.Flags("kvcachequant", "longbench-hotpotqa,longbench-2wikimqa,longbench-musique", false))
}
