package openai

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sashabaranov/go-openai"
	"github.com/ali-m07/texweave/internal/domain"
)

const requestTimeout = 120 * time.Second

const name = "openai"

// Provider implements domain.Provider using OpenAI API.
type Provider struct {
	client *openai.Client
	model  string
}

// New creates an OpenAI-backed provider. API key must be set (e.g. OPENAI_API_KEY).
func New(apiKey string, model string) (*Provider, error) {
	apiKey = strings.TrimSpace(apiKey)
	if apiKey == "" {
		return nil, fmt.Errorf("openai: %s is required", "OPENAI_API_KEY")
	}
	if model == "" {
		model = openai.GPT4o
	}
	return &Provider{
		client: openai.NewClient(apiKey),
		model:  model,
	}, nil
}

// Name returns the provider identifier.
func (p *Provider) Name() string { return name }

// Generate calls the API and returns LaTeX.
func (p *Provider) Generate(in domain.GenerateInput) (domain.GenerateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()
	prompt := buildPrompt(in)
	req := openai.ChatCompletionRequest{
		Model: p.model,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: systemPrompt},
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
		Temperature: 0.2,
	}
	resp, err := p.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return domain.GenerateResult{}, err
	}
	if len(resp.Choices) == 0 {
		return domain.GenerateResult{}, fmt.Errorf("openai: empty response")
	}
	latex := strings.TrimSpace(resp.Choices[0].Message.Content)
	latex = extractLatexBlock(latex)
	return domain.GenerateResult{Latex: latex}, nil
}

const systemPrompt = `You are a LaTeX expert. Given raw text or notes, output only valid, complete LaTeX document.
Rules:
- Output a full document with \documentclass, \begin{document}, \end{document}.
- Use the document class and title/author the user requests when provided.
- No markdown, no explanation—only LaTeX.
- If the input looks like markdown, convert headings to \section, \subsection, lists to itemize/enumerate, code to verbatim or listings as appropriate.
- Keep math in \( \) or \[ \] or equation environment.`

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
	const docBegin = "\\begin{document}"
	const docEnd = "\\end{document}"
	const docClass = "\\documentclass"
	if i := strings.Index(s, docBegin); i >= 0 {
		s = s[i:]
		if j := strings.Index(s, docEnd); j >= 0 {
			return s[:j+len(docEnd)]
		}
		return s
	}
	if i := strings.Index(s, docClass); i >= 0 {
		s = s[i:]
		if j := strings.Index(s, docEnd); j >= 0 {
			return s[:j+len(docEnd)]
		}
	}
	return s
}
