package config

import (
	"github.com/mislav/everyenv/utils"
	"os"
)

var (
	Root            = "$HOME/.rbenv"
	RootEnvName     = "RBENV_ROOT"
	VersionFilename = ".ruby-version"
	VersionEnvName  = "RBENV_VERSION"
	DirEnvName      = "RBENV_DIR"
	ShellEnvName    = "RBENV_SHELL"
	MainExecutable  = "ruby"
	BuildVersion    = "0.0.0"
)

func VersionEnv() string {
	return os.Getenv(VersionEnvName)
}

func DirEnv() string {
	return os.Getenv(DirEnvName)
}

func ShimsDir() utils.Pathname {
	return utils.NewPathname(Root, "shims")
}

func VersionsDir() utils.Pathname {
	return utils.NewPathname(Root, "versions")
}

func VersionDir(name string) utils.Pathname {
	return VersionsDir().Join(name)
}

func GlobalVersionFile() utils.Pathname {
	return utils.NewPathname(Root, "version")
}

func init() {
	customRoot := os.Getenv(RootEnvName)
	if customRoot != "" {
		Root = customRoot
	} else {
		Root = os.ExpandEnv(Root)
	}
}
