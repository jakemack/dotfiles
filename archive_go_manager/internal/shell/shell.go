package shell

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// Manager handles shell configuration.
type Manager struct {
	DryRun bool
}

// NewManager creates a new ShellManager.
func NewManager(dryRun bool) *Manager {
	return &Manager{DryRun: dryRun}
}

// InstallOhMyZsh installs Oh My Zsh if not present.
func (m *Manager) InstallOhMyZsh() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	omzPath := filepath.Join(homeDir, ".oh-my-zsh")

	if _, err := os.Stat(omzPath); err == nil {
		fmt.Println("  [SKIP] Oh My Zsh is already installed")
		return nil
	}

	fmt.Println("  [INSTALL] Installing Oh My Zsh...")
	if m.DryRun {
		fmt.Println("  [DRY-RUN] git clone https://github.com/ohmyzsh/ohmyzsh.git " + omzPath)
		return nil
	}

	cmd := exec.Command("git", "clone", "https://github.com/ohmyzsh/ohmyzsh.git", omzPath)
	return cmd.Run()
}

// SetDefaultShell sets Zsh as the default shell.
func (m *Manager) SetDefaultShell() error {
	zshPath, err := exec.LookPath("zsh")
	if err != nil {
		return fmt.Errorf("zsh not found")
	}

	shell := os.Getenv("SHELL")
	if shell == zshPath {
		fmt.Println("  [SKIP] Zsh is already the default shell")
		return nil
	}

	fmt.Println("  [CONFIG] Setting Zsh as default shell...")
	if m.DryRun {
		fmt.Printf("  [DRY-RUN] chsh -s %s\n", zshPath)
		return nil
	}

	cmd := exec.Command("chsh", "-s", zshPath)
	return cmd.Run()
}
