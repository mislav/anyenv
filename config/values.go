package config

import (
	"os"
)

var (
	Root            = os.ExpandEnv("$HOME/.rbenv")
	RootEnvName     = "RBENV_ROOT"
	VersionFilename = ".ruby-version"
	VersionEnvName  = "RBENV_VERSION"
)

func VersionEnv() string {
	return os.Getenv(VersionEnvName)
}

func init() {
	customRoot := os.Getenv(RootEnvName)
	if customRoot != "" {
		Root = customRoot
	}
}
