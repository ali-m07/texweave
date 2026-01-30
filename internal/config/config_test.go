package config

import (
	"os"
	"testing"
)

func TestProviderName(t *testing.T) {
	c := &Config{}
	if got := c.ProviderName(); got != DefaultProvider {
		t.Errorf("ProviderName() = %q, want %q", got, DefaultProvider)
	}
	c.Provider = "  OPENAI  "
	if got := c.ProviderName(); got != "openai" {
		t.Errorf("ProviderName() = %q, want openai", got)
	}
}

func TestValidate(t *testing.T) {
	t.Run("openai missing key", func(t *testing.T) {
		c := &Config{Provider: "openai", OpenAIKey: ""}
		if err := c.Validate(); err != ErrMissingOpenAIKey {
			t.Errorf("Validate() = %v, want ErrMissingOpenAIKey", err)
		}
	})
	t.Run("openai with key", func(t *testing.T) {
		c := &Config{Provider: "openai", OpenAIKey: "sk-x"}
		if err := c.Validate(); err != nil {
			t.Errorf("Validate() = %v, want nil", err)
		}
	})
	t.Run("anthropic missing key", func(t *testing.T) {
		c := &Config{Provider: "anthropic", AnthropicKey: ""}
		if err := c.Validate(); err != ErrMissingAnthropicKey {
			t.Errorf("Validate() = %v, want ErrMissingAnthropicKey", err)
		}
	})
	t.Run("anthropic with key", func(t *testing.T) {
		c := &Config{Provider: "anthropic", AnthropicKey: "sk-ant-x"}
		if err := c.Validate(); err != nil {
			t.Errorf("Validate() = %v, want nil", err)
		}
	})
}

func TestLoad(t *testing.T) {
	os.Setenv(EnvProvider, "anthropic")
	os.Setenv(EnvOpenAIKey, "sk-openai")
	os.Setenv(EnvAnthropicKey, "sk-ant")
	defer func() {
		os.Unsetenv(EnvProvider)
		os.Unsetenv(EnvOpenAIKey)
		os.Unsetenv(EnvAnthropicKey)
	}()
	c := Load()
	if c.ProviderName() != "anthropic" {
		t.Errorf("ProviderName() = %q, want anthropic", c.ProviderName())
	}
	if c.OpenAIKey != "sk-openai" || c.AnthropicKey != "sk-ant" {
		t.Errorf("Load() keys not set correctly")
	}
}
