## baseline

| benchmark | cached_tokens | completion_tokens | draft_acceptance | errors | latency_s | n | overhead_s | overhead_tokens | prompt_s | prompt_tokens | reasoning_tokens | score | source | tag | task | think | throughput_tokens_per_s | tokens_per_s | total_s | truncated | ttft_s | wall_s | workers |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| baseline | 3.015 | 237.6 | 0 | 0 | 14.18 | 200 | 0 | 0 | 0.2174 | 92.6 | 0 | 0.96 | output/baseline/Q8_0/nothink | Q8_0 | gsm8k | false | 245.7 | 17.64 | 14.18 | 0 | 0.4347 | 193.4 | 16 |
| baseline | 3.105 | 795.9 | 0 | 0 | 53.08 | 200 | 0 | 0 | 0.2197 | 97.81 | 0 | 0.75 | output/baseline/Q8_0/nothink | Q8_0 | math500 | false | 168.4 | 16.1 | 53.08 | 5 | 0.3308 | 945.3 | 16 |
| baseline | 3.08 | 254.5 | 0 | 0 | 16.12 | 200 | 0 | 0 | 0.2315 | 121.9 | 0 | 0.75 | output/baseline/Q8_0/nothink | Q8_0 | mmlu | false | 153.5 | 17 | 16.12 | 2 | 0.3812 | 331.6 | 16 |
| baseline | 4.875 | 600.2 | 0 | 0 | 42.62 | 200 | 0 | 0 | 0.245 | 215.7 | 0 | 0.625 | output/baseline/Q8_0/nothink | Q8_0 | mmlupro | false | 162.9 | 15.46 | 42.62 | 8 | 0.4007 | 736.8 | 16 |
| baseline | 4.875 | 102.8 | 0 | 0 | 12.37 | 200 | 0 | 0 | 0.5088 | 853.7 | 0 | 0.202 | output/baseline/Q8_0/nothink | Q8_0 | cnndm | false | 126.8 | 8.997 | 12.37 | 0 | 0.9011 | 162.2 | 16 |
| baseline | 3.005 | 1718 | 0 | 0 | 115 | 200 | 0 | 0 | 0.1987 | 88.6 | 1453 | 0.98 | output/baseline/Q8_0/think | Q8_0 | gsm8k | true | 162.2 | 15.28 | 115 | 0 | 0.4799 | 2118 | 12 |
| baseline | 3.095 | 4943 | 0 | 0 | 452.3 | 200 | 0 | 0 | 0.2231 | 93.81 | 4332 | 0.845 | output/baseline/Q8_0/think | Q8_0 | math500 | true | 113.6 | 11.42 | 452.3 | 9 | 0.7545 | 8700 | 12 |
| baseline | 3.035 | 1659 | 0 | 0 | 124.3 | 200 | 0 | 0 | 0.204 | 117.9 | 1521 | 0.79 | output/baseline/Q8_0/think | Q8_0 | mmlu | true | 143.7 | 14.94 | 124.3 | 1 | 1.475 | 2309 | 12 |
| baseline | 3.075 | 3359 | 0 | 0 | 277.7 | 200 | 0 | 0 | 0.2372 | 211.7 | 3067 | 0.745 | output/baseline/Q8_0/think | Q8_0 | mmlupro | true | 135.1 | 12.77 | 277.7 | 2 | 0.729 | 4971 | 12 |
| baseline | 3.01 | 1588 | 0 | 0 | 126 | 200 |  |  | 0.2367 | 88.6 | 1357 | 0.905 | output/baseline/Q8_0/think-cap4096 | Q8_0 | gsm8k | true | 164.3 | 13.93 |  | 17 | 0.519 | 1932 | 16 |
| baseline | 3.075 | 3187 | 0 | 0 | 272 | 200 |  |  | 0.3371 | 93.81 | 2887 | 0.5 | output/baseline/Q8_0/think-cap4096 | Q8_0 | math500 | true | 167.9 | 12.29 |  | 87 | 1.371 | 3797 | 16 |
| baseline | 3.13 | 1354 | 0 | 0 | 96.74 | 200 |  |  | 0.2308 | 117.9 | 1271 | 0.745 | output/baseline/Q8_0/think-cap4096 | Q8_0 | mmlu | true | 193.5 | 15.43 |  | 17 | 0.6664 | 1400 | 16 |
| baseline | 3.075 | 2442 | 0 | 0 | 179 | 200 |  |  | 0.2618 | 211.7 | 2301 | 0.565 | output/baseline/Q8_0/think-cap4096 | Q8_0 | mmlupro | true | 204.7 | 14.01 |  | 54 | 0.8072 | 2385 | 16 |
| baseline | 7.54 | 7.16 | 0 | 0 | 24.07 | 50 | 0 | 0 | 13.77 | 1.342e+04 | 0 | 0.5908 | output/baseline/Q8_0-long/nothink | Q8_0-long | longbench-hotpotqa | false | 1.184 | 1.77 | 24.07 | 0 | 18.57 | 302.3 | 4 |
| baseline | 8 | 7.9 | 0 | 0 | 14.22 | 50 | 0 | 0 | 7.023 | 8246 | 0 | 0.3451 | output/baseline/Q8_0-long/nothink | Q8_0-long | longbench-2wikimqa | false | 2.202 | 2.688 | 14.22 | 0 | 10.07 | 179.4 | 4 |
| baseline | 147.8 | 7.78 | 0 | 0 | 29.75 | 50 | 0 | 0 | 17.04 | 1.65e+04 | 0 | 0.3676 | output/baseline/Q8_0-long/nothink | Q8_0-long | longbench-musique | false | 1.027 | 1.763 | 29.75 | 0 | 23.58 | 378.8 | 4 |

## cascade

| cheap | escalation_rate | mean_completion_tokens | mean_latency_s | mean_strong_tokens | n | policy | score | source | task |
|---|---|---|---|---|---|---|---|---|---|
| Qwen3-1.7B-Q8_0 | 0 | 226.8 | 4.89 | 0 | 200 | cheap | 0.845 | output/cascade/Qwen3-1.7B-Q8_0 | gsm8k |
| Qwen3-1.7B-Q8_0 | 1 | 1718 | 115 | 1718 | 200 | strong | 0.98 | output/cascade/Qwen3-1.7B-Q8_0 | gsm8k |
| Qwen3-1.7B-Q8_0 | 0.065 | 852.4 | 17.14 | 165.7 | 200 | consistency-2of3 | 0.905 | output/cascade/Qwen3-1.7B-Q8_0 | gsm8k |
| Qwen3-1.7B-Q8_0 | 0.25 | 1323 | 48.15 | 636.2 | 200 | consistency-3of3 | 0.96 | output/cascade/Qwen3-1.7B-Q8_0 | gsm8k |
| Qwen3-1.7B-Q8_0 | 0 | 228.8 | 4.997 | 0 | 200 | verify | 0.845 | output/cascade/Qwen3-1.7B-Q8_0 | gsm8k |
| Qwen3-1.7B-Q8_0 | 0.155 | 585 | 29.36 | 358.1 | 200 | oracle | 0.985 | output/cascade/Qwen3-1.7B-Q8_0 | gsm8k |
| Qwen3-1.7B-Q8_0 | 0 | 93.98 | 2.454 | 0 | 200 | cheap | 0.62 | output/cascade/Qwen3-1.7B-Q8_0 | mmlu |
| Qwen3-1.7B-Q8_0 | 1 | 1659 | 124.3 | 1659 | 200 | strong | 0.79 | output/cascade/Qwen3-1.7B-Q8_0 | mmlu |
| Qwen3-1.7B-Q8_0 | 0.025 | 415.5 | 13.74 | 126.9 | 200 | consistency-2of3 | 0.62 | output/cascade/Qwen3-1.7B-Q8_0 | mmlu |
| Qwen3-1.7B-Q8_0 | 0.165 | 714.9 | 35.97 | 426.3 | 200 | consistency-3of3 | 0.685 | output/cascade/Qwen3-1.7B-Q8_0 | mmlu |
| Qwen3-1.7B-Q8_0 | 0.09 | 374 | 24.72 | 277.9 | 200 | verify | 0.675 | output/cascade/Qwen3-1.7B-Q8_0 | mmlu |
| Qwen3-1.7B-Q8_0 | 0.38 | 909.4 | 64.5 | 815.4 | 200 | oracle | 0.875 | output/cascade/Qwen3-1.7B-Q8_0 | mmlu |
| Qwen3-1.7B-Q8_0 | 0 | 305.3 | 7.491 | 0 | 200 | cheap | 0.415 | output/cascade/Qwen3-1.7B-Q8_0 | mmlupro |
| Qwen3-1.7B-Q8_0 | 1 | 3359 | 277.7 | 3359 | 200 | strong | 0.745 | output/cascade/Qwen3-1.7B-Q8_0 | mmlupro |
| Qwen3-1.7B-Q8_0 | 0.115 | 1565 | 64 | 653.8 | 200 | consistency-2of3 | 0.485 | output/cascade/Qwen3-1.7B-Q8_0 | mmlupro |
| Qwen3-1.7B-Q8_0 | 0.405 | 2851 | 170 | 1940 | 200 | consistency-3of3 | 0.61 | output/cascade/Qwen3-1.7B-Q8_0 | mmlupro |
| Qwen3-1.7B-Q8_0 | 0.14 | 1051 | 71.24 | 743.8 | 200 | verify | 0.485 | output/cascade/Qwen3-1.7B-Q8_0 | mmlupro |
| Qwen3-1.7B-Q8_0 | 0.585 | 2675 | 206.4 | 2370 | 200 | oracle | 0.77 | output/cascade/Qwen3-1.7B-Q8_0 | mmlupro |
| Qwen3-8B-Q3_K_M | 0 | 222.1 | 8.416 | 0 | 200 | cheap | 0.955 | output/cascade/Qwen3-8B-Q3_K_M | gsm8k |
| Qwen3-8B-Q3_K_M | 1 | 1718 | 115 | 1718 | 200 | strong | 0.98 | output/cascade/Qwen3-8B-Q3_K_M | gsm8k |
| Qwen3-8B-Q3_K_M | 0.02 | 759.3 | 16.82 | 88.69 | 200 | consistency-2of3 | 0.975 | output/cascade/Qwen3-8B-Q3_K_M | gsm8k |
| Qwen3-8B-Q3_K_M | 0.06 | 897.1 | 26.69 | 226.5 | 200 | consistency-3of3 | 0.975 | output/cascade/Qwen3-8B-Q3_K_M | gsm8k |
| Qwen3-8B-Q3_K_M | 0.005 | 269.5 | 11.61 | 43.45 | 200 | verify | 0.96 | output/cascade/Qwen3-8B-Q3_K_M | gsm8k |
| Qwen3-8B-Q3_K_M | 0.045 | 423 | 22.1 | 201 | 200 | oracle | 0.98 | output/cascade/Qwen3-8B-Q3_K_M | gsm8k |
| Qwen3-8B-Q3_K_M | 0 | 115.6 | 5.171 | 0 | 200 | cheap | 0.75 | output/cascade/Qwen3-8B-Q3_K_M | mmlu |
| Qwen3-8B-Q3_K_M | 1 | 1659 | 124.3 | 1659 | 200 | strong | 0.79 | output/cascade/Qwen3-8B-Q3_K_M | mmlu |
| Qwen3-8B-Q3_K_M | 0.015 | 443.6 | 14.64 | 80.31 | 200 | consistency-2of3 | 0.765 | output/cascade/Qwen3-8B-Q3_K_M | mmlu |
| Qwen3-8B-Q3_K_M | 0.135 | 745.9 | 36.74 | 382.6 | 200 | consistency-3of3 | 0.79 | output/cascade/Qwen3-8B-Q3_K_M | mmlu |
| Qwen3-8B-Q3_K_M | 0.18 | 578.6 | 39.4 | 459.2 | 200 | verify | 0.775 | output/cascade/Qwen3-8B-Q3_K_M | mmlu |
| Qwen3-8B-Q3_K_M | 0.25 | 730 | 53.16 | 614.4 | 200 | oracle | 0.85 | output/cascade/Qwen3-8B-Q3_K_M | mmlu |
| Qwen3-8B-Q3_K_M | 0 | 305.3 | 12.5 | 0 | 200 | cheap | 0.595 | output/cascade/Qwen3-8B-Q3_K_M | mmlupro |
| Qwen3-8B-Q3_K_M | 1 | 3359 | 277.7 | 3359 | 200 | strong | 0.745 | output/cascade/Qwen3-8B-Q3_K_M | mmlupro |
| Qwen3-8B-Q3_K_M | 0.04 | 1304 | 46.73 | 332.6 | 200 | consistency-2of3 | 0.635 | output/cascade/Qwen3-8B-Q3_K_M | mmlupro |
| Qwen3-8B-Q3_K_M | 0.285 | 2456 | 141.9 | 1485 | 200 | consistency-3of3 | 0.69 | output/cascade/Qwen3-8B-Q3_K_M | mmlupro |
| Qwen3-8B-Q3_K_M | 0.26 | 1278 | 91.95 | 968.7 | 200 | verify | 0.655 | output/cascade/Qwen3-8B-Q3_K_M | mmlupro |
| Qwen3-8B-Q3_K_M | 0.405 | 2212 | 173.2 | 1907 | 200 | oracle | 0.77 | output/cascade/Qwen3-8B-Q3_K_M | mmlupro |

## embedquant

| bytes_per_entry | cache_size | dataset | dimensions | error_rate | f1 | hit_rate | hits | llm_calls | lookup_us | missed_hits | nn_agreement | precision | recall | requests | reusable | source | threshold | variant | wrong_hits |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| 6144 | 2688 | lmarena-openai | 1536 | 0.2103 | 0.8615 | 0.8656 | 17312 | 2688 | 1.828e+04 | 9 | 1 | 0.757 | 0.9993 | 20000 | 13115 | output/embedquant/lmarena-openai | 0.5 | float32 | 4206 |
| 6144 | 2864 | lmarena-openai | 1536 | 0.1571 | 0.8984 | 0.8568 | 17136 | 2864 | 1.828e+04 | 24 | 1 | 0.8166 | 0.9983 | 20000 | 14018 | output/embedquant/lmarena-openai | 0.55 | float32 | 3142 |
| 6144 | 2987 | lmarena-openai | 1536 | 0.1293 | 0.9164 | 0.8507 | 17013 | 2987 | 1.828e+04 | 49 | 1 | 0.8481 | 0.9966 | 20000 | 14477 | output/embedquant/lmarena-openai | 0.6 | float32 | 2585 |
| 6144 | 3110 | lmarena-openai | 1536 | 0.1095 | 0.9278 | 0.8445 | 16890 | 3110 | 1.828e+04 | 99 | 1 | 0.8704 | 0.9933 | 20000 | 14800 | output/embedquant/lmarena-openai | 0.65 | float32 | 2189 |
| 6144 | 3250 | lmarena-openai | 1536 | 0.0972 | 0.9328 | 0.8375 | 16750 | 3250 | 1.828e+04 | 189 | 1 | 0.8839 | 0.9874 | 20000 | 14995 | output/embedquant/lmarena-openai | 0.7 | float32 | 1944 |
| 6144 | 3509 | lmarena-openai | 1536 | 0.08825 | 0.9314 | 0.8246 | 16491 | 3509 | 1.828e+04 | 403 | 1 | 0.893 | 0.9734 | 20000 | 15129 | output/embedquant/lmarena-openai | 0.75 | float32 | 1765 |
| 6144 | 4043 | lmarena-openai | 1536 | 0.07935 | 0.9209 | 0.7978 | 15957 | 4043 | 1.828e+04 | 881 | 1 | 0.9005 | 0.9422 | 20000 | 15251 | output/embedquant/lmarena-openai | 0.8 | float32 | 1587 |
| 6144 | 4899 | lmarena-openai | 1536 | 0.07065 | 0.8988 | 0.755 | 15101 | 4899 | 1.828e+04 | 1670 | 1 | 0.9064 | 0.8913 | 20000 | 15358 | output/embedquant/lmarena-openai | 0.84 | float32 | 1413 |
| 6144 | 6412 | lmarena-openai | 1536 | 0.0592 | 0.854 | 0.6794 | 13588 | 6412 | 1.828e+04 | 3056 | 1 | 0.9129 | 0.8023 | 20000 | 15460 | output/embedquant/lmarena-openai | 0.88 | float32 | 1184 |
| 6144 | 7551 | lmarena-openai | 1536 | 0.05315 | 0.8146 | 0.6224 | 12449 | 7551 | 1.828e+04 | 4120 | 1 | 0.9146 | 0.7343 | 20000 | 15506 | output/embedquant/lmarena-openai | 0.9 | float32 | 1063 |
| 6144 | 9111 | lmarena-openai | 1536 | 0.04495 | 0.7554 | 0.5444 | 10889 | 9111 | 1.828e+04 | 5571 | 1 | 0.9174 | 0.642 | 20000 | 15561 | output/embedquant/lmarena-openai | 0.92 | float32 | 899 |
| 6144 | 11188 | lmarena-openai | 1536 | 0.037 | 0.6619 | 0.4406 | 8812 | 11188 | 1.828e+04 | 7508 | 1 | 0.916 | 0.5181 | 20000 | 15580 | output/embedquant/lmarena-openai | 0.94 | float32 | 740 |
| 6144 | 13784 | lmarena-openai | 1536 | 0.0254 | 0.5228 | 0.3108 | 6216 | 13784 | 1.828e+04 | 9914 | 1 | 0.9183 | 0.3654 | 20000 | 15622 | output/embedquant/lmarena-openai | 0.96 | float32 | 508 |
| 6144 | 16660 | lmarena-openai | 1536 | 0.015 | 0.3206 | 0.167 | 3340 | 16660 | 1.828e+04 | 12582 | 1 | 0.9102 | 0.1946 | 20000 | 15622 | output/embedquant/lmarena-openai | 0.98 | float32 | 300 |
| 1536 | 2691 | lmarena-openai | 1536 | 0.2089 | 0.8625 | 0.8655 | 17309 | 2691 | 1.177e+04 | 10 | 0.9883 | 0.7586 | 0.9992 | 20000 | 13141 | output/embedquant/lmarena-openai | 0.5 | int8 | 4178 |
| 1536 | 2866 | lmarena-openai | 1536 | 0.1564 | 0.8989 | 0.8567 | 17134 | 2866 | 1.177e+04 | 24 | 0.9933 | 0.8174 | 0.9983 | 20000 | 14030 | output/embedquant/lmarena-openai | 0.55 | int8 | 3128 |
| 1536 | 2991 | lmarena-openai | 1536 | 0.1285 | 0.9168 | 0.8505 | 17009 | 2991 | 1.177e+04 | 50 | 0.9942 | 0.8488 | 0.9965 | 20000 | 14488 | output/embedquant/lmarena-openai | 0.6 | int8 | 2571 |
| 1536 | 3112 | lmarena-openai | 1536 | 0.109 | 0.9281 | 0.8444 | 16888 | 3112 | 1.177e+04 | 99 | 0.9961 | 0.8709 | 0.9933 | 20000 | 14806 | output/embedquant/lmarena-openai | 0.65 | int8 | 2181 |
| 1536 | 3257 | lmarena-openai | 1536 | 0.09735 | 0.9325 | 0.8371 | 16743 | 3257 | 1.177e+04 | 195 | 0.993 | 0.8837 | 0.987 | 20000 | 14991 | output/embedquant/lmarena-openai | 0.7 | int8 | 1947 |
| 1536 | 3516 | lmarena-openai | 1536 | 0.0885 | 0.9311 | 0.8242 | 16484 | 3516 | 1.177e+04 | 408 | 0.9951 | 0.8926 | 0.973 | 20000 | 15122 | output/embedquant/lmarena-openai | 0.75 | int8 | 1770 |
| 1536 | 4052 | lmarena-openai | 1536 | 0.0793 | 0.9207 | 0.7974 | 15948 | 4052 | 1.177e+04 | 889 | 0.9948 | 0.9006 | 0.9417 | 20000 | 15251 | output/embedquant/lmarena-openai | 0.8 | int8 | 1586 |
| 1536 | 4914 | lmarena-openai | 1536 | 0.0705 | 0.8984 | 0.7543 | 15086 | 4914 | 1.177e+04 | 1684 | 0.9948 | 0.9065 | 0.8904 | 20000 | 15360 | output/embedquant/lmarena-openai | 0.84 | int8 | 1410 |
| 1536 | 6433 | lmarena-openai | 1536 | 0.05905 | 0.8534 | 0.6784 | 13567 | 6433 | 1.177e+04 | 3076 | 0.9928 | 0.913 | 0.8011 | 20000 | 15462 | output/embedquant/lmarena-openai | 0.88 | int8 | 1181 |
| 1536 | 7580 | lmarena-openai | 1536 | 0.05285 | 0.8138 | 0.621 | 12420 | 7580 | 1.177e+04 | 4144 | 0.9923 | 0.9149 | 0.7328 | 20000 | 15507 | output/embedquant/lmarena-openai | 0.9 | int8 | 1057 |
| 1536 | 9147 | lmarena-openai | 1536 | 0.04485 | 0.7539 | 0.5426 | 10853 | 9147 | 1.177e+04 | 5603 | 0.9916 | 0.9174 | 0.6399 | 20000 | 15559 | output/embedquant/lmarena-openai | 0.92 | int8 | 897 |
| 1536 | 11223 | lmarena-openai | 1536 | 0.03675 | 0.6604 | 0.4389 | 8777 | 11223 | 1.177e+04 | 7537 | 0.9936 | 0.9163 | 0.5162 | 20000 | 15579 | output/embedquant/lmarena-openai | 0.94 | int8 | 735 |
| 1536 | 13812 | lmarena-openai | 1536 | 0.02525 | 0.5211 | 0.3094 | 6188 | 13812 | 1.177e+04 | 9939 | 0.9933 | 0.9184 | 0.3638 | 20000 | 15622 | output/embedquant/lmarena-openai | 0.96 | int8 | 505 |
| 1536 | 16677 | lmarena-openai | 1536 | 0.0149 | 0.3193 | 0.1661 | 3323 | 16677 | 1.177e+04 | 12598 | 0.9934 | 0.9103 | 0.1936 | 20000 | 15623 | output/embedquant/lmarena-openai | 0.98 | int8 | 298 |
| 192 | 3322 | lmarena-openai | 1536 | 0.09585 | 0.9316 | 0.8339 | 16678 | 3322 | 428.1 | 251 | 0.7674 | 0.8851 | 0.9833 | 20000 | 15012 | output/embedquant/lmarena-openai | 0.5 | binary | 1917 |
| 192 | 3664 | lmarena-openai | 1536 | 0.08625 | 0.928 | 0.8168 | 16336 | 3664 | 428.1 | 542 | 0.8064 | 0.8944 | 0.9642 | 20000 | 15153 | output/embedquant/lmarena-openai | 0.55 | binary | 1725 |
| 192 | 4290 | lmarena-openai | 1536 | 0.0772 | 0.9142 | 0.7855 | 15710 | 4290 | 428.1 | 1116 | 0.8069 | 0.9017 | 0.927 | 20000 | 15282 | output/embedquant/lmarena-openai | 0.6 | binary | 1544 |
| 192 | 5395 | lmarena-openai | 1536 | 0.0664 | 0.8852 | 0.7302 | 14605 | 5395 | 428.1 | 2116 | 0.7691 | 0.9091 | 0.8625 | 20000 | 15393 | output/embedquant/lmarena-openai | 0.65 | binary | 1328 |
| 192 | 7116 | lmarena-openai | 1536 | 0.0563 | 0.8293 | 0.6442 | 12884 | 7116 | 428.1 | 3713 | 0.7044 | 0.9126 | 0.76 | 20000 | 15471 | output/embedquant/lmarena-openai | 0.7 | binary | 1126 |
| 192 | 9604 | lmarena-openai | 1536 | 0.0429 | 0.7347 | 0.5198 | 10396 | 9604 | 428.1 | 6031 | 0.6219 | 0.9175 | 0.6126 | 20000 | 15569 | output/embedquant/lmarena-openai | 0.75 | binary | 858 |
| 192 | 12612 | lmarena-openai | 1536 | 0.0296 | 0.5912 | 0.3694 | 7388 | 12612 | 428.1 | 8806 | 0.5602 | 0.9199 | 0.4356 | 20000 | 15602 | output/embedquant/lmarena-openai | 0.8 | binary | 592 |
| 192 | 15050 | lmarena-openai | 1536 | 0.021 | 0.4406 | 0.2475 | 4950 | 15050 | 428.1 | 11084 | 0.5383 | 0.9152 | 0.2901 | 20000 | 15614 | output/embedquant/lmarena-openai | 0.84 | binary | 420 |
| 192 | 16940 | lmarena-openai | 1536 | 0.01345 | 0.299 | 0.153 | 3060 | 16940 | 428.1 | 12820 | 0.5578 | 0.9121 | 0.1788 | 20000 | 15611 | output/embedquant/lmarena-openai | 0.88 | binary | 269 |
| 192 | 17643 | lmarena-openai | 1536 | 0.01085 | 0.2383 | 0.1178 | 2357 | 17643 | 428.1 | 13467 | 0.5879 | 0.9079 | 0.1371 | 20000 | 15607 | output/embedquant/lmarena-openai | 0.9 | binary | 217 |
| 192 | 18146 | lmarena-openai | 1536 | 0.0088 | 0.1922 | 0.0927 | 1854 | 18146 | 428.1 | 13930 | 0.632 | 0.9051 | 0.1075 | 20000 | 15608 | output/embedquant/lmarena-openai | 0.92 | binary | 176 |
| 192 | 18524 | lmarena-openai | 1536 | 0.007 | 0.1563 | 0.0738 | 1476 | 18524 | 428.1 | 14278 | 0.6903 | 0.9051 | 0.08556 | 20000 | 15614 | output/embedquant/lmarena-openai | 0.94 | binary | 140 |
| 192 | 18757 | lmarena-openai | 1536 | 0.00645 | 0.1322 | 0.06215 | 1243 | 18757 | 428.1 | 14502 | 0.7557 | 0.8962 | 0.07134 | 20000 | 15616 | output/embedquant/lmarena-openai | 0.96 | binary | 129 |
| 192 | 18836 | lmarena-openai | 1536 | 0.00605 | 0.1243 | 0.0582 | 1164 | 18836 | 428.1 | 14571 | 0.8202 | 0.896 | 0.0668 | 20000 | 15614 | output/embedquant/lmarena-openai | 0.98 | binary | 121 |
| 3072 | 2619 | lmarena-openai | 768 | 0.2301 | 0.8472 | 0.869 | 17381 | 2619 | 5612 | 8 | 0.8972 | 0.7352 | 0.9994 | 20000 | 12786 | output/embedquant/lmarena-openai | 0.5 | truncate-768 | 4603 |
| 3072 | 2821 | lmarena-openai | 768 | 0.1703 | 0.8895 | 0.8589 | 17179 | 2821 | 5612 | 17 | 0.9355 | 0.8017 | 0.9988 | 20000 | 13790 | output/embedquant/lmarena-openai | 0.55 | truncate-768 | 3406 |
| 3072 | 2975 | lmarena-openai | 768 | 0.1333 | 0.9138 | 0.8512 | 17025 | 2975 | 5612 | 44 | 0.9523 | 0.8434 | 0.9969 | 20000 | 14403 | output/embedquant/lmarena-openai | 0.6 | truncate-768 | 2666 |
| 3072 | 3087 | lmarena-openai | 768 | 0.1128 | 0.926 | 0.8457 | 16913 | 3087 | 5612 | 87 | 0.9573 | 0.8666 | 0.9941 | 20000 | 14744 | output/embedquant/lmarena-openai | 0.65 | truncate-768 | 2256 |
| 3072 | 3217 | lmarena-openai | 768 | 0.09965 | 0.9321 | 0.8391 | 16783 | 3217 | 5612 | 162 | 0.9545 | 0.8812 | 0.9892 | 20000 | 14952 | output/embedquant/lmarena-openai | 0.7 | truncate-768 | 1993 |
| 3072 | 3450 | lmarena-openai | 768 | 0.09095 | 0.9314 | 0.8275 | 16550 | 3450 | 5612 | 352 | 0.9579 | 0.8901 | 0.9767 | 20000 | 15083 | output/embedquant/lmarena-openai | 0.75 | truncate-768 | 1819 |
| 3072 | 3947 | lmarena-openai | 768 | 0.08105 | 0.9227 | 0.8026 | 16053 | 3947 | 5612 | 797 | 0.9568 | 0.899 | 0.9477 | 20000 | 15229 | output/embedquant/lmarena-openai | 0.8 | truncate-768 | 1621 |
| 3072 | 4745 | lmarena-openai | 768 | 0.07135 | 0.9036 | 0.7628 | 15255 | 4745 | 5612 | 1524 | 0.9509 | 0.9065 | 0.9007 | 20000 | 15352 | output/embedquant/lmarena-openai | 0.84 | truncate-768 | 1427 |
| 3072 | 6185 | lmarena-openai | 768 | 0.0608 | 0.8614 | 0.6907 | 13815 | 6185 | 5612 | 2840 | 0.9431 | 0.912 | 0.8161 | 20000 | 15439 | output/embedquant/lmarena-openai | 0.88 | truncate-768 | 1216 |
| 3072 | 7295 | lmarena-openai | 768 | 0.05435 | 0.8239 | 0.6352 | 12705 | 7295 | 5612 | 3880 | 0.9382 | 0.9144 | 0.7496 | 20000 | 15498 | output/embedquant/lmarena-openai | 0.9 | truncate-768 | 1087 |
| 3072 | 8804 | lmarena-openai | 768 | 0.0463 | 0.768 | 0.5598 | 11196 | 8804 | 5612 | 5278 | 0.936 | 0.9173 | 0.6605 | 20000 | 15548 | output/embedquant/lmarena-openai | 0.92 | truncate-768 | 926 |
| 3072 | 10913 | lmarena-openai | 768 | 0.0377 | 0.6758 | 0.4543 | 9087 | 10913 | 5612 | 7242 | 0.9383 | 0.917 | 0.535 | 20000 | 15575 | output/embedquant/lmarena-openai | 0.94 | truncate-768 | 754 |
| 3072 | 13566 | lmarena-openai | 768 | 0.026 | 0.5364 | 0.3217 | 6434 | 13566 | 5612 | 9701 | 0.9403 | 0.9192 | 0.3787 | 20000 | 15615 | output/embedquant/lmarena-openai | 0.96 | truncate-768 | 520 |
| 3072 | 16557 | lmarena-openai | 768 | 0.0154 | 0.3288 | 0.1721 | 3443 | 16557 | 5612 | 12490 | 0.9445 | 0.9105 | 0.2006 | 20000 | 15625 | output/embedquant/lmarena-openai | 0.98 | truncate-768 | 308 |
| 768 | 2615 | lmarena-openai | 768 | 0.2307 | 0.8468 | 0.8692 | 17385 | 2615 | 0 | 8 | 0.898 | 0.7346 | 0.9994 | 20000 | 12779 | output/embedquant/lmarena-openai | 0.5 | truncate-768-int8 | 4614 |
| 768 | 2822 | lmarena-openai | 768 | 0.1701 | 0.8896 | 0.8589 | 17178 | 2822 | 0 | 17 | 0.9351 | 0.802 | 0.9988 | 20000 | 13793 | output/embedquant/lmarena-openai | 0.55 | truncate-768-int8 | 3402 |
| 768 | 2975 | lmarena-openai | 768 | 0.1333 | 0.9138 | 0.8512 | 17025 | 2975 | 0 | 44 | 0.952 | 0.8434 | 0.9969 | 20000 | 14403 | output/embedquant/lmarena-openai | 0.6 | truncate-768-int8 | 2666 |
| 768 | 3088 | lmarena-openai | 768 | 0.1128 | 0.926 | 0.8456 | 16912 | 3088 | 0 | 87 | 0.9572 | 0.8667 | 0.9941 | 20000 | 14744 | output/embedquant/lmarena-openai | 0.65 | truncate-768-int8 | 2255 |
| 768 | 3219 | lmarena-openai | 768 | 0.09965 | 0.9321 | 0.839 | 16781 | 3219 | 0 | 163 | 0.9543 | 0.8812 | 0.9891 | 20000 | 14951 | output/embedquant/lmarena-openai | 0.7 | truncate-768-int8 | 1993 |
| 768 | 3449 | lmarena-openai | 768 | 0.09095 | 0.9314 | 0.8276 | 16551 | 3449 | 0 | 350 | 0.9578 | 0.8901 | 0.9768 | 20000 | 15082 | output/embedquant/lmarena-openai | 0.75 | truncate-768-int8 | 1819 |
| 768 | 3951 | lmarena-openai | 768 | 0.08105 | 0.9226 | 0.8024 | 16049 | 3951 | 0 | 801 | 0.9567 | 0.899 | 0.9474 | 20000 | 15229 | output/embedquant/lmarena-openai | 0.8 | truncate-768-int8 | 1621 |
| 768 | 4747 | lmarena-openai | 768 | 0.07135 | 0.9035 | 0.7627 | 15253 | 4747 | 0 | 1526 | 0.9503 | 0.9064 | 0.9006 | 20000 | 15352 | output/embedquant/lmarena-openai | 0.84 | truncate-768-int8 | 1427 |
| 768 | 6192 | lmarena-openai | 768 | 0.0605 | 0.8613 | 0.6904 | 13808 | 6192 | 0 | 2846 | 0.9436 | 0.9124 | 0.8157 | 20000 | 15444 | output/embedquant/lmarena-openai | 0.88 | truncate-768-int8 | 1210 |
| 768 | 7299 | lmarena-openai | 768 | 0.05435 | 0.8237 | 0.6351 | 12701 | 7299 | 0 | 3884 | 0.9379 | 0.9144 | 0.7494 | 20000 | 15498 | output/embedquant/lmarena-openai | 0.9 | truncate-768-int8 | 1087 |
| 768 | 8817 | lmarena-openai | 768 | 0.04595 | 0.7678 | 0.5592 | 11183 | 8817 | 0 | 5288 | 0.9364 | 0.9178 | 0.66 | 20000 | 15552 | output/embedquant/lmarena-openai | 0.92 | truncate-768-int8 | 919 |
| 768 | 10920 | lmarena-openai | 768 | 0.0377 | 0.6754 | 0.454 | 9080 | 10920 | 0 | 7249 | 0.9389 | 0.917 | 0.5346 | 20000 | 15575 | output/embedquant/lmarena-openai | 0.94 | truncate-768-int8 | 754 |
| 768 | 13576 | lmarena-openai | 768 | 0.02595 | 0.5358 | 0.3212 | 6424 | 13576 | 0 | 9711 | 0.9399 | 0.9192 | 0.3781 | 20000 | 15616 | output/embedquant/lmarena-openai | 0.96 | truncate-768-int8 | 519 |
| 768 | 16566 | lmarena-openai | 768 | 0.01535 | 0.3281 | 0.1717 | 3434 | 16566 | 0 | 12499 | 0.9449 | 0.9106 | 0.2001 | 20000 | 15626 | output/embedquant/lmarena-openai | 0.98 | truncate-768-int8 | 307 |
| 2048 | 2555 | lmarena-openai | 512 | 0.2478 | 0.8343 | 0.8722 | 17445 | 2555 | 4418 | 4 | 0.8488 | 0.7159 | 0.9997 | 20000 | 12493 | output/embedquant/lmarena-openai | 0.5 | truncate-512 | 4956 |
| 2048 | 2793 | lmarena-openai | 512 | 0.1806 | 0.8823 | 0.8603 | 17207 | 2793 | 4418 | 15 | 0.9111 | 0.7901 | 0.9989 | 20000 | 13610 | output/embedquant/lmarena-openai | 0.55 | truncate-512 | 3612 |
| 2048 | 2948 | lmarena-openai | 512 | 0.1398 | 0.9096 | 0.8526 | 17052 | 2948 | 4418 | 39 | 0.9322 | 0.836 | 0.9973 | 20000 | 14295 | output/embedquant/lmarena-openai | 0.6 | truncate-512 | 2796 |
| 2048 | 3062 | lmarena-openai | 512 | 0.1165 | 0.9239 | 0.8469 | 16938 | 3062 | 4418 | 76 | 0.9406 | 0.8624 | 0.9948 | 20000 | 14684 | output/embedquant/lmarena-openai | 0.65 | truncate-512 | 2330 |
| 2048 | 3191 | lmarena-openai | 512 | 0.1012 | 0.9316 | 0.8405 | 16809 | 3191 | 4418 | 146 | 0.9396 | 0.8796 | 0.9902 | 20000 | 14931 | output/embedquant/lmarena-openai | 0.7 | truncate-512 | 2024 |
| 2048 | 3413 | lmarena-openai | 512 | 0.0914 | 0.9322 | 0.8294 | 16587 | 3413 | 4418 | 318 | 0.9442 | 0.8898 | 0.9789 | 20000 | 15077 | output/embedquant/lmarena-openai | 0.75 | truncate-512 | 1828 |
| 2048 | 3887 | lmarena-openai | 512 | 0.082 | 0.924 | 0.8056 | 16113 | 3887 | 4418 | 742 | 0.9414 | 0.8982 | 0.9512 | 20000 | 15215 | output/embedquant/lmarena-openai | 0.8 | truncate-512 | 1640 |
| 2048 | 4636 | lmarena-openai | 512 | 0.07205 | 0.9067 | 0.7682 | 15364 | 4636 | 4418 | 1426 | 0.9336 | 0.9062 | 0.9071 | 20000 | 15349 | output/embedquant/lmarena-openai | 0.84 | truncate-512 | 1441 |
| 2048 | 6020 | lmarena-openai | 512 | 0.06115 | 0.8673 | 0.699 | 13980 | 6020 | 4418 | 2682 | 0.9201 | 0.9125 | 0.8263 | 20000 | 15439 | output/embedquant/lmarena-openai | 0.88 | truncate-512 | 1223 |
| 2048 | 7081 | lmarena-openai | 512 | 0.0556 | 0.8314 | 0.646 | 12919 | 7081 | 4418 | 3675 | 0.9162 | 0.9139 | 0.7626 | 20000 | 15482 | output/embedquant/lmarena-openai | 0.9 | truncate-512 | 1112 |
| 2048 | 8606 | lmarena-openai | 512 | 0.04785 | 0.7751 | 0.5697 | 11394 | 8606 | 4418 | 5101 | 0.9114 | 0.916 | 0.6717 | 20000 | 15538 | output/embedquant/lmarena-openai | 0.92 | truncate-512 | 957 |
| 2048 | 10702 | lmarena-openai | 512 | 0.03855 | 0.6855 | 0.4649 | 9298 | 10702 | 4418 | 7055 | 0.916 | 0.9171 | 0.5472 | 20000 | 15582 | output/embedquant/lmarena-openai | 0.94 | truncate-512 | 771 |
| 2048 | 13377 | lmarena-openai | 512 | 0.02705 | 0.5469 | 0.3311 | 6623 | 13377 | 4418 | 9536 | 0.9188 | 0.9183 | 0.3894 | 20000 | 15618 | output/embedquant/lmarena-openai | 0.96 | truncate-512 | 541 |
| 2048 | 16461 | lmarena-openai | 512 | 0.01575 | 0.3364 | 0.1769 | 3539 | 16461 | 4418 | 12406 | 0.9264 | 0.911 | 0.2063 | 20000 | 15630 | output/embedquant/lmarena-openai | 0.98 | truncate-512 | 315 |
| 512 | 2553 | lmarena-openai | 512 | 0.2484 | 0.8339 | 0.8723 | 17447 | 2553 | 0 | 4 | 0.8494 | 0.7153 | 0.9997 | 20000 | 12483 | output/embedquant/lmarena-openai | 0.5 | truncate-512-int8 | 4968 |
| 512 | 2792 | lmarena-openai | 512 | 0.1806 | 0.8823 | 0.8604 | 17208 | 2792 | 0 | 15 | 0.9112 | 0.7901 | 0.9989 | 20000 | 13611 | output/embedquant/lmarena-openai | 0.55 | truncate-512-int8 | 3612 |
| 512 | 2946 | lmarena-openai | 512 | 0.1398 | 0.9096 | 0.8527 | 17054 | 2946 | 0 | 39 | 0.9322 | 0.8361 | 0.9973 | 20000 | 14297 | output/embedquant/lmarena-openai | 0.6 | truncate-512-int8 | 2796 |
| 512 | 3062 | lmarena-openai | 512 | 0.1167 | 0.9237 | 0.8469 | 16938 | 3062 | 0 | 77 | 0.9406 | 0.8622 | 0.9948 | 20000 | 14681 | output/embedquant/lmarena-openai | 0.65 | truncate-512-int8 | 2334 |
| 512 | 3192 | lmarena-openai | 512 | 0.1012 | 0.9316 | 0.8404 | 16808 | 3192 | 0 | 147 | 0.9394 | 0.8796 | 0.9902 | 20000 | 14931 | output/embedquant/lmarena-openai | 0.7 | truncate-512-int8 | 2024 |
| 512 | 3415 | lmarena-openai | 512 | 0.09145 | 0.9321 | 0.8293 | 16585 | 3415 | 0 | 320 | 0.9439 | 0.8897 | 0.9788 | 20000 | 15076 | output/embedquant/lmarena-openai | 0.75 | truncate-512-int8 | 1829 |
| 512 | 3886 | lmarena-openai | 512 | 0.08205 | 0.9239 | 0.8057 | 16114 | 3886 | 0 | 742 | 0.9413 | 0.8982 | 0.9512 | 20000 | 15215 | output/embedquant/lmarena-openai | 0.8 | truncate-512-int8 | 1641 |
| 512 | 4636 | lmarena-openai | 512 | 0.07205 | 0.9067 | 0.7682 | 15364 | 4636 | 0 | 1426 | 0.9333 | 0.9062 | 0.9071 | 20000 | 15349 | output/embedquant/lmarena-openai | 0.84 | truncate-512-int8 | 1441 |
| 512 | 6023 | lmarena-openai | 512 | 0.0612 | 0.8671 | 0.6988 | 13977 | 6023 | 0 | 2686 | 0.9199 | 0.9124 | 0.826 | 20000 | 15439 | output/embedquant/lmarena-openai | 0.88 | truncate-512-int8 | 1224 |
| 512 | 7087 | lmarena-openai | 512 | 0.0556 | 0.8312 | 0.6456 | 12913 | 7087 | 0 | 3681 | 0.9162 | 0.9139 | 0.7622 | 20000 | 15482 | output/embedquant/lmarena-openai | 0.9 | truncate-512-int8 | 1112 |
| 512 | 8616 | lmarena-openai | 512 | 0.04805 | 0.7744 | 0.5692 | 11384 | 8616 | 0 | 5112 | 0.912 | 0.9156 | 0.6709 | 20000 | 15535 | output/embedquant/lmarena-openai | 0.92 | truncate-512-int8 | 961 |
| 512 | 10718 | lmarena-openai | 512 | 0.03845 | 0.6847 | 0.4641 | 9282 | 10718 | 0 | 7070 | 0.9165 | 0.9172 | 0.5463 | 20000 | 15583 | output/embedquant/lmarena-openai | 0.94 | truncate-512-int8 | 769 |
| 512 | 13388 | lmarena-openai | 512 | 0.027 | 0.5463 | 0.3306 | 6612 | 13388 | 0 | 9544 | 0.9193 | 0.9183 | 0.3888 | 20000 | 15616 | output/embedquant/lmarena-openai | 0.96 | truncate-512-int8 | 540 |
| 512 | 16469 | lmarena-openai | 512 | 0.0158 | 0.3356 | 0.1766 | 3531 | 16469 | 0 | 12413 | 0.9274 | 0.9105 | 0.2057 | 20000 | 15628 | output/embedquant/lmarena-openai | 0.98 | truncate-512-int8 | 316 |
| 1024 | 2420 | lmarena-openai | 256 | 0.2873 | 0.8046 | 0.879 | 17580 | 2420 | 1962 | 2 | 0.7391 | 0.6732 | 0.9998 | 20000 | 11837 | output/embedquant/lmarena-openai | 0.5 | truncate-256 | 5745 |
| 1024 | 2715 | lmarena-openai | 256 | 0.2011 | 0.8679 | 0.8642 | 17285 | 2715 | 1962 | 13 | 0.834 | 0.7673 | 0.999 | 20000 | 13275 | output/embedquant/lmarena-openai | 0.55 | truncate-256 | 4023 |
| 1024 | 2910 | lmarena-openai | 256 | 0.152 | 0.9016 | 0.8545 | 17090 | 2910 | 1962 | 28 | 0.8746 | 0.8221 | 0.998 | 20000 | 14078 | output/embedquant/lmarena-openai | 0.6 | truncate-256 | 3040 |
| 1024 | 3046 | lmarena-openai | 256 | 0.1197 | 0.922 | 0.8477 | 16954 | 3046 | 1962 | 71 | 0.9008 | 0.8588 | 0.9951 | 20000 | 14631 | output/embedquant/lmarena-openai | 0.65 | truncate-256 | 2394 |
| 1024 | 3189 | lmarena-openai | 256 | 0.1037 | 0.9299 | 0.8406 | 16811 | 3189 | 1962 | 146 | 0.8991 | 0.8766 | 0.9902 | 20000 | 14882 | output/embedquant/lmarena-openai | 0.7 | truncate-256 | 2075 |
| 1024 | 3410 | lmarena-openai | 256 | 0.0928 | 0.9313 | 0.8295 | 16590 | 3410 | 1962 | 317 | 0.9074 | 0.8881 | 0.9789 | 20000 | 15051 | output/embedquant/lmarena-openai | 0.75 | truncate-256 | 1856 |
| 1024 | 3859 | lmarena-openai | 256 | 0.08285 | 0.9243 | 0.8071 | 16141 | 3859 | 1962 | 716 | 0.9005 | 0.8973 | 0.9529 | 20000 | 15200 | output/embedquant/lmarena-openai | 0.8 | truncate-256 | 1657 |
| 1024 | 4592 | lmarena-openai | 256 | 0.0735 | 0.9072 | 0.7704 | 15408 | 4592 | 1962 | 1383 | 0.8897 | 0.9046 | 0.9097 | 20000 | 15321 | output/embedquant/lmarena-openai | 0.84 | truncate-256 | 1470 |
| 1024 | 5975 | lmarena-openai | 256 | 0.06235 | 0.8679 | 0.7013 | 14025 | 5975 | 1962 | 2643 | 0.8696 | 0.9111 | 0.8286 | 20000 | 15421 | output/embedquant/lmarena-openai | 0.88 | truncate-256 | 1247 |
| 1024 | 7048 | lmarena-openai | 256 | 0.05495 | 0.8332 | 0.6476 | 12952 | 7048 | 1962 | 3647 | 0.8612 | 0.9151 | 0.7647 | 20000 | 15500 | output/embedquant/lmarena-openai | 0.9 | truncate-256 | 1099 |
| 1024 | 8521 | lmarena-openai | 256 | 0.0474 | 0.7794 | 0.5739 | 11479 | 8521 | 1962 | 5014 | 0.8617 | 0.9174 | 0.6775 | 20000 | 15545 | output/embedquant/lmarena-openai | 0.92 | truncate-256 | 948 |
| 1024 | 10584 | lmarena-openai | 256 | 0.0392 | 0.6907 | 0.4708 | 9416 | 10584 | 1962 | 6948 | 0.858 | 0.9167 | 0.554 | 20000 | 15580 | output/embedquant/lmarena-openai | 0.94 | truncate-256 | 784 |
| 1024 | 13259 | lmarena-openai | 256 | 0.02815 | 0.553 | 0.3371 | 6741 | 13259 | 1962 | 9423 | 0.8651 | 0.9165 | 0.396 | 20000 | 15601 | output/embedquant/lmarena-openai | 0.96 | truncate-256 | 563 |
| 1024 | 16396 | lmarena-openai | 256 | 0.01575 | 0.3421 | 0.1802 | 3604 | 16396 | 1962 | 12334 | 0.8718 | 0.9126 | 0.2105 | 20000 | 15623 | output/embedquant/lmarena-openai | 0.98 | truncate-256 | 315 |
| 256 | 2424 | lmarena-openai | 256 | 0.2843 | 0.807 | 0.8788 | 17576 | 2424 | 0 | 4 | 0.7406 | 0.6765 | 0.9997 | 20000 | 11895 | output/embedquant/lmarena-openai | 0.5 | truncate-256-int8 | 5685 |
| 256 | 2727 | lmarena-openai | 256 | 0.1982 | 0.87 | 0.8637 | 17273 | 2727 | 0 | 13 | 0.8371 | 0.7705 | 0.999 | 20000 | 13322 | output/embedquant/lmarena-openai | 0.55 | truncate-256-int8 | 3964 |
| 256 | 2912 | lmarena-openai | 256 | 0.1505 | 0.9026 | 0.8544 | 17088 | 2912 | 0 | 28 | 0.8783 | 0.8239 | 0.998 | 20000 | 14106 | output/embedquant/lmarena-openai | 0.6 | truncate-256-int8 | 3010 |
| 256 | 3046 | lmarena-openai | 256 | 0.1197 | 0.922 | 0.8477 | 16954 | 3046 | 0 | 71 | 0.9005 | 0.8588 | 0.9951 | 20000 | 14631 | output/embedquant/lmarena-openai | 0.65 | truncate-256-int8 | 2394 |
| 256 | 3195 | lmarena-openai | 256 | 0.1032 | 0.9302 | 0.8403 | 16805 | 3195 | 0 | 149 | 0.8985 | 0.8772 | 0.99 | 20000 | 14890 | output/embedquant/lmarena-openai | 0.7 | truncate-256-int8 | 2064 |
| 256 | 3420 | lmarena-openai | 256 | 0.09245 | 0.9313 | 0.829 | 16580 | 3420 | 0 | 325 | 0.9077 | 0.8885 | 0.9784 | 20000 | 15056 | output/embedquant/lmarena-openai | 0.75 | truncate-256-int8 | 1849 |
| 256 | 3876 | lmarena-openai | 256 | 0.08255 | 0.924 | 0.8062 | 16124 | 3876 | 0 | 730 | 0.9011 | 0.8976 | 0.952 | 20000 | 15203 | output/embedquant/lmarena-openai | 0.8 | truncate-256-int8 | 1651 |
| 256 | 4597 | lmarena-openai | 256 | 0.07345 | 0.9071 | 0.7702 | 15403 | 4597 | 0 | 1386 | 0.8904 | 0.9046 | 0.9095 | 20000 | 15320 | output/embedquant/lmarena-openai | 0.84 | truncate-256-int8 | 1469 |
| 256 | 6005 | lmarena-openai | 256 | 0.06205 | 0.8671 | 0.6997 | 13995 | 6005 | 0 | 2668 | 0.871 | 0.9113 | 0.827 | 20000 | 15422 | output/embedquant/lmarena-openai | 0.88 | truncate-256-int8 | 1241 |
| 256 | 7073 | lmarena-openai | 256 | 0.0548 | 0.8323 | 0.6463 | 12927 | 7073 | 0 | 3670 | 0.8627 | 0.9152 | 0.7632 | 20000 | 15501 | output/embedquant/lmarena-openai | 0.9 | truncate-256-int8 | 1096 |
| 256 | 8566 | lmarena-openai | 256 | 0.04735 | 0.7773 | 0.5717 | 11434 | 8566 | 0 | 5061 | 0.8635 | 0.9172 | 0.6745 | 20000 | 15548 | output/embedquant/lmarena-openai | 0.92 | truncate-256-int8 | 947 |
| 256 | 10627 | lmarena-openai | 256 | 0.0391 | 0.6886 | 0.4687 | 9373 | 10627 | 0 | 6988 | 0.8593 | 0.9166 | 0.5514 | 20000 | 15579 | output/embedquant/lmarena-openai | 0.94 | truncate-256-int8 | 782 |
| 256 | 13295 | lmarena-openai | 256 | 0.02795 | 0.551 | 0.3352 | 6705 | 13295 | 0 | 9456 | 0.8659 | 0.9166 | 0.3939 | 20000 | 15602 | output/embedquant/lmarena-openai | 0.96 | truncate-256-int8 | 559 |
| 256 | 16417 | lmarena-openai | 256 | 0.01565 | 0.3404 | 0.1792 | 3583 | 16417 | 0 | 12357 | 0.8717 | 0.9126 | 0.2093 | 20000 | 15627 | output/embedquant/lmarena-openai | 0.98 | truncate-256-int8 | 313 |
| 512 | 2222 | lmarena-openai | 128 | 0.3442 | 0.7599 | 0.8889 | 17778 | 2222 | 995.8 | 2 | 0.6367 | 0.6128 | 0.9998 | 20000 | 10896 | output/embedquant/lmarena-openai | 0.5 | truncate-128 | 6884 |
| 512 | 2609 | lmarena-openai | 128 | 0.2354 | 0.8432 | 0.8696 | 17391 | 2609 | 995.8 | 8 | 0.7541 | 0.7293 | 0.9994 | 20000 | 12691 | output/embedquant/lmarena-openai | 0.55 | truncate-128 | 4708 |
| 512 | 2885 | lmarena-openai | 128 | 0.1592 | 0.8968 | 0.8558 | 17115 | 2885 | 995.8 | 24 | 0.8235 | 0.814 | 0.9983 | 20000 | 13956 | output/embedquant/lmarena-openai | 0.6 | truncate-128 | 3183 |
| 512 | 3037 | lmarena-openai | 128 | 0.1244 | 0.919 | 0.8481 | 16963 | 3037 | 995.8 | 65 | 0.8636 | 0.8533 | 0.9955 | 20000 | 14540 | output/embedquant/lmarena-openai | 0.65 | truncate-128 | 2488 |
| 512 | 3197 | lmarena-openai | 128 | 0.1027 | 0.9306 | 0.8401 | 16803 | 3197 | 995.8 | 147 | 0.8636 | 0.8778 | 0.9901 | 20000 | 14896 | output/embedquant/lmarena-openai | 0.7 | truncate-128 | 2054 |
| 512 | 3425 | lmarena-openai | 128 | 0.0933 | 0.9305 | 0.8287 | 16575 | 3425 | 995.8 | 332 | 0.8758 | 0.8874 | 0.9779 | 20000 | 15041 | output/embedquant/lmarena-openai | 0.75 | truncate-128 | 1866 |
| 512 | 3886 | lmarena-openai | 128 | 0.08145 | 0.9245 | 0.8057 | 16114 | 3886 | 995.8 | 736 | 0.8602 | 0.8989 | 0.9516 | 20000 | 15221 | output/embedquant/lmarena-openai | 0.8 | truncate-128 | 1629 |
| 512 | 4575 | lmarena-openai | 128 | 0.07285 | 0.9084 | 0.7712 | 15425 | 4575 | 995.8 | 1360 | 0.8527 | 0.9055 | 0.9113 | 20000 | 15328 | output/embedquant/lmarena-openai | 0.84 | truncate-128 | 1457 |
| 512 | 5929 | lmarena-openai | 128 | 0.06205 | 0.87 | 0.7036 | 14071 | 5929 | 995.8 | 2594 | 0.8319 | 0.9118 | 0.8318 | 20000 | 15424 | output/embedquant/lmarena-openai | 0.88 | truncate-128 | 1241 |
| 512 | 6996 | lmarena-openai | 128 | 0.0547 | 0.8358 | 0.6502 | 13004 | 6996 | 995.8 | 3584 | 0.811 | 0.9159 | 0.7687 | 20000 | 15494 | output/embedquant/lmarena-openai | 0.9 | truncate-128 | 1094 |
| 512 | 8453 | lmarena-openai | 128 | 0.0476 | 0.7824 | 0.5774 | 11547 | 8453 | 995.8 | 4940 | 0.8105 | 0.9176 | 0.682 | 20000 | 15535 | output/embedquant/lmarena-openai | 0.92 | truncate-128 | 952 |
| 512 | 10475 | lmarena-openai | 128 | 0.0391 | 0.6969 | 0.4763 | 9525 | 10475 | 995.8 | 6824 | 0.8058 | 0.9179 | 0.5616 | 20000 | 15567 | output/embedquant/lmarena-openai | 0.94 | truncate-128 | 782 |
| 512 | 13112 | lmarena-openai | 128 | 0.0289 | 0.5616 | 0.3444 | 6888 | 13112 | 995.8 | 9275 | 0.8125 | 0.9161 | 0.4049 | 20000 | 15585 | output/embedquant/lmarena-openai | 0.96 | truncate-128 | 578 |
| 512 | 16351 | lmarena-openai | 128 | 0.0158 | 0.3465 | 0.1825 | 3649 | 16351 | 995.8 | 12258 | 0.819 | 0.9134 | 0.2138 | 20000 | 15591 | output/embedquant/lmarena-openai | 0.98 | truncate-128 | 316 |
| 128 | 2304 | lmarena-openai | 128 | 0.3205 | 0.7788 | 0.8848 | 17696 | 2304 | 0 | 3 | 0.6562 | 0.6378 | 0.9997 | 20000 | 11290 | output/embedquant/lmarena-openai | 0.5 | truncate-128-int8 | 6409 |
| 128 | 2675 | lmarena-openai | 128 | 0.2173 | 0.8563 | 0.8662 | 17325 | 2675 | 0 | 8 | 0.7654 | 0.7491 | 0.9994 | 20000 | 12986 | output/embedquant/lmarena-openai | 0.55 | truncate-128-int8 | 4347 |
| 128 | 2913 | lmarena-openai | 128 | 0.1542 | 0.8999 | 0.8544 | 17087 | 2913 | 0 | 31 | 0.8348 | 0.8195 | 0.9978 | 20000 | 14034 | output/embedquant/lmarena-openai | 0.6 | truncate-128-int8 | 3084 |
| 128 | 3071 | lmarena-openai | 128 | 0.1195 | 0.9217 | 0.8465 | 16929 | 3071 | 0 | 81 | 0.8653 | 0.8588 | 0.9945 | 20000 | 14620 | output/embedquant/lmarena-openai | 0.65 | truncate-128-int8 | 2390 |
| 128 | 3231 | lmarena-openai | 128 | 0.1011 | 0.9306 | 0.8385 | 16769 | 3231 | 0 | 179 | 0.8668 | 0.8795 | 0.988 | 20000 | 14927 | output/embedquant/lmarena-openai | 0.7 | truncate-128-int8 | 2021 |
| 128 | 3473 | lmarena-openai | 128 | 0.0916 | 0.9303 | 0.8264 | 16527 | 3473 | 0 | 370 | 0.8754 | 0.8892 | 0.9754 | 20000 | 15065 | output/embedquant/lmarena-openai | 0.75 | truncate-128-int8 | 1832 |
| 128 | 3983 | lmarena-openai | 128 | 0.0814 | 0.9214 | 0.8008 | 16017 | 3983 | 0 | 826 | 0.8588 | 0.8984 | 0.9457 | 20000 | 15215 | output/embedquant/lmarena-openai | 0.8 | truncate-128-int8 | 1628 |
| 128 | 4731 | lmarena-openai | 128 | 0.07145 | 0.9044 | 0.7634 | 15269 | 4731 | 0 | 1498 | 0.8491 | 0.9064 | 0.9023 | 20000 | 15338 | output/embedquant/lmarena-openai | 0.84 | truncate-128-int8 | 1429 |
| 128 | 6170 | lmarena-openai | 128 | 0.06015 | 0.8627 | 0.6915 | 13830 | 6170 | 0 | 2816 | 0.831 | 0.913 | 0.8177 | 20000 | 15443 | output/embedquant/lmarena-openai | 0.88 | truncate-128-int8 | 1203 |
| 128 | 7243 | lmarena-openai | 128 | 0.05335 | 0.8272 | 0.6379 | 12757 | 7243 | 0 | 3817 | 0.8114 | 0.9164 | 0.7539 | 20000 | 15507 | output/embedquant/lmarena-openai | 0.9 | truncate-128-int8 | 1067 |
| 128 | 8705 | lmarena-openai | 128 | 0.04645 | 0.7727 | 0.5647 | 11295 | 8705 | 0 | 5168 | 0.8093 | 0.9178 | 0.6673 | 20000 | 15534 | output/embedquant/lmarena-openai | 0.92 | truncate-128-int8 | 929 |
| 128 | 10792 | lmarena-openai | 128 | 0.0371 | 0.6832 | 0.4604 | 9208 | 10792 | 0 | 7109 | 0.8045 | 0.9194 | 0.5436 | 20000 | 15575 | output/embedquant/lmarena-openai | 0.94 | truncate-128-int8 | 742 |
| 128 | 13407 | lmarena-openai | 128 | 0.0276 | 0.5448 | 0.3296 | 6593 | 13407 | 0 | 9542 | 0.8108 | 0.9163 | 0.3877 | 20000 | 15583 | output/embedquant/lmarena-openai | 0.96 | truncate-128-int8 | 552 |
| 128 | 16475 | lmarena-openai | 128 | 0.01525 | 0.337 | 0.1762 | 3525 | 16475 | 0 | 12367 | 0.8155 | 0.9135 | 0.2066 | 20000 | 15587 | output/embedquant/lmarena-openai | 0.98 | truncate-128-int8 | 305 |
| 256 | 1767 | lmarena-openai | 64 | 0.485 | 0.6375 | 0.9116 | 18233 | 1767 | 337.8 | 2 | 0.455 | 0.4679 | 0.9998 | 20000 | 8534 | output/embedquant/lmarena-openai | 0.5 | truncate-64 | 9701 |
| 256 | 2382 | lmarena-openai | 64 | 0.309 | 0.7869 | 0.8809 | 17618 | 2382 | 337.8 | 12 | 0.6129 | 0.6492 | 0.999 | 20000 | 11449 | output/embedquant/lmarena-openai | 0.55 | truncate-64 | 6181 |
| 256 | 2814 | lmarena-openai | 64 | 0.1862 | 0.878 | 0.8593 | 17186 | 2814 | 337.8 | 20 | 0.7345 | 0.7834 | 0.9985 | 20000 | 13483 | output/embedquant/lmarena-openai | 0.6 | truncate-64 | 3723 |
| 256 | 3041 | lmarena-openai | 64 | 0.1318 | 0.9139 | 0.8479 | 16959 | 3041 | 337.8 | 63 | 0.8063 | 0.8445 | 0.9956 | 20000 | 14385 | output/embedquant/lmarena-openai | 0.65 | truncate-64 | 2637 |
| 256 | 3252 | lmarena-openai | 64 | 0.1047 | 0.9283 | 0.8374 | 16748 | 3252 | 337.8 | 170 | 0.8248 | 0.875 | 0.9885 | 20000 | 14824 | output/embedquant/lmarena-openai | 0.7 | truncate-64 | 2094 |
| 256 | 3508 | lmarena-openai | 64 | 0.09215 | 0.9296 | 0.8246 | 16492 | 3508 | 337.8 | 375 | 0.8309 | 0.8882 | 0.975 | 20000 | 15024 | output/embedquant/lmarena-openai | 0.75 | truncate-64 | 1843 |
| 256 | 3988 | lmarena-openai | 64 | 0.08285 | 0.9211 | 0.8006 | 16012 | 3988 | 337.8 | 801 | 0.8288 | 0.8965 | 0.9471 | 20000 | 15156 | output/embedquant/lmarena-openai | 0.8 | truncate-64 | 1657 |
| 256 | 4726 | lmarena-openai | 64 | 0.07325 | 0.9039 | 0.7637 | 15274 | 4726 | 337.8 | 1471 | 0.8093 | 0.9041 | 0.9037 | 20000 | 15280 | output/embedquant/lmarena-openai | 0.84 | truncate-64 | 1465 |
| 256 | 6091 | lmarena-openai | 64 | 0.06085 | 0.8659 | 0.6955 | 13909 | 6091 | 337.8 | 2715 | 0.7782 | 0.9125 | 0.8238 | 20000 | 15407 | output/embedquant/lmarena-openai | 0.88 | truncate-64 | 1217 |
| 256 | 7122 | lmarena-openai | 64 | 0.0555 | 0.831 | 0.6439 | 12878 | 7122 | 337.8 | 3675 | 0.762 | 0.9138 | 0.762 | 20000 | 15443 | output/embedquant/lmarena-openai | 0.9 | truncate-64 | 1110 |
| 256 | 8558 | lmarena-openai | 64 | 0.04805 | 0.7782 | 0.5721 | 11442 | 8558 | 337.8 | 5012 | 0.7538 | 0.916 | 0.6765 | 20000 | 15493 | output/embedquant/lmarena-openai | 0.92 | truncate-64 | 961 |
| 256 | 10513 | lmarena-openai | 64 | 0.03875 | 0.6964 | 0.4743 | 9487 | 10513 | 337.8 | 6820 | 0.7518 | 0.9183 | 0.5609 | 20000 | 15532 | output/embedquant/lmarena-openai | 0.94 | truncate-64 | 775 |
| 256 | 13200 | lmarena-openai | 64 | 0.02805 | 0.5581 | 0.34 | 6800 | 13200 | 337.8 | 9318 | 0.7574 | 0.9175 | 0.401 | 20000 | 15557 | output/embedquant/lmarena-openai | 0.96 | truncate-64 | 561 |
| 256 | 16356 | lmarena-openai | 64 | 0.0156 | 0.3468 | 0.1822 | 3644 | 16356 | 337.8 | 12238 | 0.762 | 0.9144 | 0.214 | 20000 | 15570 | output/embedquant/lmarena-openai | 0.98 | truncate-64 | 312 |
| 64 | 1852 | lmarena-openai | 64 | 0.4709 | 0.6495 | 0.9074 | 18148 | 1852 | 0 | 4 | 0.4562 | 0.481 | 0.9995 | 20000 | 8733 | output/embedquant/lmarena-openai | 0.5 | truncate-64-int8 | 9419 |
| 64 | 2455 | lmarena-openai | 64 | 0.2932 | 0.7992 | 0.8772 | 17545 | 2455 | 0 | 9 | 0.6219 | 0.6658 | 0.9992 | 20000 | 11691 | output/embedquant/lmarena-openai | 0.55 | truncate-64-int8 | 5863 |
| 64 | 2847 | lmarena-openai | 64 | 0.1802 | 0.8818 | 0.8577 | 17153 | 2847 | 0 | 29 | 0.7398 | 0.7899 | 0.9979 | 20000 | 13578 | output/embedquant/lmarena-openai | 0.6 | truncate-64-int8 | 3604 |
| 64 | 3105 | lmarena-openai | 64 | 0.1262 | 0.9168 | 0.8448 | 16895 | 3105 | 0 | 87 | 0.7998 | 0.8507 | 0.994 | 20000 | 14459 | output/embedquant/lmarena-openai | 0.65 | truncate-64-int8 | 2523 |
| 64 | 3316 | lmarena-openai | 64 | 0.1021 | 0.9289 | 0.8342 | 16684 | 3316 | 0 | 202 | 0.8246 | 0.8777 | 0.9864 | 20000 | 14845 | output/embedquant/lmarena-openai | 0.7 | truncate-64-int8 | 2041 |
| 64 | 3670 | lmarena-openai | 64 | 0.089 | 0.9275 | 0.8165 | 16330 | 3670 | 0 | 494 | 0.8243 | 0.891 | 0.9672 | 20000 | 15044 | output/embedquant/lmarena-openai | 0.75 | truncate-64-int8 | 1780 |
| 64 | 4266 | lmarena-openai | 64 | 0.0785 | 0.9159 | 0.7867 | 15734 | 4266 | 0 | 1030 | 0.8149 | 0.9002 | 0.9322 | 20000 | 15194 | output/embedquant/lmarena-openai | 0.8 | truncate-64-int8 | 1570 |
| 64 | 5099 | lmarena-openai | 64 | 0.0684 | 0.8952 | 0.745 | 14901 | 5099 | 0 | 1799 | 0.7969 | 0.9082 | 0.8827 | 20000 | 15332 | output/embedquant/lmarena-openai | 0.84 | truncate-64-int8 | 1368 |
| 64 | 6613 | lmarena-openai | 64 | 0.0591 | 0.8479 | 0.6694 | 13387 | 6613 | 0 | 3197 | 0.7656 | 0.9117 | 0.7924 | 20000 | 15402 | output/embedquant/lmarena-openai | 0.88 | truncate-64-int8 | 1182 |
| 64 | 7710 | lmarena-openai | 64 | 0.05275 | 0.8101 | 0.6145 | 12290 | 7710 | 0 | 4213 | 0.7508 | 0.9142 | 0.7273 | 20000 | 15448 | output/embedquant/lmarena-openai | 0.9 | truncate-64-int8 | 1055 |
| 64 | 9237 | lmarena-openai | 64 | 0.04595 | 0.75 | 0.5382 | 10763 | 9237 | 0 | 5643 | 0.7395 | 0.9146 | 0.6356 | 20000 | 15487 | output/embedquant/lmarena-openai | 0.92 | truncate-64-int8 | 919 |
| 64 | 11229 | lmarena-openai | 64 | 0.03675 | 0.6615 | 0.4385 | 8771 | 11229 | 0 | 7488 | 0.7392 | 0.9162 | 0.5177 | 20000 | 15524 | output/embedquant/lmarena-openai | 0.94 | truncate-64-int8 | 735 |
| 64 | 13803 | lmarena-openai | 64 | 0.02595 | 0.5223 | 0.3099 | 6197 | 13803 | 0 | 9866 | 0.7461 | 0.9162 | 0.3653 | 20000 | 15544 | output/embedquant/lmarena-openai | 0.96 | truncate-64-int8 | 519 |
| 64 | 16702 | lmarena-openai | 64 | 0.0142 | 0.3197 | 0.1649 | 3298 | 16702 | 0 | 12542 | 0.7513 | 0.9139 | 0.1938 | 20000 | 15556 | output/embedquant/lmarena-openai | 0.98 | truncate-64-int8 | 284 |

## kvcachequant

| benchmark | cached_tokens | completion_tokens | draft_acceptance | errors | latency_s | n | overhead_s | overhead_tokens | prompt_s | prompt_tokens | reasoning_tokens | score | source | tag | task | think | throughput_tokens_per_s | tokens_per_s | total_s | truncated | ttft_s | wall_s | workers |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| kvcachequant | 3.02 | 241.8 | 0 | 0 | 9.767 | 100 | 0 | 0 | 0.12 | 91.81 | 0 | 0.95 | output/kvcachequant/f16/nothink | f16 | gsm8k | false | 97.05 | 25.42 | 9.767 | 0 | 0.2814 | 249.1 | 4 |
| kvcachequant | 3.51 | 253.5 | 0 | 0 | 10.75 | 100 | 0 | 0 | 0.1281 | 124.1 | 0 | 0.72 | output/kvcachequant/f16/nothink | f16 | mmlu | false | 91.14 | 24.26 | 10.75 | 1 | 0.1925 | 278.1 | 4 |
| kvcachequant | 3.01 | 245.1 | 0 | 0 | 10.06 | 100 | 0 | 0 | 0.1232 | 91.81 | 0 | 0.93 | output/kvcachequant/q4_0/nothink | q4_0 | gsm8k | false | 95.74 | 25.01 | 10.06 | 0 | 0.2197 | 256 | 4 |
| kvcachequant | 3.96 | 211.2 | 0 | 0 | 8.811 | 100 | 0 | 0 | 0.1317 | 124.1 | 0 | 0.73 | output/kvcachequant/q4_0/nothink | q4_0 | mmlu | false | 90.88 | 24.76 | 8.811 | 0 | 0.2057 | 232.4 | 4 |
| kvcachequant | 3.02 | 237.5 | 0 | 0 | 9.754 | 100 | 0 | 0 | 0.1251 | 91.81 | 0 | 0.95 | output/kvcachequant/q8_0/nothink | q8_0 | gsm8k | false | 95.4 | 24.97 | 9.754 | 0 | 0.2367 | 248.9 | 4 |
| kvcachequant | 3.97 | 316.5 | 0 | 0 | 13.01 | 100 | 0 | 0 | 0.133 | 124.1 | 0 | 0.72 | output/kvcachequant/q8_0/nothink | q8_0 | mmlu | false | 80.55 | 24.59 | 13.01 | 2 | 0.2079 | 392.9 | 4 |

## prefixcache

| cached_tokens | file_mb | latency_s | mode | n | prompt_s | prompt_tokens | restore_s | save_s | score | source | task | ttft_plus_restore_s | ttft_s |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| 0 | 0 | 6.162 | cold | 30 | 5.311 | 1.348e+04 | 0 | 0 | 0 | output/prefixcache | longbench-hotpotqa | 5.338 | 5.338 |
| 0 | 0 | 5.56 | nocache | 30 | 5.323 | 1.351e+04 | 0 | 0 | 0.6468 | output/prefixcache | longbench-hotpotqa | 5.35 | 5.35 |
| 1.346e+04 | 0 | 0.3547 | warm | 30 | 0.06589 | 1.351e+04 | 0 | 0 | 0.661 | output/prefixcache | longbench-hotpotqa | 0.09233 | 0.09233 |
| 1.346e+04 | 1991 | 0.355 | restored | 30 | 0.06586 | 1.351e+04 | 0.1644 | 1.531 | 0.661 | output/prefixcache | longbench-hotpotqa | 0.2568 | 0.09236 |
| 0 | 0 | 3.941 | cold | 30 | 2.967 | 8105 | 0 | 0 | 0 | output/prefixcache | longbench-2wikimqa | 2.986 | 2.986 |
| 0 | 0 | 3.278 | nocache | 30 | 2.966 | 8134 | 0 | 0 | 0.3572 | output/prefixcache | longbench-2wikimqa | 2.985 | 2.985 |
| 8082 | 0 | 0.369 | warm | 30 | 0.05971 | 8134 | 0 | 0 | 0.3572 | output/prefixcache | longbench-2wikimqa | 0.07863 | 0.07863 |
| 8082 | 1199 | 0.3688 | restored | 30 | 0.05993 | 8134 | 0.1044 | 0.8569 | 0.3572 | output/prefixcache | longbench-2wikimqa | 0.1829 | 0.0785 |
| 0 | 0 | 7.59 | cold | 30 | 6.618 | 1.636e+04 | 0 | 0 | 0 | output/prefixcache | longbench-musique | 6.653 | 6.653 |
| 0 | 0 | 7.096 | nocache | 30 | 6.647 | 1.639e+04 | 0 | 0 | 0.3793 | output/prefixcache | longbench-musique | 6.683 | 6.683 |
| 1.634e+04 | 0 | 0.5419 | warm | 30 | 0.06978 | 1.639e+04 | 0 | 0 | 0.346 | output/prefixcache | longbench-musique | 0.104 | 0.104 |
| 1.634e+04 | 2415 | 0.5422 | restored | 30 | 0.06975 | 1.639e+04 | 0.2017 | 1.669 | 0.346 | output/prefixcache | longbench-musique | 0.306 | 0.1043 |

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
| 5565 | classification | 0.1902 | 0.8306 | 0.8763 | 39435 | 5565 | 4034 | static | 0.783 | 0.8845 | 45000 | 34912 | 53.42 | output/semanticcache/classification | 0.7 | 8557 |
| 12035 | classification | 0.1433 | 0.7751 | 0.7326 | 32965 | 12035 | 8937 | static | 0.8044 | 0.7479 | 45000 | 35453 | 91.33 | output/semanticcache/classification | 0.75 | 6449 |
| 21570 | classification | 0.08309 | 0.6633 | 0.5207 | 23430 | 21570 | 16254 | static | 0.8404 | 0.5478 | 45000 | 35945 | 134.2 | output/semanticcache/classification | 0.8 | 3739 |
| 28975 | classification | 0.041 | 0.5433 | 0.3561 | 16025 | 28975 | 21995 | static | 0.8849 | 0.392 | 45000 | 36175 | 161.4 | output/semanticcache/classification | 0.84 | 1845 |
| 31742 | classification | 0.0258 | 0.4893 | 0.2946 | 13258 | 31742 | 24089 | static | 0.9124 | 0.3343 | 45000 | 36186 | 162.5 | output/semanticcache/classification | 0.86 | 1161 |
| 33791 | classification | 0.01533 | 0.4437 | 0.2491 | 11209 | 33791 | 25682 | static | 0.9384 | 0.2906 | 45000 | 36201 | 174.9 | output/semanticcache/classification | 0.88 | 690 |
| 35155 | classification | 0.008644 | 0.4106 | 0.2188 | 9845 | 35155 | 26764 | static | 0.9605 | 0.2611 | 45000 | 36220 | 177.9 | output/semanticcache/classification | 0.9 | 389 |
| 36113 | classification | 0.005067 | 0.3837 | 0.1975 | 8887 | 36113 | 27585 | static | 0.9743 | 0.2389 | 45000 | 36244 | 173.5 | output/semanticcache/classification | 0.92 | 228 |
| 36791 | classification | 0.0032 | 0.3628 | 0.1824 | 8209 | 36791 | 28185 | static | 0.9825 | 0.2225 | 45000 | 36250 | 182 | output/semanticcache/classification | 0.94 | 144 |
| 37062 | classification | 0.002889 | 0.3533 | 0.1764 | 7938 | 37062 | 28454 | static | 0.9836 | 0.2153 | 45000 | 36262 | 183.4 | output/semanticcache/classification | 0.95 | 130 |
| 37375 | classification | 0.002333 | 0.3427 | 0.1694 | 7625 | 37375 | 28746 | static | 0.9862 | 0.2074 | 45000 | 36266 | 182.2 | output/semanticcache/classification | 0.96 | 105 |
| 37684 | classification | 0.0016 | 0.3323 | 0.1626 | 7316 | 37684 | 29038 | static | 0.9902 | 0.1997 | 45000 | 36282 | 184.3 | output/semanticcache/classification | 0.97 | 72 |
| 37981 | classification | 0.001 | 0.322 | 0.156 | 7019 | 37981 | 29318 | static | 0.9936 | 0.1922 | 45000 | 36292 | 185.2 | output/semanticcache/classification | 0.98 | 45 |
| 38283 | classification | 0.0006222 | 0.311 | 0.1493 | 6717 | 38283 | 29609 | static | 0.9958 | 0.1843 | 45000 | 36298 | 184 | output/semanticcache/classification | 0.99 | 28 |
| 3086 | lmarena | 0.1557 | 0.9103 | 0.9516 | 60710 | 3086 | 70 | static | 0.8364 | 0.9986 | 63796 | 50846 | 63.92 | output/semanticcache/lmarena | 0.7 | 9934 |
| 3438 | lmarena | 0.1156 | 0.9329 | 0.9461 | 60358 | 3438 | 247 | static | 0.8778 | 0.9954 | 63796 | 53232 | 67.18 | output/semanticcache/lmarena | 0.75 | 7373 |
| 3962 | lmarena | 0.09813 | 0.9392 | 0.9379 | 59834 | 3962 | 682 | static | 0.8954 | 0.9874 | 63796 | 54256 | 49.46 | output/semanticcache/lmarena | 0.8 | 6260 |
| 4914 | lmarena | 0.0858 | 0.9384 | 0.923 | 58882 | 4914 | 1540 | static | 0.907 | 0.972 | 63796 | 54948 | 72.03 | output/semanticcache/lmarena | 0.84 | 5474 |
| 5726 | lmarena | 0.07999 | 0.9347 | 0.9102 | 58070 | 5726 | 2301 | static | 0.9121 | 0.9584 | 63796 | 55268 | 79.95 | output/semanticcache/lmarena | 0.86 | 5103 |
| 6995 | lmarena | 0.07461 | 0.9265 | 0.8904 | 56801 | 6995 | 3500 | static | 0.9162 | 0.937 | 63796 | 55541 | 92.4 | output/semanticcache/lmarena | 0.88 | 4760 |
| 8970 | lmarena | 0.06862 | 0.9121 | 0.8594 | 54826 | 8970 | 5346 | static | 0.9201 | 0.9042 | 63796 | 55794 | 106.9 | output/semanticcache/lmarena | 0.9 | 4378 |
| 11939 | lmarena | 0.06135 | 0.8884 | 0.8129 | 51857 | 11939 | 8134 | static | 0.9245 | 0.8549 | 63796 | 56077 | 130.9 | output/semanticcache/lmarena | 0.92 | 3914 |
| 16780 | lmarena | 0.054 | 0.8438 | 0.737 | 47016 | 16780 | 12685 | static | 0.9267 | 0.7745 | 63796 | 56256 | 159.1 | output/semanticcache/lmarena | 0.94 | 3445 |
| 20384 | lmarena | 0.04958 | 0.8073 | 0.6805 | 43412 | 20384 | 16053 | static | 0.9271 | 0.7149 | 63796 | 56302 | 169.7 | output/semanticcache/lmarena | 0.95 | 3163 |
| 24939 | lmarena | 0.04355 | 0.7579 | 0.6091 | 38857 | 24939 | 20275 | static | 0.9285 | 0.6402 | 63796 | 56354 | 213.6 | output/semanticcache/lmarena | 0.96 | 2778 |
| 30778 | lmarena | 0.03706 | 0.6859 | 0.5176 | 33018 | 30778 | 25707 | static | 0.9284 | 0.5439 | 63796 | 56361 | 235.6 | output/semanticcache/lmarena | 0.97 | 2364 |
| 38251 | lmarena | 0.02851 | 0.579 | 0.4004 | 25545 | 38251 | 32684 | static | 0.9288 | 0.4206 | 63796 | 56410 | 264 | output/semanticcache/lmarena | 0.98 | 1819 |
| 47140 | lmarena | 0.0195 | 0.4217 | 0.2611 | 16656 | 47140 | 41023 | static | 0.9253 | 0.2731 | 63796 | 56435 | 316.7 | output/semanticcache/lmarena | 0.99 | 1244 |

## specdecode

| benchmark | cached_tokens | completion_tokens | draft_acceptance | errors | latency_s | n | overhead_s | overhead_tokens | prompt_s | prompt_tokens | reasoning_tokens | score | source | tag | task | think | throughput_tokens_per_s | tokens_per_s | total_s | truncated | ttft_s | wall_s | workers |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| specdecode | 3.025 | 239.9 | 0.4119 | 0 | 3.38 | 40 | 0 | 0 | 0.06166 | 91.3 | 0 | 0.95 | output/specdecode/draft0.6B-n16/nothink | draft0.6B-n16 | gsm8k | false | 70.98 | 76.72 | 3.38 | 0 | 0.08195 | 135.2 | 1 |
| specdecode | 5.05 | 105 | 0.1177 | 0 | 4.427 | 40 | 0 | 0 | 0.3195 | 875.8 | 0 | 0.1965 | output/specdecode/draft0.6B-n16/nothink | draft0.6B-n16 | cnndm | false | 23.72 | 26.3 | 4.427 | 0 | 0.3772 | 177.1 | 1 |
| specdecode | 17.9 | 214.2 | 0.6567 | 0 | 2.096 | 40 | 0 | 0 | 0.07461 | 161.2 | 0 | 0 | output/specdecode/draft0.6B-n16/nothink | draft0.6B-n16 | humaneval | false | 102.2 | 113.9 | 2.096 | 2 | 0.1012 | 83.86 | 1 |
| specdecode | 3.067 | 1385 | 0.3049 | 0 | 26.8 | 15 | 0 | 0 | 0.06112 | 86.13 | 1190 | 0.6667 | output/specdecode/draft0.6B-n16/think | draft0.6B-n16 | gsm8k | true | 51.66 | 55.61 | 26.8 | 5 | 0.2592 | 402 | 1 |
| specdecode | 3.025 | 240.1 | 0.604 | 0 | 3.057 | 40 | 0 | 0 | 0.06161 | 91.3 | 0 | 0.95 | output/specdecode/draft0.6B-n8/nothink | draft0.6B-n8 | gsm8k | false | 78.52 | 81.53 | 3.057 | 1 | 0.08231 | 122.3 | 1 |
| specdecode | 5.05 | 104.4 | 0.2213 | 0 | 3.28 | 40 | 0 | 0 | 0.3202 | 875.8 | 0 | 0.1982 | output/specdecode/draft0.6B-n8/nothink | draft0.6B-n8 | cnndm | false | 31.83 | 36.29 | 3.28 | 0 | 0.3774 | 131.2 | 1 |
| specdecode | 17.9 | 214.8 | 0.7933 | 0 | 2.261 | 40 | 0 | 0 | 0.07436 | 161.2 | 0 | 0 | output/specdecode/draft0.6B-n8/nothink | draft0.6B-n8 | humaneval | false | 94.99 | 101.8 | 2.261 | 2 | 0.1019 | 90.43 | 1 |
| specdecode | 3.067 | 1384 | 0.51 | 0 | 21.07 | 15 | 0 | 0 | 0.06035 | 86.13 | 1115 | 0.8 | output/specdecode/draft0.6B-n8/think | draft0.6B-n8 | gsm8k | true | 65.66 | 67.8 | 21.07 | 3 | 0.2271 | 316.1 | 1 |
| specdecode | 3.025 | 240.7 | 0.5039 | 0 | 5.096 | 40 | 0 | 0 | 0.06198 | 91.3 | 0 | 0.95 | output/specdecode/draft1.7B-n16/nothink | draft1.7B-n16 | gsm8k | false | 47.23 | 49.33 | 5.096 | 1 | 0.08535 | 203.9 | 1 |
| specdecode | 5.05 | 105 | 0.1727 | 0 | 5.834 | 40 | 0 | 0 | 0.3383 | 875.8 | 0 | 0.1965 | output/specdecode/draft1.7B-n16/nothink | draft1.7B-n16 | cnndm | false | 18.01 | 19.66 | 5.834 | 0 | 0.3993 | 233.4 | 1 |
| specdecode | 17.9 | 213.9 | 0.7419 | 0 | 3.263 | 40 | 0 | 0 | 0.07415 | 161.2 | 0 | 0 | output/specdecode/draft1.7B-n16/nothink | draft1.7B-n16 | humaneval | false | 65.57 | 69.44 | 3.263 | 2 | 0.104 | 130.5 | 1 |
| specdecode | 3.067 | 1289 | 0.3927 | 0 | 34.82 | 15 | 0 | 0 | 0.06047 | 86.13 | 1067 | 0.7333 | output/specdecode/draft1.7B-n16/think | draft1.7B-n16 | gsm8k | true | 37.01 | 39.4 | 34.82 | 4 | 0.343 | 522.3 | 1 |
| specdecode | 3.025 | 240.1 | 0.6887 | 0 | 4.39 | 40 | 0 | 0 | 0.06196 | 91.3 | 0 | 0.95 | output/specdecode/draft1.7B-n8/nothink | draft1.7B-n8 | gsm8k | false | 54.68 | 55.59 | 4.39 | 1 | 0.08508 | 175.6 | 1 |
| specdecode | 5.05 | 104.4 | 0.3106 | 0 | 4.035 | 40 | 0 | 0 | 0.3384 | 875.8 | 0 | 0.1977 | output/specdecode/draft1.7B-n8/nothink | draft1.7B-n8 | cnndm | false | 25.87 | 28.99 | 4.035 | 0 | 0.3968 | 161.4 | 1 |
| specdecode | 17.9 | 214.8 | 0.858 | 0 | 3.332 | 40 | 0 | 0 | 0.0747 | 161.2 | 0 | 0 | output/specdecode/draft1.7B-n8/nothink | draft1.7B-n8 | humaneval | false | 64.46 | 66.88 | 3.332 | 2 | 0.1033 | 133.3 | 1 |
| specdecode | 3.067 | 1420 | 0.5754 | 0 | 31.02 | 15 | 0 | 0 | 0.06085 | 86.13 | 1197 | 0.6667 | output/specdecode/draft1.7B-n8/think | draft1.7B-n8 | gsm8k | true | 45.75 | 47.77 | 31.02 | 5 | 0.2745 | 465.4 | 1 |
| specdecode | 3.025 | 240.9 | 0.06095 | 0 | 8.674 | 40 | 0 | 0 | 0.05914 | 91.3 | 0 | 0.925 | output/specdecode/ngram-simple/nothink | ngram-simple | gsm8k | false | 27.77 | 27.89 | 8.674 | 1 | 0.07191 | 347 | 1 |
| specdecode | 5.05 | 103 | 0.04861 | 0 | 4.085 | 40 | 0 | 0 | 0.2967 | 875.8 | 0 | 0.204 | output/specdecode/ngram-simple/nothink | ngram-simple | cnndm | false | 25.22 | 27.16 | 4.085 | 0 | 0.3272 | 163.4 | 1 |
| specdecode | 17.9 | 214.9 | 0.7715 | 0 | 3.759 | 40 | 0 | 0 | 0.0716 | 161.2 | 0 | 0 | output/specdecode/ngram-simple/nothink | ngram-simple | humaneval | false | 57.17 | 63.85 | 3.759 | 2 | 0.08764 | 150.4 | 1 |
| specdecode | 3.067 | 1400 | 0.1105 | 0 | 49.68 | 15 | 0 | 0 | 0.05768 | 86.13 | 1169 | 0.8 | output/specdecode/ngram-simple/think | ngram-simple | gsm8k | true | 28.18 | 28.08 | 49.68 | 3 | 0.1852 | 745.2 | 1 |
| specdecode | 3.025 | 241 | 0 | 0 | 8.627 | 40 | 0 | 0 | 0.05923 | 91.3 | 0 | 0.925 | output/specdecode/none/nothink | none | gsm8k | false | 27.93 | 28.07 | 8.627 | 1 | 0.07309 | 345.1 | 1 |
| specdecode | 5.05 | 103.1 | 0 | 0 | 4.045 | 40 | 0 | 0 | 0.2967 | 875.8 | 0 | 0.2043 | output/specdecode/none/nothink | none | cnndm | false | 25.48 | 27.46 | 4.045 | 0 | 0.3265 | 161.8 | 1 |
| specdecode | 17.9 | 214.7 | 0 | 0 | 7.724 | 40 | 0 | 0 | 0.07091 | 161.2 | 0 | 0 | output/specdecode/none/nothink | none | humaneval | false | 27.8 | 28.03 | 7.724 | 2 | 0.08693 | 308.9 | 1 |
| specdecode | 3.067 | 1424 | 0 | 0 | 51.66 | 15 | 0 | 0 | 0.05636 | 86.13 | 1223 | 0.6 | output/specdecode/none/think | none | gsm8k | true | 27.56 | 27.68 | 51.66 | 6 | 0.1836 | 774.9 | 1 |

## thinkingbudget

| benchmark | cached_tokens | completion_tokens | draft_acceptance | errors | latency_s | n | overhead_s | overhead_tokens | prompt_s | prompt_tokens | reasoning_tokens | score | source | tag | task | think | throughput_tokens_per_s | tokens_per_s | total_s | truncated | ttft_s | wall_s | workers |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| thinkingbudget | 3.01 | 238 | 0 | 0 | 12.46 | 200 | 0 | 0 | 0.1595 | 92.6 | 0 | 0.955 | output/thinkingbudget/budget-0/think | budget-0 | gsm8k | true | 208.3 | 19.66 | 12.46 | 0 | 0.2444 | 228.5 | 12 |
| thinkingbudget | 3.065 | 1319 | 0 | 0 | 125.1 | 200 | 0 | 0 | 0.2043 | 97.81 | 0 | 0.74 | output/thinkingbudget/budget-0/think | budget-0 | math500 | true | 79.77 | 14.32 | 125.1 | 8 | 0.323 | 3306 | 12 |
| thinkingbudget | 3.065 | 1027 | 0 | 0 | 96.98 | 200 | 0 | 0 | 0.2291 | 215.7 | 0 | 0.625 | output/thinkingbudget/budget-0/think | budget-0 | mmlupro | true | 83.22 | 14.64 | 96.98 | 7 | 2.291 | 2469 | 12 |
| thinkingbudget | 3.02 | 1058 | 0 | 0 | 57.18 | 200 | 0 | 0 | 0.1756 | 88.6 | 823.5 | 0.975 | output/thinkingbudget/budget-1024/think | budget-1024 | gsm8k | true | 213.4 | 18.87 | 57.18 | 0 | 0.4994 | 991.8 | 12 |
| thinkingbudget | 3.015 | 519.5 | 0 | 0 | 31.8 | 200 | 0 | 0 | 0.1771 | 88.6 | 126.9 | 0.97 | output/thinkingbudget/budget-128/think | budget-128 | gsm8k | true | 133.5 | 17.73 | 31.8 | 0 | 2.188 | 778.4 | 12 |
| thinkingbudget | 3.09 | 2554 | 0 | 0 | 235 | 200 | 0 | 0 | 0.2155 | 93.81 | 127 | 0.83 | output/thinkingbudget/budget-128/think | budget-128 | math500 | true | 106.9 | 12.12 | 235 | 5 | 0.6048 | 4777 | 12 |
| thinkingbudget | 3.045 | 1433 | 0 | 0 | 123.2 | 200 | 0 | 0 | 0.2415 | 211.7 | 127 | 0.665 | output/thinkingbudget/budget-128/think | budget-128 | mmlupro | true | 105 | 13.26 | 123.2 | 3 | 1.842 | 2730 | 12 |
| thinkingbudget | 79.12 | 869.6 | 0.3487 | 0 | 92 | 8 | 0 | 0 | 0.05251 | 80.12 | 735.6 | 0.75 | output/thinkingbudget/budget-2048/think | budget-2048 | gsm8k | true | 42.08 | 10.29 | 92 | 2 | 1.693 | 165.3 | 8 |
| thinkingbudget | 3 | 739.9 | 0 | 0 | 45.82 | 200 | 0 | 0 | 0.1824 | 88.6 | 254.7 | 0.97 | output/thinkingbudget/budget-256/think | budget-256 | gsm8k | true | 148.9 | 17.48 | 45.82 | 0 | 1.675 | 993.4 | 12 |
| thinkingbudget | 3.14 | 3319 | 0 | 0 | 287 | 200 | 0 | 0 | 0.2165 | 93.81 | 255 | 0.835 | output/thinkingbudget/budget-256/think | budget-256 | math500 | true | 119.7 | 12.32 | 287 | 6 | 0.629 | 5546 | 12 |
| thinkingbudget | 3.09 | 1477 | 0 | 0 | 130.5 | 200 | 0 | 0 | 0.241 | 211.7 | 255 | 0.69 | output/thinkingbudget/budget-256/think | budget-256 | mmlupro | true | 99.55 | 12.74 | 130.5 | 2 | 1.952 | 2968 | 12 |
| thinkingbudget | 3.015 | 816.7 | 0 | 0 | 46.79 | 200 | 0 | 0 | 0.1804 | 88.6 | 498 | 0.98 | output/thinkingbudget/budget-512/think | budget-512 | gsm8k | true | 153.7 | 18.28 | 46.79 | 0 | 0.5064 | 1063 | 12 |
| thinkingbudget | 3.11 | 3640 | 0 | 0 | 338.5 | 200 | 0 | 0 | 0.2196 | 93.81 | 510.9 | 0.825 | output/thinkingbudget/budget-512/think | budget-512 | math500 | true | 110.4 | 11.69 | 338.5 | 9 | 0.668 | 6594 | 12 |
| thinkingbudget | 3.07 | 1485 | 0 | 0 | 114.7 | 200 | 0 | 0 | 0.2287 | 211.7 | 507.1 | 0.71 | output/thinkingbudget/budget-512/think | budget-512 | mmlupro | true | 119.3 | 14.25 | 114.7 | 1 | 1.758 | 2488 | 12 |

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
| 10704 | classification | 0.005 | 0.0024 | 0.2014 | 0.08776 | 3949 | 41051 | 30347 | vcache | 0.9727 | 0.1123 | 45000 | 34188 | 49.66 | output/vcache/classification | 108 |
| 10606 | classification | 0.01 | 0.005533 | 0.3146 | 0.1482 | 6668 | 38332 | 27726 | vcache | 0.9627 | 0.188 | 45000 | 34145 | 92.4 | output/vcache/classification | 249 |
| 10440 | classification | 0.015 | 0.007644 | 0.4285 | 0.217 | 9767 | 35233 | 24793 | vcache | 0.9648 | 0.2754 | 45000 | 34216 | 97.62 | output/vcache/classification | 344 |
| 10201 | classification | 0.02 | 0.01173 | 0.5439 | 0.3006 | 13526 | 31474 | 21273 | vcache | 0.961 | 0.3793 | 45000 | 34271 | 99.15 | output/vcache/classification | 528 |
| 10057 | classification | 0.03 | 0.01589 | 0.5715 | 0.3266 | 14695 | 30305 | 20248 | vcache | 0.9513 | 0.4084 | 45000 | 34228 | 99.03 | output/vcache/classification | 715 |
| 9555 | classification | 0.05 | 0.02507 | 0.6405 | 0.3961 | 17826 | 27174 | 17619 | vcache | 0.9367 | 0.4866 | 45000 | 34317 | 96.76 | output/vcache/classification | 1128 |
| 9408 | classification | 0.07 | 0.0292 | 0.6518 | 0.4116 | 18522 | 26478 | 17070 | vcache | 0.9291 | 0.502 | 45000 | 34278 | 93.72 | output/vcache/classification | 1314 |
| 8885 | classification | 0.1 | 0.0378 | 0.6892 | 0.4598 | 20689 | 24311 | 15426 | vcache | 0.9178 | 0.5518 | 45000 | 34414 | 85.19 | output/vcache/classification | 1701 |
| 7345 | lmarena | 0.005 | 0.0005016 | 0.2171 | 0.1083 | 6906 | 56890 | 49545 | vcache | 0.9954 | 0.1218 | 63796 | 56419 | 82.58 | output/vcache/lmarena | 32 |
| 7311 | lmarena | 0.01 | 0.001129 | 0.3766 | 0.2065 | 13176 | 50620 | 43309 | vcache | 0.9945 | 0.2323 | 63796 | 56413 | 78.47 | output/vcache/lmarena | 72 |
| 7284 | lmarena | 0.015 | 0.001693 | 0.5008 | 0.2976 | 18986 | 44810 | 37526 | vcache | 0.9943 | 0.3347 | 63796 | 56404 | 84.95 | output/vcache/lmarena | 108 |
| 7259 | lmarena | 0.02 | 0.002242 | 0.5983 | 0.3805 | 24276 | 39520 | 32261 | vcache | 0.9941 | 0.4279 | 63796 | 56394 | 84.54 | output/vcache/lmarena | 143 |
| 7187 | lmarena | 0.03 | 0.003401 | 0.6556 | 0.4361 | 27823 | 35973 | 28786 | vcache | 0.9922 | 0.4895 | 63796 | 56392 | 81.13 | output/vcache/lmarena | 217 |
| 7142 | lmarena | 0.05 | 0.004734 | 0.7152 | 0.4991 | 31842 | 31954 | 24812 | vcache | 0.9905 | 0.5597 | 63796 | 56352 | 85.11 | output/vcache/lmarena | 302 |
| 7031 | lmarena | 0.07 | 0.00674 | 0.7468 | 0.5369 | 34255 | 29541 | 22510 | vcache | 0.9874 | 0.6004 | 63796 | 56335 | 88.98 | output/vcache/lmarena | 430 |
| 6991 | lmarena | 0.1 | 0.008151 | 0.7751 | 0.5717 | 36469 | 27327 | 20336 | vcache | 0.9857 | 0.6387 | 63796 | 56285 | 93.15 | output/vcache/lmarena | 520 |

