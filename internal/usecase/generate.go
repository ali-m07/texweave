package usecase

import (
	"github.com/ali-m07/texweave/internal/domain"
)

// Generator orchestrates LaTeX generation via a provider.
type Generator struct {
	Provider domain.Provider
}

// NewGenerator returns a Generator that uses the given provider.
func NewGenerator(p domain.Provider) *Generator {
	return &Generator{Provider: p}
}

// Generate produces LaTeX from the given input.
func (g *Generator) Generate(in domain.GenerateInput) (domain.GenerateResult, error) {
	return g.Provider.Generate(in)
}
