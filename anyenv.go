package main

import (
	"github.com/mislav/anyenv/cli"
	_ "github.com/mislav/anyenv/commands"
	"github.com/mislav/anyenv/utils"
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
		results := utils.SearchInPath(exeName)
		if len(results) == 0 {
			cli.Errorf("%s: command not found\n", exeName)
			cli.Exit(1)
		}
		exeCmd := results[0]

		argv := []string{exeName}
		argv = append(argv, args.ARGV[2:]...)

		err := syscall.Exec(exeCmd.String(), argv, os.Environ())
		if err != nil {
			cli.Errorf("%s: %s\n", exeName, err)
			cli.Exit(1)
		}
	}
}
