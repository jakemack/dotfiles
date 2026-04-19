package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jakemack/dotfiles/internal/bootstrap"
	"github.com/jakemack/dotfiles/internal/git"
	"github.com/jakemack/dotfiles/internal/shell"
	"github.com/jakemack/dotfiles/internal/symlink"
)

func main() {
	dryRun := flag.Bool("dry-run", false, "Print what would happen without making changes")
	flag.Parse()

	fmt.Println("Dotfiles Manager")
	if *dryRun {
		fmt.Println("Running in DRY-RUN mode")
	}

	// Initialize Managers
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting current directory: %v\n", err)
		os.Exit(1)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting home directory: %v\n", err)
		os.Exit(1)
	}

	symlinkMgr := symlink.NewManager(cwd, homeDir, *dryRun)
	bootstrapMgr := bootstrap.NewManager(*dryRun)
	shellMgr := shell.NewManager(*dryRun)
	gitMgr := git.NewManager(cwd, *dryRun)

	// 1. Symlinks
	if err := symlinkMgr.Install(); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating symlinks: %v\n", err)
	}

	// 2. Bootstrap (Homebrew)
	if err := bootstrapMgr.InstallHomebrew(); err != nil {
		fmt.Fprintf(os.Stderr, "Error installing Homebrew: %v\n", err)
	}
	// Only run bundle if brew is available
	if err := bootstrapMgr.BrewBundle(); err != nil {
		// Don't fail hard, just warn
		fmt.Printf("  [WARN] Brew bundle failed or skipped: %v\n", err)
	}

	// 3. Git Config
	if err := gitMgr.SetupGitConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Error setting up git config: %v\n", err)
	}

	// 4. Shell (OMZ)
	if err := shellMgr.InstallOhMyZsh(); err != nil {
		fmt.Fprintf(os.Stderr, "Error installing Oh My Zsh: %v\n", err)
	}
	if err := shellMgr.SetDefaultShell(); err != nil {
		fmt.Fprintf(os.Stderr, "Error setting default shell: %v\n", err)
	}

	fmt.Println("\nDone!")
}
