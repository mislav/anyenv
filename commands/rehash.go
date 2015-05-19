package commands

import (
	"github.com/mislav/everyenv/cli"
)

var rehashHelp = `
Usage: $program_name rehash

Rebuild $program_name shims. Run this after installing executables.
`

func rehashCmd(args cli.Args) {
	cli.Printf("rehash: %#v\n", args)
}

func init() {
	cli.Register("rehash", rehashCmd, rehashHelp)
}
