package root

import (
	"github.com/spf13/cobra"
	"github.com/texweave/texweave/cmd/texweave/root/generate"
)

// Execute runs the root command.
func Execute() error {
	cmd := NewCmd()
	cmd.AddCommand(generate.NewCmd())
	return cmd.Execute()
}

// NewCmd returns the root command.
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "texweave",
		Short: "Weave content into LaTeX using AI",
		Long:  "TexWeave reads your file and generates a complete LaTeX document via OpenAI or Anthropic.",
	}
	cmd.CompletionOptions.DisableDefaultCmd = false
	return cmd
}
