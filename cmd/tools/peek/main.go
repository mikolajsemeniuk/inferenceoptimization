// Prints one example of every task: a quick check that datasets are downloaded and parsed.
package main

import (
	"fmt"
	"os"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/dataset"
)

func main() {
	tasks := []string{"gsm8k", "math500", "mmlu", "mmlupro", "cnndm", "humaneval", "longbench-hotpotqa", "longbench-2wikimqa", "longbench-musique", "longbench-multi_news", "longbench-samsum"}
	if len(os.Args) > 1 {
		tasks = os.Args[1:]
	}
	for _, t := range tasks {
		ex, err := dataset.Load(t)
		if err != nil {
			fmt.Println(t, "ERROR", err)
			continue
		}
		e := dataset.Sample(ex, 1, 1)[0]
		text := e.Text()
		if len(text) > 300 {
			text = text[:150] + " ... " + text[len(text)-150:]
		}
		fmt.Printf("== %s  n=%d  kind=%s  answers=%q  chars=%d\n%s\n\n", t, len(ex), e.Kind, e.Answers, len(e.Text()), text)
	}
}
