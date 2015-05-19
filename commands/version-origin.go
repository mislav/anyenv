package commands

import (
	"github.com/mislav/everyenv/cli"
)

var versionOriginHelp = `
Usage: $ProgramName version-origin

Explain how the current version is set
`

func versionOriginCmd(args cli.Args) {
	currentVersion := detectVersion()
	cli.Println(currentVersion.Origin)
}

func init() {
	cli.Register("version-origin", versionOriginCmd, versionOriginHelp)
}
