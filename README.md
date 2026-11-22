# Inference optimization

## Optimizations
* semantic cache
* speculative decoding
* speculative decoding hint from cache
* setting cache threshold based on other parameters like thinking budget etc
* quantization of model, embedding kv cache
* kv cache
* model, cache, quantization cascade
* prefix caching
* cacheblend
* cache-to-cache C2C
* vCache
* storing question and answer in cache
* kompresja śladów rozumowania
* prompt compression
* early exit/layer


## Bibliography
* [GPTCache: An Open-Source Semantic Cache for LLM Applications Enabling Faster Answers and Cost Savings](https://aclanthology.org/2023.nlposs-1.24/)
* [MeanCache: User-Centric Semantic Caching for LLM Web Services](https://arxiv.org/abs/2403.02694)
* [SCALM: Towards Semantic Caching for Automated Chat Services with Large Language Models](https://arxiv.org/abs/2406.00025)
* [GPT Semantic Cache: Reducing LLM Costs and Latency via Semantic Embedding Caching](https://arxiv.org/abs/2411.05276)
* [vCache: Verified Semantic Prompt Caching](https://arxiv.org/abs/2502.03771)
* [A Generative Caching System for Large Language Models](https://arxiv.org/abs/2503.17603)
* [Semantic Caching for Low-Cost LLM Serving: From Offline Learning to Online Adaptation](https://arxiv.org/abs/2508.07675)
* [Continuous Semantic Caching for Low-Cost LLM Serving](https://arxiv.org/abs/2604.20021)
* [LaCache: Robust Semantic Caching for LLM Serving](https://arxiv.org/abs/2608.01718)
* [MemBoost: A Memory-Boosted Framework for Cost-Aware LLM Inference](https://arxiv.org/abs/2603.26557)
* [Cortex: Achieving Low-Latency, Cost-Efficient Remote Data Access For LLM via Semantic-Aware Knowledge Caching](https://arxiv.org/abs/2509.17360)
* [Fast Inference from Transformers via Speculative Decoding](https://arxiv.org/abs/2211.17192)
* [Accelerating Large Language Model Decoding with Speculative Sampling](https://arxiv.org/abs/2302.01318)
* [SpecInfer: Accelerating Generative Large Language Model Serving with Tree-based Speculative Inference and Verification](https://arxiv.org/abs/2305.09781)
* [REST: Retrieval-Based Speculative Decoding](https://arxiv.org/abs/2311.08252)
* [SuffixDecoding: Extreme Speculative Decoding for Emerging AI Applications](https://arxiv.org/abs/2411.04975)
* [Hybrid Verified Decoding: Learning to Allocate Verification in Speculative Decoding](https://arxiv.org/abs/2606.01019)
* [Medusa: Simple LLM Inference Acceleration Framework with Multiple Decoding Heads](https://arxiv.org/abs/2401.10774)
* [EAGLE: Speculative Sampling Requires Rethinking Feature Uncertainty](https://arxiv.org/abs/2401.15077)
* [Break the Sequential Dependency of LLM Inference Using Lookahead Decoding](https://arxiv.org/abs/2402.02057)
* [Draft & Verify: Lossless Large Language Model Acceleration via Self-Speculative Decoding](https://arxiv.org/abs/2309.08168)
* [LayerSkip: Enabling Early Exit Inference and Self-Speculative Decoding](https://arxiv.org/abs/2404.16710)
* [Confident Adaptive Language Modeling](https://arxiv.org/abs/2207.07061)
* [DeeBERT: Dynamic Early Exiting for Accelerating BERT Inference](https://arxiv.org/abs/2004.12993)
* [Efficient Memory Management for Large Language Model Serving with PagedAttention](https://arxiv.org/abs/2309.06180)
* [Prompt Cache: Modular Attention Reuse for Low-Latency Inference](https://arxiv.org/abs/2311.04934)
* [Don't Break the Cache: An Evaluation of Prompt Caching for Long-Horizon Agentic Tasks](https://arxiv.org/abs/2601.06007)
* [CacheBlend: Fast Large Language Model Serving for RAG with Cached Knowledge Fusion](https://arxiv.org/abs/2405.16444)
* [CacheGen: KV Cache Compression and Streaming for Fast Large Language Model Serving](https://arxiv.org/abs/2310.07240)
* [Cache-to-Cache: Direct Semantic Communication Between Large Language Models](https://arxiv.org/abs/2510.03215)
* [KVCOMM: Online Cross-context KV-cache Communication for Efficient LLM-based Multi-agent Systems](https://arxiv.org/abs/2510.12872)
* [KVQuant: Towards 10 Million Context Length LLM Inference with KV Cache Quantization](https://arxiv.org/abs/2401.18079)
* [H2O: Heavy-Hitter Oracle for Efficient Generative Inference of Large Language Models](https://arxiv.org/abs/2306.14048)
* [Scissorhands: Exploiting the Persistence of Importance Hypothesis for LLM KV Cache Compression at Test Time](https://arxiv.org/abs/2305.17118)
* [Model Tells You What to Discard: Adaptive KV Cache Compression for LLMs](https://arxiv.org/abs/2310.01801)
* [GPTQ: Accurate Post-Training Quantization for Generative Pre-trained Transformers](https://arxiv.org/abs/2210.17323)
* [AWQ: Activation-aware Weight Quantization for LLM Compression and Acceleration](https://arxiv.org/abs/2306.00978)
* [A Comprehensive Evaluation of Quantization Strategies for Large Language Models](https://arxiv.org/abs/2402.16775)
* [Matryoshka Representation Learning](https://arxiv.org/abs/2205.13147)
* [FrugalGPT: How to Use Large Language Models While Reducing Cost and Improving Performance](https://arxiv.org/abs/2305.05176)
* [RouteLLM: Learning to Route LLMs with Preference Data](https://arxiv.org/abs/2406.18665)
* [AutoMix: Automatically Mixing Language Models](https://arxiv.org/abs/2310.12963)
* [Hybrid LLM: Cost-Efficient and Quality-Aware Query Routing](https://arxiv.org/abs/2404.14618)
* [LLMLingua: Compressing Prompts for Accelerated Inference of Large Language Models](https://arxiv.org/abs/2310.05736)
* [LLMLingua-2: Data Distillation for Efficient and Faithful Task-Agnostic Prompt Compression](https://arxiv.org/abs/2403.12968)
* [Cache-Aware Prompt Compression: A Two-Tier Cost Model for LLM API Caching](https://arxiv.org/abs/2607.15516)
* [TokenSkip: Controllable Chain-of-Thought Compression in LLMs](https://arxiv.org/abs/2502.12067)
* [Token-Budget-Aware LLM Reasoning](https://arxiv.org/abs/2412.18547)
* [When to Reason: Semantic Router for vLLM](https://arxiv.org/abs/2510.08731)
* [Stop Overthinking: A Survey on Efficient Reasoning for Large Language Models](https://arxiv.org/abs/2503.16419)
