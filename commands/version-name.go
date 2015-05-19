package commands

import (
	"github.com/mislav/everyenv/cli"
)

var versionNameHelp = `
Usage: $ProgramName version-name

Show the current version
`

func versionNameCmd(args cli.Args) {
	currentVersion := detectVersion()
	cli.Println(currentVersion.Name)
}

func init() {
	cli.Register("version-name", versionNameCmd, versionNameHelp)
}
