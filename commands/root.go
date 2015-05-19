package commands

import (
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
)

var rootHelp = `
Usage: $ProgramName root

Display the root directory where versions and shims are kept
`

func rootCmd(args cli.Args) {
	cli.Println(config.Root)
}

func init() {
	cli.Register("root", rootCmd, rootHelp)
}
