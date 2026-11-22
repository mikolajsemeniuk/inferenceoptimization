// Package llm is a small client for llama-server's OpenAI-compatible API.
package llm

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"net/http"
	"strings"
	"time"
)

// Message is one chat turn.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Request is a chat completion request. Extra is merged into the JSON body
// (llama-server options such as cache_prompt, id_slot, reasoning_budget, chat_template_kwargs).
type Request struct {
	Messages    []Message
	MaxTokens   int
	Temperature float64
	Seed        int
	Extra       map[string]any
}

// Response is the answer plus everything the server reports about how it was produced.
type Response struct {
	Content          string  `json:"content"`
	Reasoning        string  `json:"reasoning"`
	FinishReason     string  `json:"finish_reason"`
	PromptTokens     int     `json:"prompt_tokens"`
	CompletionTokens int     `json:"completion_tokens"` // reasoning + content
	ReasoningTokens  int     `json:"reasoning_tokens"`  // filled by CountReasoning
	CachedTokens     int     `json:"cached_tokens"`     // prompt tokens reused from the KV cache
	DraftTokens      int     `json:"draft_tokens"`      // speculative decoding: drafted
	DraftAccepted    int     `json:"draft_accepted"`    // speculative decoding: accepted
	TTFT             float64 `json:"ttft_s"`            // client-side time to first token
	Latency          float64 `json:"latency_s"`         // client-side total time
	PromptSeconds    float64 `json:"prompt_s"`          // server-side prompt processing
	PredictSeconds   float64 `json:"predict_s"`         // server-side generation
	TokensPerSecond  float64 `json:"tokens_per_s"`      // server-side generation speed
}

type chunk struct {
	Choices []struct {
		Delta struct {
			Content   string `json:"content"`
			Reasoning string `json:"reasoning_content"`
		} `json:"delta"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage *struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
	} `json:"usage"`
	Timings *struct {
		CacheN        int     `json:"cache_n"`
		PromptMS      float64 `json:"prompt_ms"`
		PredictedMS   float64 `json:"predicted_ms"`
		PerSecond     float64 `json:"predicted_per_second"`
		DraftN        int     `json:"draft_n"`
		DraftAccepted int     `json:"draft_n_accepted"`
	} `json:"timings"`
}

// Think returns the Extra option that switches thinking on or off (Qwen3-style templates).
func Think(on bool) map[string]any {
	return map[string]any{"chat_template_kwargs": map[string]any{"enable_thinking": on}}
}

// User builds a single-turn conversation.
func User(prompt string) []Message { return []Message{{"user", prompt}} }

func post(ctx context.Context, url string, body any) (*http.Response, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		msg, _ := io.ReadAll(res.Body)
		res.Body.Close()
		return nil, fmt.Errorf("%s: %s: %s", url, res.Status, msg)
	}
	return res, nil
}

// Chat sends the request to the server at base (e.g. http://localhost:8101) and streams the answer.
func Chat(ctx context.Context, base string, r Request) (Response, error) {
	body := map[string]any{
		"messages": r.Messages, "temperature": r.Temperature, "seed": r.Seed,
		"stream": true, "stream_options": map[string]any{"include_usage": true}, "timings_per_token": true,
	}
	if r.MaxTokens > 0 {
		body["max_tokens"] = r.MaxTokens
	}
	maps.Copy(body, r.Extra)

	start := time.Now()
	res, err := post(ctx, base+"/v1/chat/completions", body)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	var out Response
	var content, reasoning strings.Builder
	sc := bufio.NewScanner(res.Body)
	sc.Buffer(make([]byte, 1<<20), 1<<26)
	for sc.Scan() {
		line, ok := strings.CutPrefix(sc.Text(), "data: ")
		if !ok || line == "[DONE]" {
			continue
		}
		var c chunk
		if err := json.Unmarshal([]byte(line), &c); err != nil {
			return out, fmt.Errorf("bad chunk %q: %w", line, err)
		}
		for _, ch := range c.Choices {
			if out.TTFT == 0 && ch.Delta.Content+ch.Delta.Reasoning != "" {
				out.TTFT = time.Since(start).Seconds()
			}
			content.WriteString(ch.Delta.Content)
			reasoning.WriteString(ch.Delta.Reasoning)
			if ch.FinishReason != "" {
				out.FinishReason = ch.FinishReason
			}
		}
		if c.Usage != nil {
			out.PromptTokens, out.CompletionTokens = c.Usage.PromptTokens, c.Usage.CompletionTokens
		}
		if t := c.Timings; t != nil {
			out.CachedTokens, out.DraftTokens, out.DraftAccepted = t.CacheN, t.DraftN, t.DraftAccepted
			out.PromptSeconds, out.PredictSeconds, out.TokensPerSecond = t.PromptMS/1000, t.PredictedMS/1000, t.PerSecond
		}
	}
	out.Content, out.Reasoning = strings.TrimSpace(content.String()), strings.TrimSpace(reasoning.String())
	out.Latency = time.Since(start).Seconds()
	return out, sc.Err()
}

// Tokens counts the tokens of text with the server's tokenizer.
func Tokens(ctx context.Context, base, text string) (int, error) {
	if text == "" {
		return 0, nil
	}
	res, err := post(ctx, base+"/tokenize", map[string]any{"content": text})
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	var out struct {
		Tokens []int `json:"tokens"`
	}
	return len(out.Tokens), json.NewDecoder(res.Body).Decode(&out)
}

// Embed returns unit-length embeddings for the texts from an embedding server.
func Embed(ctx context.Context, base string, texts []string) ([][]float32, error) {
	res, err := post(ctx, base+"/v1/embeddings", map[string]any{"input": texts})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var out struct {
		Data []struct {
			Embedding []float32 `json:"embedding"`
		} `json:"data"`
	}
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return nil, err
	}
	vecs := make([][]float32, len(out.Data))
	for i, d := range out.Data {
		vecs[i] = d.Embedding
	}
	return vecs, nil
}
