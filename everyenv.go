package main

import (
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/commands"
	"os"
	"syscall"
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
		exeName := args.ProgramName() + "-" + cmdName
		exeCmd := commands.FindInPath(exeName)

		argv := []string{exeName}
		argv = append(argv, args.ARGV[2:]...)

		err := syscall.Exec(exeCmd.String(), argv, os.Environ())
		if err != nil {
			cli.Errorf("%s: %s\n", exeName, err)
			cli.Exit(1)
		}
	}
}
