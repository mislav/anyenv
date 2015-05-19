package main

import (
	"github.com/mislav/everyenv/cli"
	_ "github.com/mislav/everyenv/commands"
	"os"
)

func main() {
	args := cli.Args{os.Args}
	cmdName := args.CommandName()
	if cmdName == "" {
		cmdName = "help"
	}

	cmd := cli.Lookup(cmdName)
	if cmd != nil {
		cmd(args)
	} else {
		cli.Errorf("%s %s: no such command\n", args.ProgramName(), cmdName)
		cli.Exit(1)
	}
}
