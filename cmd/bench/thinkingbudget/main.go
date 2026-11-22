// Benchmark: thinking-budget control on a reasoning model.
//
//	budget-N  fixed cap of N thinking tokens (budget-0 = thinking off; unlimited = baseline run)
//	estimate  TALE-style: the model estimates its own budget first
//	route     "When to Reason"-style: a small model decides thinking on/off
//
//	make bench-thinkingbudget
package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/bench"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/dataset"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/thinkingbudget"
)

func main() {
	router := flag.String("router-url", "http://localhost:8102", "small model used by the route strategy")
	c := bench.Flags("thinkingbudget", "gsm8k,math500,mmlupro", true)

	strategies := map[string]func(context.Context, dataset.Example) thinkingbudget.Decision{
		"estimate": func(ctx context.Context, e dataset.Example) thinkingbudget.Decision {
			return thinkingbudget.Estimate(ctx, c.URL, e.Prompt)
		},
		"route": func(ctx context.Context, e dataset.Example) thinkingbudget.Decision {
			return thinkingbudget.Route(ctx, *router, e.Prompt)
		},
	}
	for _, b := range []int{0, 128, 256, 512, 1024, 2048} {
		strategies[fmt.Sprint("budget-", b)] = func(context.Context, dataset.Example) thinkingbudget.Decision { return thinkingbudget.Fixed(b) }
	}
	for _, name := range []string{"budget-0", "budget-128", "budget-256", "budget-512", "budget-1024", "budget-2048", "estimate", "route"} {
		c.Tag = name
		c.Ask = func(ctx context.Context, c bench.Config, e dataset.Example) bench.Row {
			d := strategies[name](ctx, e)
			c.Extra, c.Think = d.Extra, d.Think
			row := bench.Ask(ctx, c, e)
			row.OverheadSeconds, row.OverheadTokens, row.Note = d.Seconds, d.Tokens, fmt.Sprintf("think=%v budget=%d", d.Think, d.Budget)
			return row
		}
		bench.Run(c)
	}
}
