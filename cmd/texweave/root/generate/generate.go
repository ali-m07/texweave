package generate

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/ali-m07/texweave/internal/config"
	"github.com/ali-m07/texweave/internal/domain"
	"github.com/ali-m07/texweave/internal/provider"
	"github.com/ali-m07/texweave/internal/usecase"
	fileadapter "github.com/ali-m07/texweave/internal/adapter/file"
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

	inputPath := flagInput
	if len(args) > 0 {
		inputPath = args[0]
	}
	var content []byte
	if inputPath != "" {
		content, err = fileadapter.ReadAll(inputPath)
	} else {
		content, err = readStdin()
	}
	if err != nil {
		return fmt.Errorf("read input: %w", err)
	}
	if len(content) == 0 {
		return fmt.Errorf("input is empty")
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

	if flagOutput != "" {
		return os.WriteFile(flagOutput, []byte(res.Latex), 0644)
	}
	fmt.Print(res.Latex)
	return nil
}

func readStdin() ([]byte, error) {
	return io.ReadAll(os.Stdin)
}
