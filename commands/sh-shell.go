package commands

import (
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
)

var shellHelp = `
Usage: $ProgramName shell <version>
       $ProgramName shell --unset

Sets a shell-specific Ruby version by setting the '$VersionEnvName'
environment variable in your shell. This version overrides local
application-specific versions and the global version.

<version> should be a string matching a version known to $ProgramName.
The special version string "system" will use the default system version.
See '$ProgramName versions' for a list of available versions.
`

func shellCmd(args cli.Args) {
	assign := args.Word(0)
	unset := args.HasFlag("--unset")

	shell := cli.DetectShell("")

	if unset {
		if shell == "fish" {
			cli.Printf("set -e %s\n", config.VersionEnvName)
		} else {
			cli.Printf("unset %s\n", config.VersionEnvName)
		}
	} else if assign == "" {
		if version := config.VersionEnv(); version != "" {
			cli.Printf("echo \"$%s\"\n", config.VersionEnvName)
		} else {
			cli.Errorf("%s: no shell-specific version configured\n", args.ProgramName())
			cli.Exit(1)
		}
	} else {
		if assign != "system" {
			versionDir := config.VersionDir(assign)
			if !versionDir.Exists() {
				err := VersionNotFound{assign}
				cli.Errorf("%s: %s\n", args.ProgramName(), err)
				cli.Exit(1)
			}
		}

		if shell == "fish" {
			cli.Printf("setenv %s \"%s\"\n", config.VersionEnvName, assign)
		} else {
			cli.Printf("export %s=\"%s\"\n", config.VersionEnvName, assign)
		}
	}
}

func init() {
	cli.Register("sh-shell", shellCmd, shellHelp)
}
