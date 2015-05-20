package commands

import (
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
)

var versionNameHelp = `
Usage: $ProgramName version-name

Show the current version
`

func versionNameCmd(args cli.Args) {
	currentVersion := detectVersion()

	if !currentVersion.IsSystem() {
		versionDir := config.VersionDir(currentVersion.Name)
		if !versionDir.Exists() {
			err := VersionNotFound{currentVersion.Name}
			cli.Errorf("%s: %s\n", args.ProgramName(), err)
			cli.Exit(1)
		}
	}

	cli.Println(currentVersion.Name)
}

func init() {
	cli.Register("version-name", versionNameCmd, versionNameHelp)
}
