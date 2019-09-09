# init according to man page
if (( $+commands[pyenv] ))
then
  eval "$(pyenv init -)"
  eval "$(pyenv virtualenv-init -)"
fi
