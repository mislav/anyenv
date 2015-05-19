package commands

import (
	"fmt"
	"github.com/mislav/everyenv/cli"
)

func init() {
	cli.Register("rehash", func(args cli.Args) {
		fmt.Printf("rehash: %#v\n", args)
	})
}
