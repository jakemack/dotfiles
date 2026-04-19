#!/bin/bash
set -e

# This script installs chezmoi and applies the dotfiles.
# It's intended to be used as a VS Code devcontainer dotfiles install script.

if ! command -v chezmoi &> /dev/null; then
  echo "Installing chezmoi..."
  sh -c "$(curl -fsLS get.chezmoi.io)" -- -b "$HOME/.local/bin"
  export PATH="$HOME/.local/bin:$PATH"
fi

# Apply dotfiles. 
# In a devcontainer, we assume we are running from the cloned repository.
# We use --source to point to the current directory.
# We provide default values for templates to avoid blocking for input.

echo "Applying dotfiles with chezmoi..."
# If we are in a devcontainer, we can pass default data.
chezmoi init --apply --source="$(dirname "$0")" --data "name='Jake Mack',email='jakemack@gmail.com'"
