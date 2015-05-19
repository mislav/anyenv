package main

import (
	"github.com/mislav/everyenv/cli"
	_ "github.com/mislav/everyenv/commands"
	"os"
)

func main() {
	cmdName := os.Args[1]
	cmd := cli.Lookup(cmdName)

	if cmd != nil {
		cmd(cli.Args{os.Args[2:]})
	} else {
		cli.Errorf("%s: no such command\n", cmdName)
		cli.Exit(1)
	}
}
