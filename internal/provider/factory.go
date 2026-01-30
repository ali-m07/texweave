package provider

import (
	"fmt"

	"github.com/texweave/texweave/internal/config"
	"github.com/texweave/texweave/internal/domain"
	"github.com/texweave/texweave/internal/provider/anthropic"
	"github.com/texweave/texweave/internal/provider/openai"
)

// NewFromConfig returns a domain.Provider based on config (env).
func NewFromConfig(cfg *config.Config) (domain.Provider, error) {
	switch cfg.ProviderName() {
	case "openai":
		p, err := openai.New(cfg.OpenAIKey, "")
		if err != nil {
			return nil, err
		}
		return p, nil
	case "anthropic":
		p, err := anthropic.New(cfg.AnthropicKey, "")
		if err != nil {
			return nil, err
		}
		return p, nil
	default:
		return nil, fmt.Errorf("unknown provider %q (use openai or anthropic)", cfg.ProviderName())
	}
}
