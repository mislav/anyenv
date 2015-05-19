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

	versionDir := config.VersionDir(version)
	if versionDir.Exists() {
		cli.Println(versionDir)
	} else {
		cli.Errorf("version `%s' not installed\n", version)
		cli.Exit(1)
	}
}

func init() {
	cli.Register("prefix", prefixCmd, prefixHelp)
}
