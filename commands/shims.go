package commands

import (
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
)

var shimsHelp = `
Usage: $program_name shims [--short]

List existing $program_name shims
`

func shimsCmd(args cli.Args) {
	shimsDir := config.ShimsDir()
	short := args.HasFlag("--short")

	for _, shim := range shimsDir.Entries() {
		if short {
			cli.Println(shim.Base())
		} else {
			cli.Println(shim)
		}
	}
}

func init() {
	cli.Register("shims", shimsCmd, shimsHelp)
}
