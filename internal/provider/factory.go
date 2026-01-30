package provider

import (
	"fmt"

	"github.com/ali-m07/texweave/internal/config"
	"github.com/ali-m07/texweave/internal/domain"
	"github.com/ali-m07/texweave/internal/provider/anthropic"
	"github.com/ali-m07/texweave/internal/provider/openai"
)

// NewFromConfig returns a domain.Provider based on config (env).
func NewFromConfig(cfg *config.Config) (domain.Provider, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
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
