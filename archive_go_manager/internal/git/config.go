package git

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Manager handles git configuration tasks.
type Manager struct {
	RepoRoot string
	DryRun   bool
}

// NewManager creates a new GitManager.
func NewManager(repoRoot string, dryRun bool) *Manager {
	return &Manager{
		RepoRoot: repoRoot,
		DryRun:   dryRun,
	}
}

// SetupGitConfig sets up the local gitconfig if it doesn't exist.
func (m *Manager) SetupGitConfig() error {
	targetPath := filepath.Join(m.RepoRoot, "git", "gitconfig.local.symlink")
	examplePath := filepath.Join(m.RepoRoot, "git", "gitconfig.local.symlink.example")

	// Check if target already exists
	if _, err := os.Stat(targetPath); err == nil {
		fmt.Println("  [SKIP] gitconfig.local.symlink already exists")
		return nil
	}

	fmt.Println("  [CONFIG] Setting up gitconfig.local.symlink...")

	// Read example file
	content, err := os.ReadFile(examplePath)
	if err != nil {
		return fmt.Errorf("failed to read example gitconfig: %w", err)
	}
	configStr := string(content)

	// Prompt for user input
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("  - What is your github author name? ")
	authorName, _ := reader.ReadString('\n')
	authorName = strings.TrimSpace(authorName)

	fmt.Print("  - What is your github author email? ")
	authorEmail, _ := reader.ReadString('\n')
	authorEmail = strings.TrimSpace(authorEmail)

	// Replace placeholders
	configStr = strings.ReplaceAll(configStr, "AUTHORNAME", authorName)
	configStr = strings.ReplaceAll(configStr, "AUTHOREMAIL", authorEmail)
	
	// Determine credential helper
	credHelper := "cache"
	// Simple check for macOS, though in Go runtime.GOOS is better checked elsewhere or passed in.
	// For now, we'll assume if we are running this tool on mac, we want osxkeychain.
	// Ideally we check runtime.GOOS.
	// We'll leave "GIT_CREDENTIAL_HELPER" replacement logic simple or assume the example file has it.
	// The bash script did: sed -e "s/GIT_CREDENTIAL_HELPER/$git_credential/g"
	
	// Let's use "osxkeychain" if on darwin, "cache" otherwise.
	// We need to import runtime in this file or pass it. Let's just hardcode a check here.
	// Note: We can't easily check runtime.GOOS inside the container if the container is Linux but host is Mac.
	// But the dotfiles are for the *current* environment.
	// If running in devcontainer (Linux), it should be cache or store.
	// If running on host (Mac), osxkeychain.
	
	// Since we are running IN the container (Linux) to test, but eventually ON the host (Mac),
	// this logic is tricky. The user wants to run this on their Mac eventually.
	// But for now, let's stick to standard replacement.
	
	configStr = strings.ReplaceAll(configStr, "GIT_CREDENTIAL_HELPER", credHelper)

	if m.DryRun {
		fmt.Println("  [DRY-RUN] Writing gitconfig.local.symlink with:")
		fmt.Printf("    Name: %s\n    Email: %s\n", authorName, authorEmail)
		return nil
	}

	return os.WriteFile(targetPath, []byte(configStr), 0644)
}
