package main

import (
	"github.com/mislav/anyenv/cli"
	_ "github.com/mislav/anyenv/commands"
	"github.com/mislav/anyenv/config"
	"github.com/mislav/anyenv/utils"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {
	args := cli.Args{os.Args}
	cmdName := args.CommandName()
	if cmdName == "" {
		cmdName = "help"
	}

	preparePath()

	cmd := cli.Lookup(cmdName)
	if cmd != nil {
		cmd(args)
	} else {
		exeName := args.ProgramName() + "-" + cmdName
		results := utils.SearchInPath(exeName)
		if len(results) == 0 {
			cli.Errorf("%s: command not found\n", exeName)
			cli.Exit(1)
		}
		exeCmd := results[0]

		argv := []string{exeName}
		argv = append(argv, args.ARGV[2:]...)

		err := syscall.Exec(exeCmd.String(), argv, os.Environ())
		if err != nil {
			cli.Errorf("%s: %s\n", exeName, err)
			cli.Exit(1)
		}
	}
}

func preparePath() {
	pluginPaths, err := filepath.Glob(config.PluginsDir().Join("*/bin").String())
	if err != nil {
		panic(err)
	}

	if len(pluginPaths) > 0 {
		os.Setenv("PATH", strings.Join(append(pluginPaths, os.Getenv("PATH")), ":"))
	}
}
