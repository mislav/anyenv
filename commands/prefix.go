package commands

import (
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
)

var prefixHelp = `
Usage: $ProgramName prefix [<version>]

Displays the directory where a specific version is installed. If no version is
given, display the location of the currently selected version.
`

func prefixCmd(args cli.Args) {
	version := args.At(0)
	if version == "" {
		currentVersion := detectVersion()
		version = currentVersion.Name
	}

	if version == "system" {
		exePath := findOnSystem(config.MainExecutable)
		if exePath.IsBlank() {
			cli.Errorf("%s: system version not found in PATH\n", args.ProgramName())
			cli.Exit(1)
		} else {
			cli.Println(exePath.Dir().Dir())
		}
	} else {
		versionDir := config.VersionDir(version)
		if versionDir.Exists() {
			cli.Println(versionDir)
		} else {
			err := VersionNotFound{version}
			cli.Errorf("%s: %s\n", args.ProgramName(), err)
			cli.Exit(1)
		}
	}
}

func init() {
	cli.Register("prefix", prefixCmd, prefixHelp)
}
