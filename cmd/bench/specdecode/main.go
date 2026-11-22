// Benchmark: speculative decoding. A small draft model proposes tokens, the target verifies
// them in one pass; output is unchanged, only speed differs. Run with -workers 1.
//
//	make bench-specdecode
package main

import "github.com/mikolajsemeniuk/inferenceoptimization/pkg/bench"

func main() {
	bench.Run(bench.Flags("specdecode", "gsm8k,cnndm,humaneval", false))
}
