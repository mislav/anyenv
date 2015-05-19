package commands

import (
	"github.com/mislav/everyenv/cli"
)

var helpHelp = `
Usage: $ProgramName help [<command>]
`

func helpCmd(args cli.Args) {
	commandName := args.At(0)
	if commandName == "" {
		commandName = "help"
	}
	text := cli.HelpText(commandName, args.ProgramName())
	cli.Println(text)
}

func init() {
	cli.Register("help", helpCmd, helpHelp)
}
