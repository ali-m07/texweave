package config

import (
	"errors"
	"os"
	"strings"
)

var (
	ErrMissingOpenAIKey    = errors.New("OPENAI_API_KEY is not set")
	ErrMissingAnthropicKey = errors.New("ANTHROPIC_API_KEY is not set")
)

const (
	EnvOpenAIKey    = "OPENAI_API_KEY"
	EnvAnthropicKey = "ANTHROPIC_API_KEY"
	EnvProvider     = "TEXWEAVE_PROVIDER"
	DefaultProvider = "openai"
)

// Config holds runtime configuration (env-based, no secrets in code).
type Config struct {
	Provider     string
	OpenAIKey    string
	AnthropicKey string
}

// Load reads config from environment.
func Load() *Config {
	return &Config{
		Provider:     strings.TrimSpace(os.Getenv(EnvProvider)),
		OpenAIKey:    strings.TrimSpace(os.Getenv(EnvOpenAIKey)),
		AnthropicKey: strings.TrimSpace(os.Getenv(EnvAnthropicKey)),
	}
}

func (c *Config) ProviderName() string {
	if c.Provider != "" {
		return strings.ToLower(strings.TrimSpace(c.Provider))
	}
	return DefaultProvider
}

// Validate returns an error if the selected provider has no API key set.
func (c *Config) Validate() error {
	name := c.ProviderName()
	switch name {
	case "openai":
		if strings.TrimSpace(c.OpenAIKey) == "" {
			return ErrMissingOpenAIKey
		}
	case "anthropic":
		if strings.TrimSpace(c.AnthropicKey) == "" {
			return ErrMissingAnthropicKey
		}
	}
	return nil
}
