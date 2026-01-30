package config

import (
	"os"
	"strings"
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
		return strings.ToLower(c.Provider)
	}
	return DefaultProvider
}
