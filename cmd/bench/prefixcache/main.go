// Benchmark: prefix (KV) caching. Several questions are asked about the same long document:
//
//	cold     - nothing cached, the whole prompt is processed
//	warm     - the document prefix is still in the slot's KV cache (cache_prompt)
//	restored - the KV cache was saved to disk, the slot erased, then restored from disk
//
// Needs a server with one slot: make bench-prefixcache
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/dataset"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/llm"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/metrics"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/output"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/prefixcache"
)

type row struct {
	ID      string                 `json:"id"`
	Task    string                 `json:"task"`
	Mode    string                 `json:"mode"`
	Score   float64                `json:"score"`
	Save    prefixcache.SlotResult `json:"save"`
	Restore prefixcache.SlotResult `json:"restore"`
	llm.Response
}

type summary struct {
	Task            string          `json:"task"`
	Mode            string          `json:"mode"`
	N               int             `json:"n"`
	Score           float64         `json:"score"`
	TTFT            metrics.Summary `json:"ttft_s"`
	Latency         metrics.Summary `json:"latency_s"`
	PromptSeconds   metrics.Summary `json:"prompt_s"`
	PromptTokens    metrics.Summary `json:"prompt_tokens"`
	CachedTokens    metrics.Summary `json:"cached_tokens"`
	SaveSeconds     metrics.Summary `json:"save_s"`
	RestoreSeconds  metrics.Summary `json:"restore_s"`
	FileMegabytes   metrics.Summary `json:"file_mb"`
	TTFTWithRestore metrics.Summary `json:"ttft_plus_restore_s"`
}

func main() {
	url := flag.String("url", "http://localhost:8101", "llama-server base URL (one slot, --slot-save-path set)")
	tasks := flag.String("tasks", "longbench-hotpotqa,longbench-2wikimqa,longbench-musique", "comma-separated long-context tasks")
	n := flag.Int("n", 30, "documents per task")
	flag.Parse()
	ctx := context.Background()

	ask := func(e dataset.Example, cache bool) llm.Response {
		extra := llm.Think(false)
		extra["cache_prompt"], extra["id_slot"] = cache, 0
		res, err := llm.Chat(ctx, *url, llm.Request{Messages: llm.User(e.Text()), MaxTokens: 64, Extra: extra})
		if err != nil {
			log.Fatal(err)
		}
		return res
	}
	must := func(r prefixcache.SlotResult, err error) prefixcache.SlotResult {
		if err != nil {
			log.Fatal(err)
		}
		return r
	}

	var summaries []summary
	for _, task := range strings.Split(*tasks, ",") {
		examples, err := dataset.Load(task)
		if err != nil {
			log.Fatal(err)
		}
		var rows []row
		for _, e := range dataset.Sample(examples, *n, 1) {
			// a second question about the same document: same prefix, different suffix
			other := e
			other.Prompt = "List the titles of the first three passages above, separated by commas."

			must(prefixcache.SlotResult{}, prefixcache.Erase(ctx, *url, 0))
			cold := ask(other, false) // fills the KV cache with the document
			rows = append(rows, row{ID: e.ID, Task: task, Mode: "cold", Response: cold})

			saved := must(prefixcache.Save(ctx, *url, 0, "prefix.bin"))
			warm := ask(e, true)
			rows = append(rows, row{ID: e.ID, Task: task, Mode: "warm", Score: dataset.Grade(e, warm.Content), Response: warm})

			must(prefixcache.SlotResult{}, prefixcache.Erase(ctx, *url, 0))
			restored := must(prefixcache.Restore(ctx, *url, 0, "prefix.bin"))
			disk := ask(e, true)
			rows = append(rows, row{ID: e.ID, Task: task, Mode: "restored", Score: dataset.Grade(e, disk.Content), Save: saved, Restore: restored, Response: disk})

			must(prefixcache.SlotResult{}, prefixcache.Erase(ctx, *url, 0))
			full := ask(e, false) // same question with no cache at all: quality and latency reference
			rows = append(rows, row{ID: e.ID, Task: task, Mode: "nocache", Score: dataset.Grade(e, full.Content), Response: full})
		}
		for _, mode := range []string{"cold", "nocache", "warm", "restored"} {
			s := summary{Task: task, Mode: mode}
			col := map[string][]float64{}
			for _, r := range rows {
				if r.Mode != mode {
					continue
				}
				s.N++
				for k, v := range map[string]float64{"score": r.Score, "ttft": r.TTFT, "latency": r.Latency, "prompt_s": r.PromptSeconds, "prompt": float64(r.PromptTokens),
					"cached": float64(r.CachedTokens), "save": r.Save.Seconds, "restore": r.Restore.Seconds, "mb": float64(r.Save.Bytes) / 1e6, "ttft+restore": r.TTFT + r.Restore.Seconds} {
					col[k] = append(col[k], v)
				}
			}
			s.Score = metrics.Mean(col["score"])
			s.TTFT, s.Latency, s.PromptSeconds = metrics.Summarize(col["ttft"]), metrics.Summarize(col["latency"]), metrics.Summarize(col["prompt_s"])
			s.PromptTokens, s.CachedTokens = metrics.Summarize(col["prompt"]), metrics.Summarize(col["cached"])
			s.SaveSeconds, s.RestoreSeconds, s.FileMegabytes = metrics.Summarize(col["save"]), metrics.Summarize(col["restore"]), metrics.Summarize(col["mb"])
			s.TTFTWithRestore = metrics.Summarize(col["ttft+restore"])
			summaries = append(summaries, s)
			fmt.Printf("%-20s %-9s n=%d score=%.3f ttft=%.2fs prompt=%.0f cached=%.0f restore=%.2fs file=%.0fMB\n", task, mode, s.N, s.Score, s.TTFT.Mean, s.PromptTokens.Mean, s.CachedTokens.Mean, s.RestoreSeconds.Mean, s.FileMegabytes.Mean)
		}
		if err := output.JSONL("output/prefixcache/"+task+".jsonl", rows); err != nil {
			log.Fatal(err)
		}
		if err := output.JSON("output/prefixcache/summary.json", summaries); err != nil {
			log.Fatal(err)
		}
	}
}
