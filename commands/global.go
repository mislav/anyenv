package commands

import (
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
	"os"
)

var globalHelp = `
Usage: $ProgramName global
       $ProgramName global <version>

Sets the global version. You can override the global version at
any time by setting a directory-specific version with '$ProgramName local'
or by setting the '$VersionEnvName' environment variable.

<version> should be a string matching a version known to $ProgramName.
The special version string "system" will use the default system version.
See '$ProgramName versions' for a list of available versions.
`

func globalCmd(args cli.Args) {
	versionFile := config.GlobalVersionFile()
	assign := args.At(0)

	if assign == "" {
		if versionFile.Exists() {
			version, _ := readVersionFile(versionFile)
			cli.Println(version)
		} else {
			cli.Println("system")
		}
	} else {
		if assign != "system" {
			versionDir := config.VersionDir(assign)
			if !versionDir.Exists() {
				cli.Errorf("version `%s` not installed\n", assign)
				cli.Exit(1)
			}
		}

		err := os.MkdirAll(versionFile.Dir().String(), 0755)
		if err != nil {
			panic(err)
		}

		err = writeVersionFile(versionFile, assign)
		if err != nil {
			panic(err)
		}
	}
}

func init() {
	cli.Register("global", globalCmd, globalHelp)
}
