## baseline

| benchmark | cached_tokens | completion_tokens | draft_acceptance | errors | latency_s | n | prompt_s | prompt_tokens | reasoning_tokens | score | source | tag | task | think | throughput_tokens_per_s | tokens_per_s | truncated | ttft_s | wall_s | workers |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| baseline | 3.01 | 1588 | 0 | 0 | 126 | 200 | 0.2367 | 88.6 | 1357 | 0.905 | output/baseline/Q8_0/think | Q8_0 | gsm8k | true | 164.3 | 13.93 | 17 | 0.519 | 1932 | 16 |

## prefixcache

| cached_tokens | file_mb | latency_s | mode | n | prompt_s | prompt_tokens | restore_s | save_s | score | source | task | ttft_plus_restore_s | ttft_s |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| 0 | 0 | 2.407 | cold | 3 | 0.9955 | 5366 | 0 | 0 | 0 | output/prefixcache | longbench-2wikimqa | 1.018 | 1.018 |
| 0 | 0 | 1.859 | nocache | 3 | 0.9847 | 5393 | 0 | 0 | 0.06667 | output/prefixcache | longbench-2wikimqa | 1.004 | 1.004 |
| 5343 | 0 | 0.916 | warm | 3 | 0.0507 | 5393 | 0 | 0 | 0.06667 | output/prefixcache | longbench-2wikimqa | 0.07417 | 0.07417 |
| 5343 | 619.4 | 0.8878 | restored | 3 | 0.04873 | 5393 | 0.1749 | 0.9229 | 0.06667 | output/prefixcache | longbench-2wikimqa | 0.2425 | 0.06767 |

## promptcompress

| benchmark | cached_tokens | completion_tokens | draft_acceptance | errors | latency_s | n | overhead_s | overhead_tokens | prompt_s | prompt_tokens | reasoning_tokens | score | source | tag | task | think | throughput_tokens_per_s | tokens_per_s | total_s | truncated | ttft_s | wall_s | workers |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| promptcompress | 2.667 | 3.667 | 0 | 0 | 0.5406 | 3 | 0.00222 | 0 | 0.3136 | 560.3 | 0 | 0 | output/promptcompress/bm25-0.10/nothink | bm25-0.10 | longbench-2wikimqa | false | 18.99 | 21.38 | 0.5429 | 0 | 0.4233 | 0.5794 | 3 |
| promptcompress | 106.7 | 12 | 0 | 0 | 1.385 | 3 | 0.00377 | 0 | 0.4727 | 1404 | 0 | 0 | output/promptcompress/bm25-0.25/nothink | bm25-0.25 | longbench-2wikimqa | false | 19.82 | 19.17 | 1.388 | 0 | 0.8846 | 1.817 | 3 |
| promptcompress | 246.7 | 9 | 0 | 0 | 2.577 | 3 | 0.003871 | 0 | 0.8812 | 2749 | 0 | 0 | output/promptcompress/bm25-0.50/nothink | bm25-0.50 | longbench-2wikimqa | false | 9.404 | 12.06 | 2.581 | 0 | 1.913 | 2.871 | 3 |
| promptcompress | 132 | 5.667 | 0 | 0 | 0.403 | 3 | 1.066 | 0 | 0.1311 | 615.7 | 0 | 0 | output/promptcompress/embedding-0.10/nothink | embedding-0.10 | longbench-2wikimqa | false | 9.163 | 20.78 | 1.47 | 0 | 0.1773 | 1.855 | 3 |
| promptcompress | 77 | 3 | 0 | 0 | 0.5902 | 3 | 1.023 | 0 | 0.2706 | 1409 | 0 | 0 | output/promptcompress/embedding-0.25/nothink | embedding-0.25 | longbench-2wikimqa | false | 4.531 | 23.61 | 1.613 | 0 | 0.4332 | 1.986 | 3 |
| promptcompress | 123 | 17 | 0 | 0 | 2.621 | 3 | 1.176 | 0 | 0.7771 | 2777 | 0 | 0.05333 | output/promptcompress/embedding-0.50/nothink | embedding-0.50 | longbench-2wikimqa | false | 11.13 | 12.45 | 3.797 | 0 | 1.661 | 4.583 | 3 |
| promptcompress | 495 | 12.33 | 0 | 0 | 1.255 | 3 | 0.0005771 | 0 | 0.1441 | 614.3 | 0 | 0 | output/promptcompress/first-0.10/nothink | first-0.10 | longbench-2wikimqa | false | 23.45 | 14.56 | 1.256 | 0 | 0.5575 | 1.578 | 3 |
| promptcompress | 1282 | 12.67 | 0 | 0 | 1.547 | 3 | 0.0002159 | 0 | 0.1253 | 1374 | 0 | 0 | output/promptcompress/first-0.25/nothink | first-0.25 | longbench-2wikimqa | false | 19.9 | 17.93 | 1.547 | 0 | 0.9296 | 1.91 | 3 |
| promptcompress | 2576 | 6.667 | 0 | 0 | 1.54 | 3 | 0.0003781 | 0 | 0.1462 | 2689 | 0 | 0.1667 | output/promptcompress/first-0.50/nothink | first-0.50 | longbench-2wikimqa | false | 11.91 | 19.52 | 1.541 | 0 | 1.265 | 1.679 | 3 |
| promptcompress | 3 | 22 | 0 | 0 | 4.264 | 3 | 1.76e-07 | 0 | 1.333 | 5393 | 0 | 0.06667 | output/promptcompress/full/nothink | full | longbench-2wikimqa | false | 13.97 | 14.97 | 4.264 | 0 | 2.729 | 4.723 | 3 |
| promptcompress | 18.33 | 8 | 0 | 0 | 1.063 | 3 | 2.846 | 5427 | 0.08493 | 82 | 0 | 0.1667 | output/promptcompress/llm-0.03/nothink | llm-0.03 | longbench-2wikimqa | false | 5.005 | 18 | 3.909 | 0 | 0.63 | 4.796 | 3 |
| promptcompress | 2.667 | 11.67 | 0 | 0 | 1.052 | 3 | 6.614 | 5469 | 0.0808 | 123 | 0 | 0.1667 | output/promptcompress/llm-0.10/nothink | llm-0.10 | longbench-2wikimqa | false | 3.785 | 25.21 | 7.666 | 0 | 0.573 | 9.248 | 3 |
| promptcompress | 3 | 8.667 | 0 | 0 | 0.7791 | 3 | 0.0002621 | 0 | 0.3181 | 551 | 0 | 0.5 | output/promptcompress/random-0.10/nothink | random-0.10 | longbench-2wikimqa | false | 25.83 | 20.99 | 0.7794 | 0 | 0.4589 | 1.007 | 3 |
| promptcompress | 4 | 8 | 0 | 0 | 1.478 | 3 | 0.000223 | 0 | 0.6071 | 1349 | 0 | 0.1667 | output/promptcompress/random-0.25/nothink | random-0.25 | longbench-2wikimqa | false | 13.97 | 16.84 | 1.478 | 0 | 1.038 | 1.718 | 3 |
| promptcompress | 73.33 | 10.33 | 0 | 0 | 2.38 | 3 | 0.0002317 | 0 | 0.725 | 2626 | 0 | 0 | output/promptcompress/random-0.50/nothink | random-0.50 | longbench-2wikimqa | false | 10.88 | 17.33 | 2.38 | 0 | 1.844 | 2.849 | 3 |

## semanticcache

| cache_size | dataset | error_rate | f1 | hit_rate | hits | llm_calls | missed_hits | policy | precision | recall | requests | reusable | seconds | source | threshold | wrong_hits |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| 3086 | lmarena | 0.1557 | 0.9103 | 0.9516 | 6.071e+04 | 3086 | 70 | static | 0.8364 | 0.9986 | 6.38e+04 | 5.085e+04 | 63.92 | output/semanticcache/lmarena | 0.7 | 9934 |
| 3438 | lmarena | 0.1156 | 0.9329 | 0.9461 | 6.036e+04 | 3438 | 247 | static | 0.8778 | 0.9954 | 6.38e+04 | 5.323e+04 | 67.18 | output/semanticcache/lmarena | 0.75 | 7373 |
| 3962 | lmarena | 0.09813 | 0.9392 | 0.9379 | 5.983e+04 | 3962 | 682 | static | 0.8954 | 0.9874 | 6.38e+04 | 5.426e+04 | 49.46 | output/semanticcache/lmarena | 0.8 | 6260 |
| 4914 | lmarena | 0.0858 | 0.9384 | 0.923 | 5.888e+04 | 4914 | 1540 | static | 0.907 | 0.972 | 6.38e+04 | 5.495e+04 | 72.03 | output/semanticcache/lmarena | 0.84 | 5474 |
| 5726 | lmarena | 0.07999 | 0.9347 | 0.9102 | 5.807e+04 | 5726 | 2301 | static | 0.9121 | 0.9584 | 6.38e+04 | 5.527e+04 | 79.95 | output/semanticcache/lmarena | 0.86 | 5103 |
| 6995 | lmarena | 0.07461 | 0.9265 | 0.8904 | 5.68e+04 | 6995 | 3500 | static | 0.9162 | 0.937 | 6.38e+04 | 5.554e+04 | 92.4 | output/semanticcache/lmarena | 0.88 | 4760 |
| 8970 | lmarena | 0.06862 | 0.9121 | 0.8594 | 5.483e+04 | 8970 | 5346 | static | 0.9201 | 0.9042 | 6.38e+04 | 5.579e+04 | 106.9 | output/semanticcache/lmarena | 0.9 | 4378 |
| 1.194e+04 | lmarena | 0.06135 | 0.8884 | 0.8129 | 5.186e+04 | 1.194e+04 | 8134 | static | 0.9245 | 0.8549 | 6.38e+04 | 5.608e+04 | 130.9 | output/semanticcache/lmarena | 0.92 | 3914 |
| 1.678e+04 | lmarena | 0.054 | 0.8438 | 0.737 | 4.702e+04 | 1.678e+04 | 1.268e+04 | static | 0.9267 | 0.7745 | 6.38e+04 | 5.626e+04 | 159.1 | output/semanticcache/lmarena | 0.94 | 3445 |
| 2.038e+04 | lmarena | 0.04958 | 0.8073 | 0.6805 | 4.341e+04 | 2.038e+04 | 1.605e+04 | static | 0.9271 | 0.7149 | 6.38e+04 | 5.63e+04 | 169.7 | output/semanticcache/lmarena | 0.95 | 3163 |
| 2.494e+04 | lmarena | 0.04355 | 0.7579 | 0.6091 | 3.886e+04 | 2.494e+04 | 2.028e+04 | static | 0.9285 | 0.6402 | 6.38e+04 | 5.635e+04 | 213.6 | output/semanticcache/lmarena | 0.96 | 2778 |
| 3.078e+04 | lmarena | 0.03706 | 0.6859 | 0.5176 | 3.302e+04 | 3.078e+04 | 2.571e+04 | static | 0.9284 | 0.5439 | 6.38e+04 | 5.636e+04 | 235.6 | output/semanticcache/lmarena | 0.97 | 2364 |
| 3.825e+04 | lmarena | 0.02851 | 0.579 | 0.4004 | 2.554e+04 | 3.825e+04 | 3.268e+04 | static | 0.9288 | 0.4206 | 6.38e+04 | 5.641e+04 | 264 | output/semanticcache/lmarena | 0.98 | 1819 |
| 4.714e+04 | lmarena | 0.0195 | 0.4217 | 0.2611 | 1.666e+04 | 4.714e+04 | 4.102e+04 | static | 0.9253 | 0.2731 | 6.38e+04 | 5.644e+04 | 316.7 | output/semanticcache/lmarena | 0.99 | 1244 |

## thinkingbudget

| benchmark | cached_tokens | completion_tokens | draft_acceptance | errors | latency_s | n | overhead_s | overhead_tokens | prompt_s | prompt_tokens | reasoning_tokens | score | source | tag | task | think | throughput_tokens_per_s | tokens_per_s | total_s | truncated | ttft_s | wall_s | workers |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| thinkingbudget | 81.88 | 174.9 | 0.4984 | 0 | 9.513 | 8 | 0 | 0 | 0.08734 | 84.12 | 0 | 0.75 | output/thinkingbudget/budget-0/think | budget-0 | gsm8k | true | 91.9 | 18.37 | 9.513 | 0 | 0.133 | 15.22 | 8 |
| thinkingbudget | 79.12 | 796.9 | 0.3613 | 0 | 84.14 | 8 | 0 | 0 | 0.05026 | 80.12 | 618.5 | 0.875 | output/thinkingbudget/budget-1024/think | budget-1024 | gsm8k | true | 47.4 | 10.3 | 84.14 | 0 | 1.231 | 134.5 | 8 |
| thinkingbudget | 79.12 | 273.4 | 0.354 | 0 | 19.28 | 8 | 0 | 0 | 0.03126 | 80.12 | 126.9 | 0.75 | output/thinkingbudget/budget-128/think | budget-128 | gsm8k | true | 90.17 | 14.36 | 19.28 | 0 | 0.7386 | 24.25 | 8 |
| thinkingbudget | 79.12 | 869.6 | 0.3487 | 0 | 92 | 8 | 0 | 0 | 0.05251 | 80.12 | 735.6 | 0.75 | output/thinkingbudget/budget-2048/think | budget-2048 | gsm8k | true | 42.08 | 10.29 | 92 | 2 | 1.693 | 165.3 | 8 |
| thinkingbudget | 79.12 | 407.2 | 0.3743 | 0 | 30.04 | 8 | 0 | 0 | 0.05738 | 80.12 | 253.9 | 0.875 | output/thinkingbudget/budget-256/think | budget-256 | gsm8k | true | 89.62 | 13.76 | 30.04 | 0 | 0.8547 | 36.35 | 8 |
| thinkingbudget | 79.12 | 584.2 | 0.3774 | 0 | 51.78 | 8 | 0 | 0 | 0.04668 | 80.12 | 439.5 | 0.875 | output/thinkingbudget/budget-512/think | budget-512 | gsm8k | true | 68.03 | 11.92 | 51.78 | 0 | 1.096 | 68.71 | 8 |

## tracecompress

| benchmark | cached_tokens | completion_tokens | draft_acceptance | errors | latency_s | n | overhead_s | overhead_tokens | prompt_s | prompt_tokens | reasoning_tokens | score | source | tag | task | think | throughput_tokens_per_s | tokens_per_s | total_s | truncated | ttft_s | wall_s | workers |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| tracecompress | 3 | 71.67 | 0 | 0 | 5.687 | 6 | 0 | 0 | 0.2638 | 108.3 | 0 | 0.6667 | output/tracecompress/cod/nothink | cod | gsm8k | false | 54.59 | 13.01 | 5.687 | 0 | 0.3868 | 7.877 | 6 |
| tracecompress | 3 | 162.8 | 0 | 0 | 13.64 | 6 | 0 | 0 | 0.5334 | 763.7 | 0 | 1 | output/tracecompress/full/nothink | full | gsm8k | false | 47.34 | 12.04 | 13.64 | 0 | 0.8763 | 20.64 | 6 |
| tracecompress | 83 | 149.2 | 0 | 0 | 10.96 | 6 | 0 | 0 | 0.3408 | 314.8 | 0 | 1 | output/tracecompress/head-150/nothink | head-150 | gsm8k | false | 53.53 | 13.38 | 10.96 | 0 | 0.4989 | 16.72 | 6 |
| tracecompress | 67 | 142.8 | 0 | 0 | 11.05 | 6 | 0 | 0 | 0.2729 | 166.3 | 0 | 0.8333 | output/tracecompress/head-50/nothink | head-50 | gsm8k | false | 55.24 | 13.14 | 11.05 | 0 | 0.6505 | 15.51 | 6 |
| tracecompress | 2.5 | 140.3 | 0 | 0 | 7.387 | 6 | 0 | 0 | 0.2999 | 79.33 | 0 | 0.8333 | output/tracecompress/none/nothink | none | gsm8k | false | 67.58 | 22.25 | 7.387 | 0 | 0.3284 | 12.46 | 6 |
| tracecompress | 7.333 | 87.83 | 0 | 0 | 6.451 | 6 | 4.812 | 765.2 | 0.1291 | 150.3 | 0 | 1 | output/tracecompress/summary-30/nothink | summary-30 | gsm8k | false | 36.11 | 13.83 | 11.26 | 0 | 0.2528 | 14.59 | 6 |
| tracecompress | 7.5 | 80.5 | 0 | 0 | 5.171 | 6 | 6.28 | 774.7 | 0.1187 | 159.8 | 0 | 1 | output/tracecompress/summary-80/nothink | summary-80 | gsm8k | false | 32.23 | 15.54 | 11.45 | 0 | 0.2451 | 14.99 | 6 |
| tracecompress | 68.33 | 134.5 | 0 | 0 | 8.597 | 6 | 0 | 0 | 0.3806 | 331.8 | 0 | 1 | output/tracecompress/tail-150/nothink | tail-150 | gsm8k | false | 63.86 | 16.42 | 8.597 | 0 | 0.5718 | 12.64 | 6 |
| tracecompress | 11.67 | 129.2 | 0 | 0 | 10.52 | 6 | 0 | 0 | 0.2673 | 182.3 | 0 | 1 | output/tracecompress/tail-50/nothink | tail-50 | gsm8k | false | 53.23 | 12.45 | 10.52 | 0 | 0.5011 | 14.56 | 6 |

## vcache

| cache_size | dataset | delta | error_rate | f1 | hit_rate | hits | llm_calls | missed_hits | policy | precision | recall | requests | reusable | seconds | source | wrong_hits |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| 2026 | lmarena | 0.005 | 0 | 0 | 0 | 0 | 3000 | 974 | vcache | 0 | 0 | 3000 | 974 | 1.084 | output/vcache/lmarena | 0 |
| 2026 | lmarena | 0.01 | 0 | 0 | 0 | 0 | 3000 | 974 | vcache | 0 | 0 | 3000 | 974 | 0.9601 | output/vcache/lmarena | 0 |
| 2025 | lmarena | 0.015 | 0.0003333 | 0 | 0.0003333 | 1 | 2999 | 974 | vcache | 0 | 0 | 3000 | 974 | 0.9841 | output/vcache/lmarena | 1 |
| 2025 | lmarena | 0.02 | 0.0003333 | 0.002049 | 0.0006667 | 2 | 2998 | 973 | vcache | 0.5 | 0.001027 | 3000 | 974 | 0.9856 | output/vcache/lmarena | 1 |
| 2024 | lmarena | 0.03 | 0.0006667 | 0.00409 | 0.001333 | 4 | 2996 | 972 | vcache | 0.5 | 0.002053 | 3000 | 974 | 1.079 | output/vcache/lmarena | 2 |
| 2024 | lmarena | 0.05 | 0.001 | 0.01019 | 0.002667 | 8 | 2992 | 968 | vcache | 0.625 | 0.005139 | 3000 | 973 | 1.004 | output/vcache/lmarena | 3 |
| 2023 | lmarena | 0.07 | 0.001333 | 0.01221 | 0.003333 | 10 | 2990 | 967 | vcache | 0.6 | 0.006166 | 3000 | 973 | 1.003 | output/vcache/lmarena | 4 |
| 2020 | lmarena | 0.1 | 0.003 | 0.0182 | 0.006 | 18 | 2982 | 962 | vcache | 0.5 | 0.009269 | 3000 | 971 | 1.065 | output/vcache/lmarena | 9 |

