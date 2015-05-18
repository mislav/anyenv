package commands

import (
	"fmt"
	"github.com/mislav/everyenv/cli"
)

func init() {
	cli.Register("version", func(args []string) {
		fmt.Printf("version: %#v\n", args)
	})
}
