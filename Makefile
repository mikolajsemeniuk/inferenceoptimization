# Reproducible setup for inferenceoptimizationbench. Only needs curl (+ unzip for LongBench).
HF   := https://huggingface.co/datasets
DS   := datasets
CURL := curl -fL --retry 3 -C - --create-dirs

# <hf org>/<hf repo>/<file path in repo>  -> saved as datasets/<hf org>/<hf repo>/<file path>
DATASETS_SMALL := \
	nyu-mll/glue/qqp/validation-00000-of-00001.parquet \
	openai/gsm8k/main/test-00000-of-00001.parquet \
	HuggingFaceH4/MATH-500/test.jsonl \
	cais/mmlu/all/test-00000-of-00001.parquet \
	cais/mmlu/all/dev-00000-of-00001.parquet \
	TIGER-Lab/MMLU-Pro/data/test-00000-of-00001.parquet \
	hotpotqa/hotpot_qa/distractor/validation-00000-of-00001.parquet \
	abisee/cnn_dailymail/3.0.0/test-00000-of-00001.parquet \
	openai/openai_humaneval/openai_humaneval/test-00000-of-00001.parquet \
	zai-org/LongBench/data.zip

# vCache benchmarks (ship precomputed embeddings + LLM responses), ~13 GB total
DATASETS_LARGE := \
	vCache/SemBenchmarkLmArena/train.parquet \
	vCache/SemBenchmarkClassification/train.parquet \
	vCache/SemBenchmarkSearchQueries/train.parquet

define fetch
	@for f in $(1); do \
		repo=`echo $$f | cut -d/ -f1-2`; path=`echo $$f | cut -d/ -f3-`; \
		echo ">> $$f"; \
		$(CURL) -sS -o $(DS)/$$f $(HF)/$$repo/resolve/main/$$path || exit 1; \
	done
endef

MODELS := \
	unsloth/Qwen3-8B-GGUF/Qwen3-8B-Q8_0.gguf \
	unsloth/Qwen3-8B-GGUF/Qwen3-8B-Q6_K.gguf \
	unsloth/Qwen3-8B-GGUF/Qwen3-8B-Q4_K_M.gguf \
	unsloth/Qwen3-8B-GGUF/Qwen3-8B-Q3_K_M.gguf \
	unsloth/Qwen3-8B-GGUF/Qwen3-8B-Q2_K.gguf \
	unsloth/Qwen3-1.7B-GGUF/Qwen3-1.7B-Q8_0.gguf \
	unsloth/Qwen3-0.6B-GGUF/Qwen3-0.6B-Q8_0.gguf \
	gpustack/bge-m3-GGUF/bge-m3-Q8_0.gguf

.PHONY: models
models: ## download GGUF models (~31 GB) into models/
	@for f in $(MODELS); do \
		repo=`echo $$f | cut -d/ -f1-2`; file=`echo $$f | cut -d/ -f3`; \
		echo ">> $$f"; \
		$(CURL) -sS -o models/$$file https://huggingface.co/$$repo/resolve/main/$$file || exit 1; \
	done

.PHONY: datasets datasets-large datasets-all datasets-check

datasets: ## download small evaluation datasets (~0.3 GB)
	$(call fetch,$(DATASETS_SMALL))
	@cd $(DS)/zai-org/LongBench && unzip -qo data.zip

datasets-large: ## download vCache semantic-cache benchmarks (~13 GB)
	$(call fetch,$(DATASETS_LARGE))

datasets-all: datasets datasets-large

datasets-check: ## list downloaded dataset files
	@find $(DS) -type f \( -name '*.parquet' -o -name '*.jsonl' \) | sort | xargs du -h

# ---------------------------------------------------------------- infrastructure
# Servers: llm :8101 (MODEL, CTX, NP, EXTRA), llm2 :8102 (MODEL2, ...), embed :8103.
#   make up MODEL=Qwen3-8B-Q4_K_M.gguf EXTRA="-ctk q8_0 -ctv q8_0"
SERVICES ?= llm
export UID_GID := $(shell id -u):$(shell id -g)

.PHONY: image up down ready
image: ## build the llama.cpp server image
	docker compose build llm

up: ## start SERVICES and wait until they answer
	@mkdir -p slots
	docker compose up -d --force-recreate $(SERVICES)
	@$(MAKE) --no-print-directory ready

down: ## stop all servers
	docker compose down

ready: ## wait for /health of every running server, then send one real request
	@for s in $(SERVICES); do \
		case $$s in llm) p=8101;; llm2) p=8102;; embed) p=8103;; esac; \
		printf "waiting for $$s on :$$p "; \
		for i in `seq 300`; do curl -sf localhost:$$p/health >/dev/null && break; printf .; sleep 1; done; \
		curl -sf localhost:$$p/health >/dev/null || { echo " NOT READY"; docker compose logs --tail 20 $$s; exit 1; }; \
		if [ $$s = embed ]; then \
			curl -sf localhost:$$p/v1/embeddings -d '{"input":"ping"}' | grep -q embedding && echo " ok"; \
		else \
			curl -sf localhost:$$p/v1/chat/completions -d '{"messages":[{"role":"user","content":"ping"}],"max_tokens":1}' | grep -q choices && echo " ok"; \
		fi || { echo " REQUEST FAILED"; exit 1; }; \
	done

# ---------------------------------------------------------------- benchmarks
# Each target writes output/<name>/... ; long runs: nohup make bench-<name> > logs/<name>.log 2>&1 &
CACHE_SETS ?= lmarena classification search

.PHONY: build bench-semanticcache bench-vcache
build:
	go build -o bin/ ./cmd/...

bench-semanticcache: build ## static-threshold semantic cache on the vCache benchmarks (no LLM needed)
	@for d in $(CACHE_SETS); do ./bin/semanticcache -dataset $$d || exit 1; done

bench-vcache: build ## vCache policy on the vCache benchmarks (no LLM needed)
	@for d in $(CACHE_SETS); do ./bin/vcache -dataset $$d || exit 1; done

REASONING := gsm8k,math500,mmlu,mmlupro
LONG      := longbench-hotpotqa,longbench-2wikimqa,longbench-musique
REF       := Qwen3-8B-Q8_0.gguf
# Thinking runs: Qwen3 often reasons for several thousand tokens, a 4096 cap truncated up to 40%
# of MATH-500 answers. 12 slots x 21k context, generation cap 16384.
THINK_UP  := NP=12 CTX=262144
THINK_RUN := -workers 12 -max-tokens 16384

.PHONY: bench-baseline-think
bench-baseline-think: build ## only the thinking half of the baseline
	$(MAKE) up MODEL=$(REF) $(THINK_UP)
	./bin/baseline -tag Q8_0 $(THINK_RUN) -think=true -tasks $(REASONING)

.PHONY: bench-baseline
bench-baseline: build ## reference model without any optimization, thinking on and off
	$(MAKE) up MODEL=$(REF) $(THINK_UP)
	./bin/baseline -tag Q8_0 $(THINK_RUN) -think=true -tasks $(REASONING)
	./bin/baseline -tag Q8_0 -workers 12 -think=false -tasks $(REASONING),cnndm
	$(MAKE) up MODEL=$(REF) NP=4 CTX=131072
	./bin/baseline -tag Q8_0-long -workers 4 -n 50 -think=false -tasks $(LONG)

# serverinfo <benchmark> <tag>: store GPU memory of the running server next to the results
define serverinfo
	@mkdir -p output/$(1)/$(2)
	@nvidia-smi --query-compute-apps=process_name,used_memory --format=csv,noheader,nounits | grep /app/llama-server | \
		awk -F', ' '{s=s sep $$2; sep=","} END {print "{\"benchmark\":\"$(1)\",\"tag\":\"$(2)\",\"gpu_memory_mib\":[" s "]}"}' > output/$(1)/$(2)/server.json
	@cat output/$(1)/$(2)/server.json
endef

QUANTS ?= Q6_K Q4_K_M Q3_K_M Q2_K
.PHONY: bench-quantization
bench-quantization: build ## Qwen3-8B at lower GGUF precisions (Q8_0 is the baseline)
	@for q in $(QUANTS); do \
		$(MAKE) --no-print-directory up MODEL=Qwen3-8B-$$q.gguf NP=16 CTX=131072 || exit 1; \
		./bin/quantization -tag $$q -workers 16 -think=false || exit 1; \
		$(MAKE) --no-print-directory up MODEL=Qwen3-8B-$$q.gguf $(THINK_UP) || exit 1; \
		./bin/quantization -tag $$q $(THINK_RUN) -think=true -n 100 -tasks gsm8k,mmlupro || exit 1; \
		$(MAKE) --no-print-directory info NAME=quantization TAG=$$q; \
	done

KVTYPES ?= f16 q8_0 q4_0
.PHONY: bench-kvcachequant
bench-kvcachequant: build ## KV cache stored as f16 / q8_0 / q4_0, long-context tasks
	@for t in $(KVTYPES); do \
		$(MAKE) --no-print-directory up MODEL=$(REF) NP=4 CTX=131072 EXTRA="-ctk $$t -ctv $$t" || exit 1; \
		./bin/kvcachequant -tag $$t -workers 4 -n 50 || exit 1; \
		./bin/kvcachequant -tag $$t -workers 4 -n 100 -tasks gsm8k,mmlu || exit 1; \
		$(MAKE) --no-print-directory info NAME=kvcachequant TAG=$$t; \
	done

.PHONY: info
info:
	$(call serverinfo,$(NAME),$(TAG))

# speculative decoding configurations: <tag>@<server flags, + instead of spaces>
D06 := -md+/models/Qwen3-0.6B-Q8_0.gguf+--spec-type+draft-simple+--spec-draft-n-max
D17 := -md+/models/Qwen3-1.7B-Q8_0.gguf+--spec-type+draft-simple+--spec-draft-n-max
SPEC := none@ draft0.6B-n8@$(D06)+8 draft0.6B-n16@$(D06)+16 draft1.7B-n8@$(D17)+8 draft1.7B-n16@$(D17)+16 ngram-simple@--spec-type+ngram-simple
.PHONY: bench-specdecode
bench-specdecode: build ## draft-model and n-gram speculative decoding, one request at a time
	@for s in $(SPEC); do \
		tag=`echo $$s | cut -d@ -f1`; flags=`echo $$s | cut -d@ -f2 | tr + ' '`; \
		$(MAKE) --no-print-directory up MODEL=$(REF) NP=1 CTX=16384 EXTRA="$$flags" || exit 1; \
		./bin/specdecode -tag $$tag -workers 1 -n 40 -max-tokens 512 -think=false || exit 1; \
		./bin/specdecode -tag $$tag -workers 1 -n 15 -max-tokens 2048 -think=true -tasks gsm8k || exit 1; \
		$(MAKE) --no-print-directory info NAME=specdecode TAG=$$tag; \
	done

.PHONY: bench-prefixcache
bench-prefixcache: build ## prompt prefix reuse in memory and from disk
	$(MAKE) --no-print-directory up MODEL=$(REF) NP=1 CTX=32768
	./bin/prefixcache

.PHONY: bench-thinkingbudget
bench-thinkingbudget: build ## fixed budgets, self-estimated budget, small-model router
	$(MAKE) --no-print-directory up MODEL=$(REF) $(THINK_UP)
	./bin/thinkingbudget $(THINK_RUN)

CHEAP ?= Qwen3-8B-Q3_K_M Qwen3-1.7B-Q8_0
.PHONY: bench-cascade
bench-cascade: build ## cheap model first, strong model (baseline answers) on low confidence
	@for m in $(CHEAP); do \
		$(MAKE) --no-print-directory up SERVICES=llm2 MODEL2=$$m.gguf NP2=16 CTX2=65536 || exit 1; \
		./bin/cascade -tag $$m -workers 16 || exit 1; \
	done

.PHONY: bench-gpu
bench-gpu: bench-prefixcache bench-specdecode bench-kvcachequant bench-cascade bench-thinkingbudget bench-quantization ## all GPU benchmarks after the baseline, in order

.PHONY: bench-tracecompress bench-promptcompress
bench-tracecompress: build ## how much of a reasoning trace must be kept for it to stay useful
	$(MAKE) --no-print-directory up SERVICES="llm llm2" MODEL=$(REF) NP=8 CTX=262144 MODEL2=Qwen3-1.7B-Q8_0.gguf NP2=8 CTX2=262144
	./bin/tracecompress -workers 8 -max-tokens 4096

bench-promptcompress: build ## extractive and small-LLM context compression on long-context QA
	$(MAKE) --no-print-directory up SERVICES="llm llm2 embed" MODEL=$(REF) NP=4 CTX=131072 MODEL2=Qwen3-1.7B-Q8_0.gguf NP2=2 CTX2=65536
	./bin/promptcompress -workers 4 -n 50 -max-tokens 64

.PHONY: bench-cachehint
bench-cachehint: build ## cached neighbour used as an in-context hint and as a speculative draft
	$(MAKE) --no-print-directory up SERVICES="llm embed" MODEL=$(REF) $(THINK_UP)
	./bin/cachehint -tag plain $(THINK_RUN)
	$(MAKE) --no-print-directory up SERVICES="llm embed" MODEL=$(REF) NP=1 CTX=16384
	./bin/cachehint -tag plain-1slot -workers 1 -n 40 -tasks gsm8k -modes nohint-nothink,hint-qa-nothink
	$(MAKE) --no-print-directory up SERVICES="llm embed" MODEL=$(REF) NP=1 CTX=16384 EXTRA="--spec-type ngram-simple --spec-ngram-simple-size-n 4 --spec-ngram-simple-size-m 32"
	./bin/cachehint -tag ngram-1slot -workers 1 -n 40 -tasks gsm8k -modes nohint-nothink,hint-qa-nothink
