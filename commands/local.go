package commands

import (
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
	"github.com/mislav/everyenv/utils"
	"os"
)

var localHelp = `
Usage: $ProgramName local
       $ProgramName local <version>
       $ProgramName local --unset

Sets the local application-specific version by writing the version name to a
file named '$VersionFilename'.

When a command is run, $ProgramName will look for a '$VersionFilename' file in
the current directory and each parent directory. If no such file is found in the
tree, $ProgramName will use the global version specified with '$ProgramName global'.
A version specified with the '$VersionEnvName' environment variable takes
precedence over local and global versions.

<version> should be a string matching a version known to $ProgramName.
The special version string "system" will use the default system version.
See '$ProgramName versions' for a list of available versions.
`

func localCmd(args cli.Args) {
	pwd := utils.Getwd()
	versionFile := pwd.Join(config.VersionFilename)
	unset := args.HasFlag("--unset")
	assign := args.At(0)

	if unset {
		if versionFile.Exists() {
			err := os.Remove(versionFile.String())
			if err != nil {
				panic(err)
			}
		}
	} else if assign == "" {
		if versionFile.Exists() {
			version, _ := readVersionFile(versionFile)
			cli.Println(version)
		} else {
			cli.Errorf("%s: no local version configured for this directory\n", args.ProgramName())
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
		err := writeVersionFile(versionFile, assign)
		if err != nil {
			panic(err)
		}
	}
}

func writeVersionFile(filename utils.Pathname, value string) (err error) {
	file, err := os.Create(filename.String())
	if err != nil {
		return
	}

	defer file.Close()
	_, err = file.WriteString(value)
	return
}

func init() {
	cli.Register("local", localCmd, localHelp)
}
