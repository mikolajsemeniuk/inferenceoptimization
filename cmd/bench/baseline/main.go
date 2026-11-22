// Benchmark: no optimization. Reference model (Qwen3-8B Q8_0), every task, thinking on or off.
// Everything else is compared against these numbers.
//
//	go run ./cmd/bench/baseline -think=true
//	go run ./cmd/bench/baseline -think=false
package main

import "github.com/mikolajsemeniuk/inferenceoptimization/pkg/bench"

func main() {
	bench.Run(bench.Flags("baseline", "gsm8k,math500,mmlu,mmlupro,cnndm,longbench-hotpotqa,longbench-2wikimqa,longbench-musique", true))
}
