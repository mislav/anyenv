package commands

import (
	"github.com/mislav/everyenv/cli"
	"sort"
)

func commandsCmd(args cli.Args) {
	commandNames := cli.CommandNames()
	sort.Strings(commandNames)

	for _, name := range commandNames {
		cli.Println(name)
	}
}

func init() {
	cli.Register("commands", commandsCmd)
}
