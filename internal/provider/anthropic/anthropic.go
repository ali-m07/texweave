package anthropic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ali-m07/texweave/internal/domain"
)

const (
	name         = "anthropic"
	defaultModel = "claude-sonnet-4-20250514"
	apiURL       = "https://api.anthropic.com/v1/messages"
	httpTimeout  = 120 * time.Second
)

// Provider implements domain.Provider using Anthropic API.
type Provider struct {
	apiKey string
	model  string
	client *http.Client
}

// New creates an Anthropic-backed provider. API key must be set (e.g. ANTHROPIC_API_KEY).
func New(apiKey string, model string) (*Provider, error) {
	apiKey = strings.TrimSpace(apiKey)
	if apiKey == "" {
		return nil, fmt.Errorf("anthropic: %s is required", "ANTHROPIC_API_KEY")
	}
	if model == "" {
		model = defaultModel
	}
	return &Provider{
		apiKey: apiKey,
		model:  model,
		client: &http.Client{Timeout: httpTimeout},
	}, nil
}

// Name returns the provider identifier.
func (p *Provider) Name() string { return name }

// Generate calls the API and returns LaTeX.
func (p *Provider) Generate(in domain.GenerateInput) (domain.GenerateResult, error) {
	prompt := buildPrompt(in)
	body := map[string]interface{}{
		"model":      p.model,
		"max_tokens": 8192,
		"messages": []map[string]string{
			{"role": "user", "content": systemPrompt + "\n\n" + prompt},
		},
	}
	raw, err := json.Marshal(body)
	if err != nil {
		return domain.GenerateResult{}, fmt.Errorf("anthropic: encode request: %w", err)
	}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, apiURL, bytes.NewReader(raw))
	if err != nil {
		return domain.GenerateResult{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", p.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")
	resp, err := p.client.Do(req)
	if err != nil {
		return domain.GenerateResult{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		if len(body) > 0 {
			return domain.GenerateResult{}, fmt.Errorf("anthropic: %s: %s", resp.Status, string(body))
		}
		return domain.GenerateResult{}, fmt.Errorf("anthropic: %s", resp.Status)
	}
	var out struct {
		Content []struct {
			Text string `json:"text"`
		} `json:"content"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return domain.GenerateResult{}, err
	}
	if len(out.Content) == 0 {
		return domain.GenerateResult{}, fmt.Errorf("anthropic: empty response")
	}
	latex := strings.TrimSpace(out.Content[0].Text)
	latex = extractLatexBlock(latex)
	return domain.GenerateResult{Latex: latex}, nil
}

const systemPrompt = `You are a LaTeX expert. Given raw text or notes, output only valid, complete LaTeX document.
Rules: Output a full document with \documentclass, \begin{document}, \end{document}. Use the document class and title/author the user requests when provided. No markdown, no explanation—only LaTeX. Convert markdown to sections, lists, verbatim as appropriate. Keep math in \( \) or \[ \].`

func buildPrompt(in domain.GenerateInput) string {
	var b strings.Builder
	if in.DocumentClass != "" {
		b.WriteString("Document class: " + in.DocumentClass + "\n\n")
	}
	if in.Title != "" {
		b.WriteString("Title: " + in.Title + "\n\n")
	}
	if in.Author != "" {
		b.WriteString("Author: " + in.Author + "\n\n")
	}
	b.WriteString("Content to convert to LaTeX:\n\n")
	b.WriteString(in.Content)
	return b.String()
}

func extractLatexBlock(s string) string {
	s = strings.TrimSpace(s)
	if i := strings.Index(s, "\\begin{document}"); i >= 0 {
		s = s[i:]
	}
	if j := strings.Index(s, "\\end{document}"); j >= 0 {
		s = s[:j+len("\\end{document}")]
	}
	return s
}
