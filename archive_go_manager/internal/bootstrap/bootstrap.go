package bootstrap

import (
	"fmt"
	"os/exec"
	"runtime"
)

// Manager handles system bootstrapping tasks.
type Manager struct {
	DryRun bool
}

// NewManager creates a new BootstrapManager.
func NewManager(dryRun bool) *Manager {
	return &Manager{DryRun: dryRun}
}

// InstallHomebrew installs Homebrew if it's not already installed.
func (m *Manager) InstallHomebrew() error {
	if runtime.GOOS != "darwin" && runtime.GOOS != "linux" {
		fmt.Println("  [SKIP] Homebrew installation skipped (not macOS/Linux)")
		return nil
	}

	_, err := exec.LookPath("brew")
	if err == nil {
		fmt.Println("  [SKIP] Homebrew is already installed")
		return nil
	}

	fmt.Println("  [INSTALL] Installing Homebrew...")
	if m.DryRun {
		fmt.Println("  [DRY-RUN] /bin/bash -c \"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\"")
		return nil
	}

	cmd := exec.Command("/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)")
	cmd.Stdout = nil // Connect to stdout/stderr if you want to see output
	cmd.Stderr = nil
	return cmd.Run()
}

// BrewBundle runs 'brew bundle' to install dependencies.
func (m *Manager) BrewBundle() error {
	_, err := exec.LookPath("brew")
	if err != nil {
		return fmt.Errorf("brew not found, cannot run bundle")
	}

	fmt.Println("  [BUNDLE] Running brew bundle...")
	if m.DryRun {
		fmt.Println("  [DRY-RUN] brew bundle")
		return nil
	}

	cmd := exec.Command("brew", "bundle")
	// In a real app, you might want to stream output
	return cmd.Run()
}
