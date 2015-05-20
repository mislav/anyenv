package commands

import (
	"github.com/mislav/everyenv/cli"
	"sort"
	"strings"
)

var commandsHelp = `
Usage: $ProgramName commands [--sh|--no-sh]

List all available $ProgramName commands
`

func commandsCmd(args cli.Args) {
	shOnly := args.HasFlag("--sh")
	noSh := args.HasFlag("--no-sh")

	commandNames := cli.CommandNames()
	sort.Strings(commandNames)

	for _, name := range commandNames {
		isSh := strings.HasPrefix(name, "sh-")
		if (!shOnly || isSh) && !(noSh && isSh) {
			cli.Println(name)
		}
	}
}

func init() {
	cli.Register("commands", commandsCmd, commandsHelp)
}
