package domain

// GenerateInput holds content and options for LaTeX generation.
type GenerateInput struct {
	Content     string
	DocumentClass string
	Title       string
	Author      string
}

// GenerateResult is the outcome of a generation request.
type GenerateResult struct {
	Latex string
}
