# Goal

Build a reproducible benchmark of LLM inference optimizations (see `README.md`, `milestones.md`), then add one new method that beats existing ones on at least some axes, then write a paper.

## The novelty: graded cache gateway

### Before (every existing semantic cache)
One binary decision per request:
* similarity >= threshold -> return the cached answer (cost ~0)
* similarity < threshold -> full model, full thinking; the cache contributes nothing

Problems:
* a near-miss (e.g. similarity 0.85) costs exactly as much as a brand-new question, although the cache holds an almost-ready answer and its reasoning
* the threshold is a trade-off: low -> wrong answers served, high -> cache wasted

### After (our gateway)
Estimate `p` = probability that the nearest cached answer is correct for the new request (vCache-style calibrated estimate, not raw similarity). Pick one of four tiers:

| p | action | cost |
|---|---|---|
| very high | return the cached answer | ~0 |
| high | small quantized model, no thinking, checks "does this cached answer fit this question?"; yes -> return, no -> next tier | few tokens |
| medium | full model gets the neighbour (question + answer + compressed reasoning trace) as a hint and a reduced thinking budget (proportional to 1-p) | fraction of full |
| low | full model, full thinking budget | full |

All tier boundaries derive from a single user knob: maximum tolerated error rate.

### Example
Cache: "Jan has 5 apples, ate 2, how many left?" New: "Ola has 7 pears, ate 3, how many left?"
* before: miss, model reasons from scratch (~800 thinking tokens)
* after: tier 3, model sees the solved analogous problem, reasons briefly (~150 thinking tokens)

### What it changes / fixes
* near-misses are no longer wasted (main novelty: existing caches discard them)
* fewer wrong hits: doubtful hits are cheaply verified instead of blindly served
* thinking budget depends on what the cache already knows, not only on question difficulty
* one knob (max error rate) instead of a hand-tuned similarity threshold

### Claim to prove
At equal accuracy, fewer thinking tokens and lower latency than: static-threshold semantic cache, vCache, quantization cascade alone, thinking-budget router alone.
* expected win: paraphrase-heavy and reasoning-heavy workloads (GSM8K / MATH-500 variants, SemBenchmarkLmArena)
* expected tie or loss: low-redundancy workloads (LongBench, HumanEval) - only lookup overhead remains

### Closest prior work (must differ from, must compare against)
* vCache - error-bounded hit/miss, no use of near-misses
* LaCache - cache + draft-model lookahead + speculative decoding
* MemBoost - cache + cascade to a stronger model with write-back
* When to Reason - router for reasoning on/off by query difficulty
* Hybrid Verified Decoding - choosing cache draft vs model draft
Redo the prior-art search before implementing and before writing.

## Rules of the repo
* Go only for harness and methods; inference engine runs in Docker; no Python, no training
* one method = one folder `pkg/<method>/` + one entrypoint `cmd/bench/<method>/main.go`
* metrics are plain functions in `pkg/metrics/`
* minimal, simple, functional-style code (input -> output); structs only for shared connections
* every benchmark writes `output/<method>/*.jsonl` (per request) and `*.json` (summary); log as much per-request data as possible
* after each method: check results against the numbers reported by the original authors, then commit and push
