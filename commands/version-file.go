package commands

import (
	"github.com/mislav/anyenv/cli"
)

var versionFileHelp = `
Usage: $ProgramName version-file
`

func versionFileCmd(args cli.Args) {
	versionFile := detectVersionFile()
	cli.Println(versionFile)
}

func init() {
	cli.Register("version-file", versionFileCmd, versionFileHelp)
}
