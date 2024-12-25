# af-magic.zsh-theme
# Repo: https://github.com/andyfleming/oh-my-zsh
# Direct Link: https://github.com/andyfleming/oh-my-zsh/blob/master/themes/af-magic.zsh-theme

# Function to determine if the shell is remote
is_remote() {
    [[ -n "$SSH_CLIENT" || -n "$SSH_CONNECTION" || -n "$REMOTEHOST" ]]
}

if [ $UID -eq 0 ]; then NCOLOR="red"; else NCOLOR="green"; fi
local return_code="%(?..%{$fg[red]%}%? ↵%{$reset_color%})"

# primary prompt
# PROMPT='$FG[237]------------------------------------------------------------%{$reset_color%}
# $FG[032]%~\
# $(git_prompt_info) \
# $FG[105]%(?..[rc:%?])%{$reset_color%}\
# $FG[105]%(!.#.»)%{$reset_color%} '
PROMPT='$FG[032]%d\
$(git_prompt_info)\
$FG[237]------------------------------------------------------------%{$reset_color%}
$FG[105]%(?..[rc:%?])%{$reset_color%}\
$FG[105]%(!.#.»)%{$reset_color%} '
PROMPT2='%{$fg[red]%}\ %{$reset_color%}'
RPS1='${return_code}'

# color vars
if is_remote; then
  eval right_prompt_color='$FG[160]'
else
  eval right_prompt_color='$FG[245]'
fi
eval my_orange='$FG[214]'

# right prompt
if type "virtualenv_prompt_info" > /dev/null
then
	RPROMPT='$(virtualenv_prompt_info)$right_prompt_color%n@%m%{$reset_color%}%'
else
	RPROMPT='$right_prompt_color%n@%m%{$reset_color%}%'
fi

# git settings
ZSH_THEME_GIT_PROMPT_PREFIX="$FG[075]($FG[078]"
ZSH_THEME_GIT_PROMPT_CLEAN=""
ZSH_THEME_GIT_PROMPT_DIRTY="$my_orange*%{$reset_color%}"
ZSH_THEME_GIT_PROMPT_SUFFIX="$FG[075])%{$reset_color%}"
DISABLE_UNTRACKED_FILES_DIRTY='true'
