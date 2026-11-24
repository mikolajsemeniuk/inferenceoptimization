// Tool: print the thinking-budget decisions (estimate, route) for a few examples without answering them.
//
//	go run ./cmd/tools/probe -tasks gsm8k,math500,mmlupro -n 20
package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/dataset"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/thinkingbudget"
)

func main() {
	url := flag.String("url", "http://localhost:8101", "model that estimates its budget")
	tasks := flag.String("tasks", "gsm8k,math500,mmlupro", "tasks")
	n := flag.Int("n", 20, "examples per task")
	flag.Parse()
	ctx := context.Background()
	for _, task := range strings.Split(*tasks, ",") {
		examples, err := dataset.Load(task)
		if err != nil {
			panic(err)
		}
		for _, e := range dataset.Sample(examples, *n, 1) {
			est, route := thinkingbudget.Estimate(ctx, *url, e.Prompt), thinkingbudget.Route(ctx, *url, e.Prompt, 150)
			fmt.Printf("%-8s estimate: think=%-5v budget=%-5d route: think=%v\n", task, est.Think, est.Budget, route.Think)
		}
	}
}
