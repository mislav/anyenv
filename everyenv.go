package main

import (
	"github.com/mislav/everyenv/cli"
	_ "github.com/mislav/everyenv/commands"
	"log"
	"os"
)

func main() {
	cmdName := os.Args[1]
	cmd := cli.Lookup(cmdName)

	if cmd != nil {
		cmd(os.Args[2:])
	} else {
		log.Fatalf("command not found: `%s`\n", cmdName)
	}
}
