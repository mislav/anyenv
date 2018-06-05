package commands

import (
	"github.com/mislav/anyenv/cli"
	"github.com/mislav/anyenv/config"
)

var versionFlagHelp = `
Usage: $ProgramName --version

Display the build version of $ProgramName itself.
`

func versionFlagCmd(args cli.Args) {
	cli.Printf("%s %s\n", args.ProgramName(), config.BuildVersion)
}

func init() {
	cli.Register("--version", versionFlagCmd, versionFlagHelp)
}
