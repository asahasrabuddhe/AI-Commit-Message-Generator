# AI Commit Message Generator

A robust, production-ready CLI tool written in Go that uses Ollama API to automatically generate meaningful git commit messages from your staged changes.

## Problem Statement
Writing good commit messages is essential for maintaining a clean history and facilitating code reviews. However, it can be tedious and inconsistent across teams. Developers often resort to generic messages like "fix bug" or "update code" when in a rush.

## Objectives
- **Automation**: Automate the process of reading git diffs and crafting commit messages.
- **Consistency**: Enforce team conventions via custom rules.
- **Intelligence**: Use advanced AI to analyze the *intent* of changes, suggesting splits for complex diffs.
- **Speed**: Built with the Go Standard Library for zero-dependency overhead and fast execution.

## Features
- **Smart Diff Analysis**: Reads staged changes and context.
- **Split Suggestions**: Detects if a diff contains multiple logical changes and suggests breaking them down (displayed in Yellow).
- **Custom Rules**: Respects `.git-commit-rules-for-ai` in your repo root for team-specific guidelines.
- **Conventional Commits**: Generates messages in the `<type>(<scope>): <description>` format.
- **Hexagonal Architecture**: Clean, testable, and maintainable codebase.

## Directory Structure
The project follows a Hexagonal (Ports and Adapters) Architecture:

```
/
├── cmd/
│   └── generate-commit/
│       └── main.go            # Entry point. Wires up dependencies.
├── internal/
│   ├── ai/
│   │   ├── generate_commit_message.go      # Gemini API Client
│   │   └── generate_commit_message_test.go # Unit tests
│   ├── config/
│   │   ├── git_commit_rules.go             # Config Loader (.git-commit-rules-for-ai)
│   │   └── git_commit_rules_test.go        # Unit tests
│   ├── git/
│   │   ├── client.go           # Git Command Wrapper
│   │   └── client_test.go      # Integration/Unit tests
│   └── app/
│       ├── app.go              # Core Application Logic / Orchestrator
│       └── app_test.go         # Table-Driven Unit Tests (Mocked)
├── go.mod                      # Go Module definition
└── README.md                   # Project Documentation
```

## Prerequisites
- **Go 1.21+** installed.
- **Ollama API Key**: Get one from [Ollama](https://ollama.com).

## Installation

1. Clone the repository:
   ```bash
   git clone <repo-url>
   cd ai-commit-message-generator
   ```

2. Build the binary:
   ```bash
   go build -o generate-commit ./cmd/generate-commit
   ```

3. (Optional) Move to your PATH:
   ```bash
   mv generate-commit /usr/local/bin/
   ```

## Usage

1. **Set your API Key**:
   ```bash
   export OLLAMA_API_KEY="your_api_key_here"
   ```

2. **Stage your changes**:
   ```bash
   git add .
   ```

3. **Run the tool**:
   ```bash
   ./generate-commit
   ```

### Custom Rules
To enforce specific rules (e.g., "Mention Jira ID"), create a file named `.git-commit-rules-for-ai` in the root of your repository.

**Example `.git-commit-rules-for-ai`:**
```text
- Always start with a verb (Add, Fix, Update).
- If the change affects the UI, mention it.
- Max 50 characters for the subject line.
```

## Running Tests
Run the comprehensive test suite (Unit + Integration):
```bash
go test -v ./...
```
