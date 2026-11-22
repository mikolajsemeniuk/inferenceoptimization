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
