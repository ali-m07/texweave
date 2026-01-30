package domain

// Provider is the interface for AI-backed LaTeX generation.
type Provider interface {
	Generate(in GenerateInput) (GenerateResult, error)
	Name() string
}
