package commands

import (
	"fmt"
	"github.com/mislav/everyenv/cli"
	"sort"
)

func commandsCmd(args cli.Args) {
	commandNames := cli.CommandNames()
	sort.Strings(commandNames)

	for _, name := range commandNames {
		fmt.Printf("%s\n", name)
	}
}

func init() {
	cli.Register("commands", commandsCmd)
}
