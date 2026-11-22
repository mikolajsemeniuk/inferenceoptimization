# Milestones

Stack: Go (harness, gateway, caches) + llama.cpp `llama-server` in Docker (GGUF models, OpenAI-compatible API) on ASUS Ascent GX10 (GB10, arm64, 128 GB unified memory). No Python, no training/fine-tuning.
Every experiment writes `output/<name>/results.jsonl` (one row per request) and `output/<name>/summary.json` (accuracy, hit rate, error rate, TTFT, total latency, tokens/s, prompt/completion/thinking tokens, peak memory).

## M0 - infrastructure
* run `make datasets` (done, ~0.3 GB) and `make datasets-large` (vCache benchmarks, ~13 GB - free disk space first, host has ~98 GB left)
* choose models and add `make models` (curl GGUF from Hugging Face into `models/`): one reasoning-capable target in several quantizations (e.g. Qwen3 14B/32B at Q8_0, Q4_K_M, Q2_K), one small draft model of the same family (0.6B-1.7B), one embedding model (e.g. bge-m3 or Qwen3-Embedding GGUF)
* create docker-compose.yaml with services: `llm-target` (llama-server, full precision / Q8), `llm-quant` (same model, Q4/Q2), `llm-spec` (target + `--model-draft`), `embed` (llama-server `--embeddings`); verify the arm64 CUDA image `ghcr.io/ggml-org/llama.cpp:server-cuda` runs on GB10, otherwise add a Dockerfile that builds llama.cpp from source
* extend Makefile: `make up`, `make down`, `make ready` (curl `/health` and one `/v1/chat/completions` + `/v1/embeddings` call per service), `make bench-<name>`, `make bench-all`
* create go.mod
* create pkg/llm/*.go (OpenAI-compatible client: chat, completion, embeddings, streaming with TTFT, token usage, thinking on/off and budget)
* create pkg/dataset/*.go (parquet/jsonl loaders for every dataset in `datasets/`, fixed-seed sampling)
* create pkg/eval/*.go (exact match, multiple-choice accuracy, numeric answer extraction, F1, ROUGE-L, LLM-judge equivalence for cache hits)
* create pkg/metrics/*.go (latency percentiles, throughput, tokens, memory, jsonl/json writers)
* create cmd/baseline/main.go
* run cmd/baseline/main.go and check result in output/baseline/*.{json,jsonl} (no optimization, full-precision target, thinking on - reference point for everything below)

## M1 - response-level caching
* create pkg/semanticcache/*.go (GPTCache-style: embedding + in-memory cosine index + static threshold, stores question and answer)
* create cmd/semanticcache/main.go (threshold sweep 0.70-0.99)
* run cmd/semanticcache/main.go and check result in output/semanticcache/*.{json,jsonl} (datasets: SemBenchmarkLmArena, SemBenchmarkClassification, SemBenchmarkSearchQueries, QQP)
* create pkg/vcache/*.go (per-entry online-learned decision boundary with user-defined max error rate, as in vCache)
* create cmd/vcache/main.go
* run cmd/vcache/main.go and check result in output/vcache/*.{json,jsonl} (error-rate vs hit-rate curves against static threshold)
* create pkg/embedquant/*.go (cache embeddings in float32 / int8 / binary / Matryoshka-truncated form)
* create cmd/embedquant/main.go
* run cmd/embedquant/main.go and check result in output/embedquant/*.{json,jsonl} (hit rate, error rate, index memory, lookup latency)

## M2 - model-level optimizations
* create cmd/quantization/main.go (same model at F16/Q8/Q6/Q4/Q3/Q2)
* run cmd/quantization/main.go and check result in output/quantization/*.{json,jsonl} (datasets: MMLU, MMLU-Pro, GSM8K, MATH-500, HumanEval)
* create cmd/kvcachequant/main.go (llama-server `--cache-type-k/--cache-type-v` f16 / q8_0 / q4_0)
* run cmd/kvcachequant/main.go and check result in output/kvcachequant/*.{json,jsonl} (datasets: LongBench, HotpotQA)
* create cmd/prefixcache/main.go (`cache_prompt` on/off, slot save/restore to disk, shared system prompt and shared RAG document prefixes)
* run cmd/prefixcache/main.go and check result in output/prefixcache/*.{json,jsonl} (TTFT on LongBench, HotpotQA)
* create cmd/specdecode/main.go (draft model size and `--draft-max/--draft-min/--draft-p-min` sweep)
* run cmd/specdecode/main.go and check result in output/specdecode/*.{json,jsonl} (acceptance rate, tokens/s on GSM8K, CNN/DailyMail, HumanEval - the Spec-Bench task mix)

## M3 - request-level optimizations
* create pkg/cascade/*.go (FrugalGPT / AutoMix style: quantized model answers first, escalate to full model on low confidence: logprob margin, self-verification, self-consistency)
* create cmd/cascade/main.go
* run cmd/cascade/main.go and check result in output/cascade/*.{json,jsonl} (cost-quality curve on MMLU, MMLU-Pro, GSM8K)
* create pkg/thinkingbudget/*.go (no-think, fixed budgets, TALE-style self-estimated budget, difficulty router as in "When to Reason")
* create cmd/thinkingbudget/main.go
* run cmd/thinkingbudget/main.go and check result in output/thinkingbudget/*.{json,jsonl} (GSM8K, MATH-500, MMLU-Pro)
* create pkg/tracecompress/*.go (reasoning trace compression without training: Chain-of-Draft prompting, small-model summarisation of the trace before it is cached or re-fed)
* create cmd/tracecompress/main.go
* run cmd/tracecompress/main.go and check result in output/tracecompress/*.{json,jsonl}
* create pkg/promptcompress/*.go (training-free compression: Selective-Context / LLMLingua-style token pruning using small-model logprobs; LLMLingua-2 only if its classifier can be served from a container)
* create cmd/promptcompress/main.go
* run cmd/promptcompress/main.go and check result in output/promptcompress/*.{json,jsonl} (LongBench, HotpotQA, CNN/DailyMail; also measure interaction with prefix cache)

## M4 - cache as a draft / hint
* check whether llama-server can accept an externally supplied draft (lookup decoding / prompt-lookup); if not, add a thin patch or a cgo binding (go-llama.cpp) to verify a given token sequence in one batched forward pass
* create pkg/cachehint/*.go (near-miss cache entry used as: (a) speculative draft verified by the target model, REST / SuffixDecoding / LaCache style, (b) in-context hint "similar question + answer")
* create cmd/cachehint/main.go
* run cmd/cachehint/main.go and check result in output/cachehint/*.{json,jsonl}

## M5 - methods that need Python or trained components (decide: Docker-only or cite-only)
* CacheBlend / KV cache fusion: needs selective KV recompute inside the engine; only via the authors' LMCache + vLLM containers, otherwise report numbers from the paper
* Cache-to-Cache (C2C): needs a trained fuser network - out of scope, cite only
* early exit / LayerSkip: needs specially trained checkpoints and a custom decoding loop - out of scope unless a ready container exists, cite only
* GPTQ / AWQ: covered indirectly by GGUF quantization levels; no separate calibration runs

## M6 - aggregate
* create cmd/report/main.go (merge all `output/*/summary.json` into `output/report/{table.md,table.csv,pareto.json}`)
* run cmd/report/main.go and check which optimizations compose and which regions of the quality / latency / cost space are left uncovered

## M7 - novelty
* do a dedicated prior-art search before committing (arXiv, OpenReview, ACL Anthology, Semantic Scholar) - closest known work to exclude: LaCache (cache + draft-model lookahead + speculative decoding), MemBoost (cache + cascade to a stronger model with write-back), When to Reason (router for reasoning on/off), Hybrid Verified Decoding (choosing cache draft vs model draft), vCache (error-bounded hits), Cache-Aware Prompt Compression
* candidate idea (to be validated against prior art and M1-M4 data): a gateway that does not treat the cache as hit/miss but turns the calibrated probability that a cached neighbour is correct (vCache-style estimate) into a graded action - return cached answer / verify cached answer with the quantized model without thinking / answer with the full model using the neighbour's question, answer and compressed reasoning trace as a hint with a proportionally reduced thinking budget / full model with full budget - under a single user-set error bound
* claim to test: near-miss cache entries, which existing caches discard, cut thinking tokens and latency at equal accuracy; expected win on reasoning-heavy and paraphrase-heavy sets (GSM8K / MATH-500 paraphrases, SemBenchmarkLmArena), expected loss or tie on low-redundancy sets (LongBench, HumanEval)
* create pkg/gateway/*.go (OpenAI-compatible reverse proxy composing semantic cache, vCache estimator, thinking-budget controller, quantization cascade, prefix cache)
* create cmd/gateway/main.go
* create cmd/gatewaybench/main.go (gateway vs every single baseline from M1-M4 and vs best pairwise combinations, same workloads)
* run cmd/gatewaybench/main.go and check result in output/gateway/*.{json,jsonl}
* run ablations (each gateway component switched off) and check result in output/gateway-ablation/*.{json,jsonl}

## M8 - paper (later)
* freeze results, tag the repo, write the article
