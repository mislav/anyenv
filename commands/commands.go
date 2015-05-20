package commands

import (
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/utils"
	"os"
	"strings"
)

var commandsHelp = `
Usage: $ProgramName commands [--sh|--no-sh]

List all available $ProgramName commands
`

func commandsCmd(args cli.Args) {
	shOnly := args.HasFlag("--sh")
	noSh := args.HasFlag("--no-sh")

	commandNames := findAvailableCommands(args.ProgramName())
	toPrint := utils.NewSet()

	for _, name := range commandNames.Array() {
		isSh := strings.HasPrefix(name, "sh-")
		if (!shOnly || isSh) && !(noSh && isSh) {
			toPrint.Add(strings.TrimPrefix(name, "sh-"))
		}
	}

	for _, name := range toPrint.Sort() {
		cli.Println(name)
	}
}

func findAvailableCommands(programName string) utils.Set {
	commandNames := utils.NewSetFromSlice(cli.CommandNames())
	prefix := programName + "-"

	paths := strings.Split(os.Getenv("PATH"), ":")
	for _, p := range paths {
		if p == "" {
			continue
		}
		path := utils.NewPathname(p)
		for _, cmd := range path.EntriesMatching(prefix + "*") {
			if cmd.IsExecutable() {
				commandNames.Add(strings.TrimPrefix(cmd.Base(), prefix))
			}
		}
	}

	return commandNames
}

func init() {
	cli.Register("commands", commandsCmd, commandsHelp)
}
