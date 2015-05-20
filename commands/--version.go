package commands

import (
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
)

var _versionHelp = `
Usage: $ProgramName --version

Display the build version of $ProgramName itself.
`

func _versionCmd(args cli.Args) {
	cli.Printf("%s %s\n", args.ProgramName(), config.BuildVersion)
}

func init() {
	cli.Register("--version", _versionCmd, _versionHelp)
}
