package symlink

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Manager handles the creation of symlinks.
type Manager struct {
	SourceDir string
	TargetDir string
	DryRun    bool
}

// NewManager creates a new SymlinkManager.
func NewManager(sourceDir, targetDir string, dryRun bool) *Manager {
	return &Manager{
		SourceDir: sourceDir,
		TargetDir: targetDir,
		DryRun:    dryRun,
	}
}

// Install finds all *.symlink files and links them to the target directory.
func (m *Manager) Install() error {
	fmt.Printf("Scanning for .symlink files in %s...\n", m.SourceDir)

	return filepath.Walk(m.SourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasSuffix(path, ".symlink") {
			relPath, err := filepath.Rel(m.SourceDir, path)
			if err != nil {
				return err
			}

			// Determine target name: .filename (without .symlink extension)
			baseName := filepath.Base(path)
			targetName := "." + strings.TrimSuffix(baseName, ".symlink")
			targetPath := filepath.Join(m.TargetDir, targetName)

			if err := m.link(path, targetPath); err != nil {
				return fmt.Errorf("failed to link %s: %w", relPath, err)
			}
		} else if strings.HasSuffix(path, ".configlink") {
			// Handle .configlink -> ~/.config/<parent>/<filename>
			// Example: ghostty/config.configlink -> ~/.config/ghostty/config
			
			relPath, err := filepath.Rel(m.SourceDir, path)
			if err != nil {
				return err
			}

			parentDir := filepath.Base(filepath.Dir(path))
			baseName := filepath.Base(path)
			targetName := strings.TrimSuffix(baseName, ".configlink")
			
			// Target is ~/.config/<parentDir>/<targetName>
			configDir := filepath.Join(m.TargetDir, ".config", parentDir)
			if err := os.MkdirAll(configDir, 0755); err != nil {
				return fmt.Errorf("failed to create config dir %s: %w", configDir, err)
			}
			
			targetPath := filepath.Join(configDir, targetName)

			if err := m.link(path, targetPath); err != nil {
				return fmt.Errorf("failed to link %s: %w", relPath, err)
			}
		}
		return nil
	})
}

func (m *Manager) link(source, target string) error {
	// Check if target exists
	info, err := os.Lstat(target)
	if err == nil {
		// Target exists
		if info.Mode()&os.ModeSymlink != 0 {
			// It's a symlink, check where it points
			dest, err := os.Readlink(target)
			if err != nil {
				return err
			}
			if dest == source {
				fmt.Printf("  [SKIP] %s already points to %s\n", filepath.Base(target), filepath.Base(source))
				return nil
			}
		}

		// It exists but is different (file, dir, or wrong link)
		// For now, we'll just warn and skip in this simple implementation
		// TODO: Implement backup/overwrite logic
		fmt.Printf("  [WARN] %s exists and is not the correct symlink. Skipping.\n", target)
		return nil
	}

	if m.DryRun {
		fmt.Printf("  [DRY-RUN] ln -s %s %s\n", source, target)
		return nil
	}

	fmt.Printf("  [LINK] %s -> %s\n", target, source)
	return os.Symlink(source, target)
}
