# AI Commit Message Generator

A robust, production-ready CLI tool written in Go that uses Ollama API to automatically generate meaningful git commit messages from your staged changes.

## Problem Statement
Writing good commit messages is essential for maintaining a clean history and facilitating code reviews. However, it can be tedious and inconsistent across teams. Developers often resort to generic messages like "fix bug" or "update code" when in a rush.

## Objectives
- **Automation**: Automate the process of reading git diffs and crafting commit messages.
- **Consistency**: Enforce team conventions via custom rules.
- **Intelligence**: Use advanced AI to analyze the *intent* of changes, suggesting splits for complex diffs.
- **Speed**: Built with the Go Standard Library and go-git library for zero external binary dependencies and fast execution.

## Features
- **Smart Diff Analysis**: Reads staged changes and context.
- **Split Suggestions**: Detects if a diff contains multiple logical changes and suggests breaking them down (displayed in Yellow).
- **Custom Rules**: Respects `.git-commit-rules-for-ai` in your repo root for team-specific guidelines.
- **Conventional Commits**: Generates messages in the `<type>(<scope>): <description>` format.
- **No External Dependencies**: Uses the go-git library - no git binary installation required.
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
│   │   ├── client.go           # Git Operations (using go-git library)
│   │   └── client_test.go      # Integration/Unit tests
│   └── app/
│       ├── app.go              # Core Application Logic / Orchestrator
│       └── app_test.go         # Table-Driven Unit Tests (Mocked)
├── go.mod                      # Go Module definition
└── README.md                   # Project Documentation
```

## Quick Start (Recommended)

### Download Pre-built Executable

**No Go installation required!** Download the pre-built binary for your platform from the [latest GitHub Release](https://github.com/YOUR_USERNAME/YOUR_REPO/releases/latest).

| Platform | Archive | Binary Name |
|----------|---------|-------------|
| 🪟 **Windows (64-bit)** | `generate-commit-windows.zip` | `generate-commit.exe` |
| 🪟 **Windows (ARM64)** | `generate-commit-windows-arm64.zip` | `generate-commit.exe` |
| 🍎 **Mac (Apple Silicon)** | `generate-commit-mac-arm64.tar.gz` | `generate-commit` |
| 🍎 **Mac (Intel)** | `generate-commit-mac-intel.tar.gz` | `generate-commit` |
| 🐧 **Linux (64-bit)** | `generate-commit-linux.tar.gz` | `generate-commit` |
| 🐧 **Linux (ARM64)** | `generate-commit-linux-arm64.tar.gz` | `generate-commit` |

### Setup Instructions

1. **Download the release**: Go to the [latest release page](https://github.com/YOUR_USERNAME/YOUR_REPO/releases/latest) and download the archive for your platform.

2. **Extract the archive**:
   - **Windows**: Extract the `.zip` file
   - **Mac/Linux**: Extract the `.tar.gz` file:
     ```bash
     tar -xzf generate-commit-*.tar.gz
     ```

3. **Make it executable** (Mac/Linux only):
   ```bash
   chmod +x generate-commit
   ```

4. **Set your API key**:
   ```bash
   # Mac/Linux
   export OLLAMA_API_KEY="your_api_key_here"
   
   # Windows (Command Prompt)
   set OLLAMA_API_KEY=your_api_key_here
   
   # Windows (PowerShell)
   $env:OLLAMA_API_KEY="your_api_key_here"
   ```

5. **Run the tool**:
   ```bash
   # Mac/Linux
   ./generate-commit
   
   # Windows
   generate-commit.exe
   ```

### Optional: Add to PATH

For easier access, move the binary to your PATH:

**Mac/Linux:**
```bash
sudo mv generate-commit /usr/local/bin/
```

**Windows:**
Move `generate-commit.exe` to a directory in your PATH (e.g., `C:\Windows\System32`)

---

## Build from Source

### Prerequisites
- **Go 1.21+** installed
- **Ollama API Key**: Get one from [Ollama](https://ollama.com)

**Note**: The tool uses the `go-git` library and does not require the `git` binary to be installed on your system.

### Installation Steps

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

1. **Navigate to your git repository**:
   ```bash
   cd /path/to/your/project
   ```

2. **Stage your changes**:
   ```bash
   git add .
   ```

3. **Run the tool**:
   ```bash
   ./generate-commit  # or just 'generate-commit' if in PATH
   ```

4. **Review the AI-generated commit message** and use it for your commit!

### Example Output

**Single commit message (Cyan):**
```
Generating commit message...

feat(auth): add OAuth2 login support
```

**Split suggestion (Yellow):**
```
Generating commit message...

AI Suggestion (Split Changes):
This diff contains multiple logical changes:
1. Authentication module (OAuth2 implementation)
2. Database schema updates (user table)
3. UI changes (login form)

Consider splitting into separate commits for better history.
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
