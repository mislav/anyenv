package commands

import (
	"github.com/mislav/anyenv/cli"
	"github.com/mislav/anyenv/config"
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
