# init according to man page
if (( $+commands[pyenv] ))
then
  _evalcache pyenv init --path
  _evalcache pyenv init -
  _evalcache pyenv virtualenv-init -
#  eval "$(pyenv init --path)"
#  eval "$(pyenv init -)"
#  eval "$(pyenv virtualenv-init -)"
fi
