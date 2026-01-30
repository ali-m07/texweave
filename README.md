# TexWeave

**Turn plain text into LaTeX.** A minimal CLI that reads your file (or stdin), sends it to an AI provider, and outputs a complete LaTeX document.

[![Go Reference](https://pkg.go.dev/badge/github.com/texweave/texweave.svg)](https://pkg.go.dev/github.com/texweave/texweave)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

- Clean architecture: domain, use case, providers, adapters
- No secrets in code: API keys via environment variables only
- Two providers: OpenAI (GPT-4o) and Anthropic (Claude)

---

## Install

```bash
go install github.com/texweave/texweave/cmd/texweave@latest
```

Or build from source:

```bash
git clone https://github.com/texweave/texweave.git
cd texweave
make build
./texweave --help
```

---

## Usage

Set one API key (and optionally the provider):

```bash
export OPENAI_API_KEY=sk-...        # for OpenAI (default)
# or
export ANTHROPIC_API_KEY=sk-ant-... # for Anthropic

export TEXWEAVE_PROVIDER=openai     # or anthropic (default: openai)
```

Generate LaTeX from a file:

```bash
texweave generate notes.txt
texweave generate -i draft.md -o output.tex --title "My Paper" --author "Your Name"
```

From stdin:

```bash
cat notes.txt | texweave generate
```

| Flag | Description |
|------|-------------|
| `-i, --input` | Input file (or pass path as argument) |
| `-o, --output` | Output `.tex` file (default: stdout) |
| `--title` | Document title |
| `--author` | Document author |
| `--class` | Document class (default: `article`) |

---

## Project layout

```
texweave/
  cmd/texweave/          CLI entry and commands
  internal/
    domain/              Types and Provider interface
    usecase/             Generate orchestration
    provider/            OpenAI and Anthropic implementations
    config/              Environment-based config
    adapter/file/        File reading
  scripts/
    publish.sh           Maintainer: set module path and remote
  go.mod, Makefile, README.md, LICENSE
```

---

## Author & contact

**Ali Mansouri**

- Email: [ali.mansouri1998@gmail.com](mailto:ali.mansouri1998@gmail.com)
- LinkedIn: [linkedin.com/in/ali-mansouri-a7984215b](https://www.linkedin.com/in/ali-mansouri-a7984215b)

For questions, feedback, or collaboration, reach out via email or LinkedIn.

---

## License

MIT. See [LICENSE](LICENSE).
