<div align="center">

![Contribution snake](https://raw.githubusercontent.com/Platane/snk/output/github-contribution-grid-snake.svg)

<br />

<img src="https://placehold.co/120x120/0d1117/238636?text=TW" width="80" height="80" />

# **TexWeave**

### *Plain text → LaTeX. One command.*

> AI-powered CLI. Read a file, get a full LaTeX document.  
> **OpenAI** · **Anthropic** · Zero config in code.

<br />

[![Go Reference](https://pkg.go.dev/badge/github.com/ali-m07/texweave.svg)](https://pkg.go.dev/github.com/ali-m07/texweave)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/ali-m07/texweave)](https://goreportcard.com/report/github.com/ali-m07/texweave)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/ali-m07/texweave/pulls)
[![Go 1.22+](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go)](https://go.dev/)

<br />

![TexWeave — terminal to LaTeX](https://placehold.co/900x320/0d1117/238636?text=texweave+generate+%E2%9E%9C+LaTeX)

*Snake: [Platane/snk](https://github.com/Platane/snk). Fork it to show your contribution graph. Replace the banner with `docs/demo.gif` for a live terminal demo.*

---

</div>

## Table of contents

| [Demo](#-demo) | [Features](#-features) | [Quick start](#-quick-start) | [Usage](#-usage) | [Config](#-configuration) | [Author](#-author--contact) |
|----------------|------------------------|------------------------------|------------------|---------------------------|----------------------------|

---

## Demo

<div align="center">

**Before** → **After**

</div>

<table>
<tr>
<td width="50%">

**Input** — plain text / markdown

```text
# My Paper
Intro. Some $x^2$ math.
- Item one
- Item two
```

</td>
<td width="50%">

**Output** — full LaTeX

```latex
\documentclass{article}
\begin{document}
\section{My Paper}
Intro. Some \(x^2\) math.
\begin{itemize}
  \item Item one
  \item Item two
\end{itemize}
\end{document}
```

</td>
</tr>
</table>

**In the terminal:**

```
┌─────────────────────────────────────────────────────────┐
│  ~  texweave generate notes.txt -o paper.tex             │
└─────────────────────────────────────────────────────────┘
$ texweave generate notes.txt -o paper.tex
```

---

## Features

<table>
<tr>
<td width="33%" align="center">

**🏗️ Architecture**

Clean layers: domain, use case, providers, adapters. No framework lock-in.

</td>
<td width="33%" align="center">

**🔐 Security**

No API keys in code. Env vars only. Safe to share and fork.

</td>
<td width="33%" align="center">

**🤖 Providers**

OpenAI (GPT-4o) and Anthropic (Claude). Swap with one env var.

</td>
</tr>
<tr>
<td width="33%" align="center">

**⚡ Single binary**

One CLI. File or stdin → LaTeX. Title, author, document class flags.

</td>
<td width="33%" align="center">

**🔄 CI-ready**

GitHub Actions: build and test on every push. Out of the box.

</td>
<td width="33%" align="center">

**📦 Go 1.22+**

Standard library + minimal deps. Fast builds, easy to extend.

</td>
</tr>
</table>

---

## Quick start

**1. Install**

```bash
go install github.com/ali-m07/texweave/cmd/texweave@latest
```

**2. Set API key** (pick one)

```bash
export OPENAI_API_KEY=sk-...           # or
export ANTHROPIC_API_KEY=sk-ant-...
```

**3. Generate**

```bash
texweave generate my-notes.txt -o paper.tex
```

**4. Compile** (optional)

```bash
pdflatex paper.tex
```

---

## Usage

**From file**

```bash
texweave generate notes.txt
texweave generate -i draft.md -o output.tex --title "My Paper" --author "Your Name"
```

**From stdin**

```bash
cat notes.txt | texweave generate -o out.tex
```

**Provider**

```bash
export TEXWEAVE_PROVIDER=anthropic   # default: openai
```

| Flag | Short | Description |
|------|-------|-------------|
| `--input` | `-i` | Input file path |
| `--output` | `-o` | Output `.tex` file (default: stdout) |
| `--title` | | Document title |
| `--author` | | Document author |
| `--class` | | Document class (default: `article`) |

---

## Configuration

| Variable | Description |
|----------|-------------|
| `OPENAI_API_KEY` | Required when `TEXWEAVE_PROVIDER=openai` |
| `ANTHROPIC_API_KEY` | Required when `TEXWEAVE_PROVIDER=anthropic` |
| `TEXWEAVE_PROVIDER` | `openai` or `anthropic` (default: `openai`) |

---

## Project structure

<details>
<summary><b>Click to expand</b></summary>

```text
texweave/
├── cmd/texweave/           # CLI entry and commands
│   ├── main.go
│   └── root/
├── internal/
│   ├── domain/             # Types and Provider interface
│   ├── usecase/            # Generate orchestration
│   ├── provider/           # OpenAI + Anthropic
│   ├── config/             # Env-based config
│   └── adapter/file/       # File reading
├── .github/workflows/      # CI (build + test)
├── scripts/                # Maintainer scripts
├── go.mod
├── Makefile
└── README.md
```

</details>

---

## Author & contact

<div align="center">

**Ali Mansouri**

[![Email](https://img.shields.io/badge/Email-ali.mansouri1998%40gmail.com-ea4335?style=for-the-badge&logo=gmail)](mailto:ali.mansouri1998@gmail.com)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-Ali%20Mansouri-0a66c2?style=for-the-badge&logo=linkedin)](https://www.linkedin.com/in/ali-mansouri-a7984215b)

*Questions, feedback, or collaboration — reach out.*

</div>

---

<div align="center">

**License** — [MIT](LICENSE)

If you use TexWeave, consider giving it a ⭐

</div>
