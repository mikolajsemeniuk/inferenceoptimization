// Package dataset loads the benchmark datasets from datasets/ into plain structs.
package dataset

import (
	"bufio"
	"encoding/json"
	"math/rand/v2"
	"os"

	"github.com/parquet-go/parquet-go"
)

// Root is the folder created by `make datasets`.
const Root = "datasets/"

// Parquet reads all rows of a parquet file into T (columns picked by `parquet:"name"` tags).
func Parquet[T any](path string) ([]T, error) {
	return parquet.ReadFile[T](path)
}

// JSONL reads one JSON object per line into T.
func JSONL[T any](path string) ([]T, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var rows []T
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1<<20), 1<<28)
	for sc.Scan() {
		var row T
		if err := json.Unmarshal(sc.Bytes(), &row); err != nil {
			return nil, err
		}
		rows = append(rows, row)
	}
	return rows, sc.Err()
}

// Sample returns n rows picked with a fixed seed (all rows, shuffled, if n <= 0 or n > len).
func Sample[T any](rows []T, n int, seed uint64) []T {
	out := append([]T(nil), rows...)
	rand.New(rand.NewPCG(seed, seed)).Shuffle(len(out), func(i, j int) { out[i], out[j] = out[j], out[i] })
	if n > 0 && n < len(out) {
		out = out[:n]
	}
	return out
}
