// Package prefixcache drives llama-server's KV-cache reuse: the in-memory prompt prefix cache
// (cache_prompt) and saving / restoring a slot's KV cache to disk.
package prefixcache

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// SlotResult is what the server reports for a slot save or restore.
type SlotResult struct {
	Tokens  int     `json:"tokens"`  // tokens in the saved / restored KV cache
	Bytes   int     `json:"bytes"`   // size of the cache file
	Seconds float64 `json:"seconds"` // client-side duration of the call
}

func slot(ctx context.Context, base string, id int, action, file string) (SlotResult, error) {
	body, _ := json.Marshal(map[string]any{"filename": file})
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s/slots/%d?action=%s", base, id, action), bytes.NewReader(body))
	if err != nil {
		return SlotResult{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	start := time.Now()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return SlotResult{}, err
	}
	defer res.Body.Close()
	raw, _ := io.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		return SlotResult{}, fmt.Errorf("slot %s: %s: %s", action, res.Status, raw)
	}
	var out struct {
		Saved    int `json:"n_saved"`
		Restored int `json:"n_restored"`
		Written  int `json:"n_written"`
		Read     int `json:"n_read"`
	}
	err = json.Unmarshal(raw, &out)
	return SlotResult{Tokens: out.Saved + out.Restored, Bytes: out.Written + out.Read, Seconds: time.Since(start).Seconds()}, err
}

// Save writes the KV cache of the slot to file (inside the server's --slot-save-path).
func Save(ctx context.Context, base string, id int, file string) (SlotResult, error) {
	return slot(ctx, base, id, "save", file)
}

// Restore loads the KV cache of the slot from file.
func Restore(ctx context.Context, base string, id int, file string) (SlotResult, error) {
	return slot(ctx, base, id, "restore", file)
}

// Erase drops the KV cache of the slot.
func Erase(ctx context.Context, base string, id int) error {
	_, err := slot(ctx, base, id, "erase", "")
	return err
}
