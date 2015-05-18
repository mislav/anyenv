package main

import (
	"fmt"
	"github.com/mislav/everyenv/cli"
	_ "github.com/mislav/everyenv/commands"
	"os"
)

func main() {
	cmdName := os.Args[1]
	cmd := cli.Lookup(cmdName)

	if cmd != nil {
		cmd(os.Args[2:])
	} else {
		fmt.Printf("command not found: `%s`\n", cmdName)
		os.Exit(1)
	}
}
