// Package output writes benchmark results: one JSON summary and one JSONL file with a row per request.
package output

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
)

// JSON writes v as indented JSON, creating parent folders.
func JSON(path string, v any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(b, '\n'), 0o644)
}

// JSONL writes one JSON object per line, creating parent folders.
func JSONL[T any](path string, rows []T) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	enc := json.NewEncoder(w)
	for _, r := range rows {
		if err := enc.Encode(r); err != nil {
			return err
		}
	}
	return w.Flush()
}
