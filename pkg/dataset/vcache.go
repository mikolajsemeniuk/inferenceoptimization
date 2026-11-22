package dataset

import (
	"strconv"
	"strings"

	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/semanticcache"
	"github.com/mikolajsemeniuk/inferenceoptimization/pkg/vector"
)

// The vCache benchmarks ship prompts with precomputed embeddings and LLM responses,
// so cache policies can be compared without running any model (as the vCache authors do).

type lmArena struct {
	ID       int64  `parquet:"id"`
	Prompt   string `parquet:"prompt"`
	Class    int64  `parquet:"ID_Set"`
	Emb      string `parquet:"emb_gte"`
	Response string `parquet:"response_gpt-4o-mini"`
}

type lmArenaOpenAI struct {
	ID       int64  `parquet:"id"`
	Prompt   string `parquet:"prompt"`
	Class    int64  `parquet:"ID_Set"`
	Emb      string `parquet:"emb_text-embedding-3-small"`
	Response string `parquet:"response_gpt-4o-mini"`
}

type classification struct {
	ID       int64  `parquet:"id"`
	Prompt   string `parquet:"prompt"`
	Emb      string `parquet:"emb_gte"`
	Response string `parquet:"response_llama_3_8b"`
}

type searchQueries struct {
	ID       int64  `parquet:"id"`
	Prompt   string `parquet:"prompt"`
	Class    int64  `parquet:"id_set"`
	Emb      string `parquet:"emb_gte"`
	Response string `parquet:"response_llama_3_8b"`
}

// embedding parses "[0.1, 0.2, ...]" into a unit-length vector.
func embedding(s string) []float32 {
	fields := strings.FieldsFunc(s, func(r rune) bool { return r == '[' || r == ']' || r == ',' || r == ' ' || r == '\n' })
	v := make([]float32, len(fields))
	for i, f := range fields {
		x, _ := strconv.ParseFloat(f, 32)
		v[i] = float32(x)
	}
	return vector.Normalize(v)
}

// VCache loads one of the vCache benchmarks: "lmarena", "classification" or "search"
// (GTE embeddings, 1024 dimensions), or "lmarena-openai" (text-embedding-3-small, 1536
// dimensions, Matryoshka-trained so it can be truncated).
// Class is -1 for "classification", where correctness is judged by comparing responses.
func VCache(name string) ([]semanticcache.Query, error) {
	var out []semanticcache.Query
	switch name {
	case "lmarena":
		rows, err := Parquet[lmArena](Root + "vCache/SemBenchmarkLmArena/train.parquet")
		for _, r := range rows {
			out = append(out, semanticcache.Query{ID: int(r.ID), Prompt: r.Prompt, Vec: embedding(r.Emb), Class: int(r.Class), Response: r.Response})
		}
		return out, err
	case "lmarena-openai":
		rows, err := Parquet[lmArenaOpenAI](Root + "vCache/SemBenchmarkLmArena/train.parquet")
		for _, r := range rows {
			out = append(out, semanticcache.Query{ID: int(r.ID), Prompt: r.Prompt, Vec: embedding(r.Emb), Class: int(r.Class), Response: r.Response})
		}
		return out, err
	case "search":
		rows, err := Parquet[searchQueries](Root + "vCache/SemBenchmarkSearchQueries/train.parquet")
		for _, r := range rows {
			out = append(out, semanticcache.Query{ID: int(r.ID), Prompt: r.Prompt, Vec: embedding(r.Emb), Class: int(r.Class), Response: r.Response})
		}
		return out, err
	default:
		rows, err := Parquet[classification](Root + "vCache/SemBenchmarkClassification/train.parquet")
		for _, r := range rows {
			out = append(out, semanticcache.Query{ID: int(r.ID), Prompt: r.Prompt, Vec: embedding(r.Emb), Class: -1, Response: r.Response})
		}
		return out, err
	}
}
