// Benchmark: weight quantization. The same model (Qwen3-8B) at several GGUF precisions;
// the Makefile restarts the server with each file and passes its name as -tag.
//
//	make bench-quantization
package main

import "github.com/mikolajsemeniuk/inferenceoptimization/pkg/bench"

func main() {
	bench.Run(bench.Flags("quantization", "gsm8k,math500,mmlu,mmlupro,cnndm", false))
}
