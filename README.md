# Jake's Dotfiles

Personal dotfiles managed with [chezmoi](https://www.chezmoi.io/). Designed for a full macOS development environment and lightweight VS Code Devcontainers.

## 🚀 Quick Start

To set up a new machine (macOS or Linux):

```sh
sh -c "$(curl -fsLS get.chezmoi.io)" -- init --apply jakemack
```

*This will install chezmoi, clone this repo, prompt for your git identity, and apply your configurations.*

## ✨ Features

- **Cross-Platform:** One configuration for macOS and Linux (Devcontainers).
- **macOS Automation:** Automatically installs Homebrew, bundles apps via `Brewfile`, and sets macOS system defaults.
- **Devcontainer Ready:** Includes an `install.sh` for VS Code to instantly configure your shell in remote containers.
- **Modular Zsh:** Cleanly separated aliases, functions, and Oh My Zsh customizations.
- **Safe & Private:** Uses templates for sensitive identity data (name/email) and supports a non-tracked `~/.localrc` for secrets.

## 📂 Structure

- `dot_local/bin/`: Custom scripts added to `$PATH`.
- `dot_zsh/aliases/`: Modular shell aliases.
- `dot_zsh/functions/`: Autoloaded Zsh functions.
- `dot_zsh_custom/`: Custom themes and plugins for Oh My Zsh.
- `dot_config/`: XDG-compliant application configurations (e.g., Ghostty).
- `run_onchange_macos-setup.sh.tmpl`: Automated macOS provisioning (Brew, system defaults).
- `archive_go_manager/`: Legacy Go-based dotfile manager (for reference).

## 🛠️ Usage

### Applying Changes
After making changes to files in your home directory, tell chezmoi to track them:
```sh
chezmoi add ~/.zshrc
```
To apply changes from the repository to your machine:
```sh
chezmoi apply
```

### Devcontainers
VS Code will automatically detect the `install.sh` in the root of this repository. When you open a project in a Devcontainer, it will run:
```sh
./install.sh
```
This quickly sets up your Zsh environment without installing heavy macOS dependencies.

### Local Customization
Create a `~/.localrc` file for machine-specific environment variables or secrets. This file is sourced by `.zshrc` but is not tracked by git.

## 📜 Credits

Originally inspired by [Zach Holman's dotfiles](https://github.com/holman/dotfiles).
