package commands

import (
	"github.com/mislav/anyenv/cli"
)

var helpHelp = `
Usage: $ProgramName <command> [<args>]

Some useful $ProgramName commands are:
   commands    List all available commands
   local       Set or show the local application-specific version
   global      Set or show the global version
   shell       Set or show the shell-specific version
   rehash      Rehash $ProgramName shims (run this after installing executables)
   version     Show the current version and its origin
   versions    List all versions available to $ProgramName
   which       Display the full path to an executable
   whence      List all versions that contain an executable

See '$ProgramName help <command>' for information on a specific command.
`

func helpCmd(args cli.Args) {
	commandName := args.At(0)
	if commandName == "" {
		commandName = "help"
	}

	text := cli.HelpText(commandName, args.ProgramName())

	if text == "" {
		cli.Errorf("%s: %s: no such command\n", args.ProgramName(), commandName)
		cli.Exit(1)
	} else {
		cli.Println(text)
	}
}

func init() {
	cli.Register("help", helpCmd, helpHelp)
}
