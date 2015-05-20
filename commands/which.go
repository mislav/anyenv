package commands

import (
	"github.com/mislav/anyenv/cli"
	"github.com/mislav/anyenv/config"
	"github.com/mislav/anyenv/utils"
	"os"
	"strings"
)

var whichHelp = `
Usage: $ProgramName which <command>

Displays the full path to the executable that $ProgramName will invoke when
running the given command.
`

func whichCmd(args cli.Args) {
	currentVersion := detectVersion()
	exeName := args.Required(0)
	exePath, err := findExecutable(exeName, currentVersion)
	if err != nil {
		cli.Errorf("%s: %s\n", args.ProgramName(), err)
		cli.Exit(1)
	}

	if exePath.IsBlank() {
		cli.Errorf("%s: %s: command not found\n", args.ProgramName(), exeName)
		versions := whence(exeName)
		if len(versions) > 0 {
			cli.Errorf("\nThe `%s' command exists in these versions:\n  %s\n", exeName,
				strings.Join(versions, "\n  "))
		}
		cli.Exit(127)
	} else {
		cli.Println(exePath)
	}
}

func findExecutable(exeName string, currentVersion SelectedVersion) (filename utils.Pathname, err error) {
	if currentVersion.IsSystem() {
		filename = findOnSystem(exeName)
	} else {
		versionDir := config.VersionDir(currentVersion.Name)
		if !versionDir.Exists() {
			err = VersionNotFound{currentVersion.Name}
			return
		}
		filename = versionDir.Join("bin", exeName)
		if !filename.IsExecutable() {
			filename = utils.NewPathname("")
		}
	}

	return
}

func findOnSystem(exeName string) utils.Pathname {
	shimsDir := config.ShimsDir()
	dirs := strings.Split(os.Getenv("PATH"), ":")

	var dir utils.Pathname
	var filename utils.Pathname

	for _, p := range dirs {
		dir = utils.NewPathname(p)
		if dir.IsBlank() || dir.Equal(shimsDir) {
			continue
		}
		filename = dir.Join(exeName)
		if filename.IsExecutable() {
			return filename.Abs()
		}
	}

	return utils.NewPathname("")
}

func init() {
	cli.Register("which", whichCmd, whichHelp)
}
