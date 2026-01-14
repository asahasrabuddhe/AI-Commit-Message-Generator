package app

import (
	"errors"
	"fmt"
	"strings"

	"ai-commit-message-generator/internal/ai"
	"ai-commit-message-generator/internal/config"
	"ai-commit-message-generator/internal/git"
)

// App is the main application struct
type App struct {
	Git    git.Client
	Config config.Loader
	AI     ai.Client
}

// NewApp creates a new App
func NewApp(gitClient git.Client, configLoader config.Loader, aiClient ai.Client) *App {
	return &App{
		Git:    gitClient,
		Config: configLoader,
		AI:     aiClient,
	}
}

// Run executes the main logic
func (a *App) Run() error {
	// 1. Pre-flight Checks
	isRepo, err := a.Git.IsInsideRepo()
	if err != nil {
		return fmt.Errorf("failed to check repository status: %w", err)
	}
	if !isRepo {
		return errors.New("not a git repository")
	}

	hasChanges, err := a.Git.HasStagedChanges()
	if err != nil {
		return fmt.Errorf("failed to check for staged changes: %w", err)
	}
	if !hasChanges {
		return errors.New("no staged changes found. Please stage your changes using 'git add'")
	}

	// 2. Custom Rule Injection
	rules, err := a.Config.LoadRules()
	if err != nil {
		fmt.Printf("Warning: failed to load rules: %v. Proceeding without rules.\n", err)
	}

	// 3. Smart Diff Reading
	diff, err := a.Git.GetStagedDiff()
	if err != nil {
		return fmt.Errorf("failed to get diff: %w", err)
	}

	fmt.Println("Generating commit message...")

	// 4. AI Integration
	message, err := a.AI.GenerateCommitMessage(diff, rules)
	if err != nil {
		return fmt.Errorf("failed to generate commit message: %w", err)
	}

	// 5. Output
	// Check if the response suggests splitting (multi-line or specific keywords)
	// Heuristic: If it has multiple lines, it's likely a split suggestion or discussion.
	// Conventional commits are typically single line (subject).
	if strings.Contains(message, "\n") {
		// Output split suggestion in Yellow
		fmt.Println("\n\033[33mAI Suggestion (Split Changes):\033[0m")
		fmt.Println(message)
	} else {
		// Output commit message in Cyan
		fmt.Println("\n\033[36m" + message + "\033[0m")
	}
	
	return nil
}
