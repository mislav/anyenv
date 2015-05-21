# rbenv, but for anything

This is an implementation of rbenv that aims to be agnostic about what kind of
software it is managing versions for. With names of configuration files and
variables such as `.ruby-version`, `RBENV_VERSION`, etc. being configurable,
this project could in theory manage multiple versions of anything and should be
able to replace rbenv, pyenv, phantomenv, nodenv and other rbenv-inspired projects.

### Build your own version manager

For example, let's say you want to build `pyenv` with this:

1. Clone this project into your GOPATH;

2. Run `make` with appropriate configuration:

    ```sh
    $ PROGRAM_NAME=pyenv PROGRAM_EXECUTABLE=python make
    ```

3. Move the resulting `pyenv` binary somewhere into your PATH;

4. Marvel at how you can now run `pyenv version` and other commands. This binary
   is hardcoded to respect:

  * `.python-version` local files,
  * `PYENV_VERSION`,
  * `PYENV_ROOT`,
  * `PYENV_DIR`.

### A work in progress :construction:

Rbenv commands implemented so far:

- [x] `rbenv`
- [x] `rbenv---version`
- [x] `rbenv-commands`
- [ ] `rbenv-completions`
- [x] `rbenv-exec`
- [x] `rbenv-global`
- [x] `rbenv-help`
- [x] `rbenv-hooks` :warning:
- [x] `rbenv-init`
- [x] `rbenv-local`
- [x] `rbenv-prefix`
- [x] `rbenv-rehash`
- [x] `rbenv-root`
- [x] `rbenv-sh-rehash`
- [x] `rbenv-sh-shell`
- [x] `rbenv-shims`
- [x] `rbenv-version`
- [x] `rbenv-version-file`
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
