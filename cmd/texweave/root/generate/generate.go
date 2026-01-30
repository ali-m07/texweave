package generate

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/texweave/texweave/internal/config"
	"github.com/texweave/texweave/internal/domain"
	"github.com/texweave/texweave/internal/provider"
	"github.com/texweave/texweave/internal/usecase"
	fileadapter "github.com/texweave/texweave/internal/adapter/file"
)

var (
	flagInput  string
	flagOutput string
	flagTitle  string
	flagAuthor string
	flagClass  string
)

// NewCmd returns the generate command.
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate [input file]",
		Short: "Read a file and generate LaTeX",
		Long:  "Reads the given file (or stdin), sends content to the configured AI provider, and writes LaTeX to stdout or a file.",
		Args:  cobra.MaximumNArgs(1),
		RunE:  run,
	}
	cmd.Flags().StringVarP(&flagInput, "input", "i", "", "Input file path (default: stdin)")
	cmd.Flags().StringVarP(&flagOutput, "output", "o", "", "Output .tex file (default: stdout)")
	cmd.Flags().StringVar(&flagTitle, "title", "", "Document title")
	cmd.Flags().StringVar(&flagAuthor, "author", "", "Document author")
	cmd.Flags().StringVar(&flagClass, "class", "article", "Document class (e.g. article, report)")
	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	cfg := config.Load()
	prov, err := provider.NewFromConfig(cfg)
	if err != nil {
		return fmt.Errorf("provider: %w", err)
	}
	gen := usecase.NewGenerator(prov)

	var content []byte
	if len(args) > 0 {
		content, err = fileadapter.ReadAll(args[0])
	} else if flagInput != "" {
		content, err = fileadapter.ReadAll(flagInput)
	} else {
		content, err = readStdin()
	}
	if err != nil {
		return fmt.Errorf("read input: %w", err)
	}

	in := domain.GenerateInput{
		Content:      string(content),
		DocumentClass: flagClass,
		Title:        flagTitle,
		Author:       flagAuthor,
	}
	res, err := gen.Generate(in)
	if err != nil {
		return fmt.Errorf("generate: %w", err)
	}
	if res.Err != nil {
		return res.Err
	}

	if flagOutput != "" {
		return os.WriteFile(flagOutput, []byte(res.Latex), 0644)
	}
	fmt.Print(res.Latex)
	return nil
}

func readStdin() ([]byte, error) {
	return io.ReadAll(os.Stdin)
}
