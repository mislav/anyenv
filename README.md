# rbenv, but for anything

This is an implementation of rbenv that aims to be agnostic about what kind of
software it is managing versions for. With names of configuration files and
variables such as `.ruby-version`, `RBENV_VERSION`, etc. being configurable,
this project could in theory manage multiple versions of anything and should be
able to replace rbenv, pyenv, phantomenv, nodenv and other rbenv-inspired projects.

:construction: Rbenv commands implemented so far:

- [x] `rbenv`
- [x] `rbenv---version`
- [x] `rbenv-commands`
- [ ] `rbenv-completions`
- [x] `rbenv-exec`
- [x] `rbenv-global`
- [x] `rbenv-help`
- [ ] `rbenv-hooks` :warning:
- [ ] `rbenv-init`
- [x] `rbenv-local`
- [x] `rbenv-prefix`
- [x] `rbenv-rehash`
- [x] `rbenv-root`
- [ ] `rbenv-sh-rehash`
- [ ] `rbenv-sh-shell`
- [x] `rbenv-shims`
- [x] `rbenv-version`
- [ ] `rbenv-version-file`
- [ ] `rbenv-version-file-read`
- [ ] `rbenv-version-file-write`
- [x] `rbenv-version-name`
- [x] `rbenv-version-origin`
- [x] `rbenv-versions`
- [x] `rbenv-whence`
- [x] `rbenv-which`

:warning: Big hurdle to overcome: support rbenv plugin (hook) system that is
right now dependent on sourcing bash scripts.

Stay tuned.
