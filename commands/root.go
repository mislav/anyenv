package commands

import (
	"github.com/mislav/anyenv/cli"
	"github.com/mislav/anyenv/config"
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
