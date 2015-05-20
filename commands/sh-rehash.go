package commands

import (
	"github.com/mislav/everyenv/cli"
)

func shellRehashCmd(args cli.Args) {
	shell := cli.DetectShell("")

	cli.Printf("command %s rehash\n", args.ProgramName())
	if shell != "fish" {
		cli.Println("hash -r 2>/dev/null || true")
	}
}

func init() {
	cli.Register("sh-rehash", shellRehashCmd, "")
}
