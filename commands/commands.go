package commands

import (
	"github.com/mislav/everyenv/cli"
	"sort"
)

var commandsHelp = `
Usage: $ProgramName commands

List all available $ProgramName commands
`

func commandsCmd(args cli.Args) {
	commandNames := cli.CommandNames()
	sort.Strings(commandNames)

	for _, name := range commandNames {
		cli.Println(name)
	}
}

func init() {
	cli.Register("commands", commandsCmd, commandsHelp)
}
