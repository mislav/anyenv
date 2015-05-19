package commands

import (
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
	"github.com/mislav/everyenv/utils"
	"os"
	"strings"
)

func whichCmd(args cli.Args) {
	currentVersion := detectVersion()
	exeName := args.List[0]
	exePath := findExecutable(exeName, currentVersion)

	if exePath.IsBlank() {
		cli.Errorf("%s: command not found\n", exeName)
		cli.Exit(127)
	} else {
		cli.Println(exePath)
	}
}

func findExecutable(exeName string, currentVersion SelectedVersion) utils.Pathname {
	if currentVersion.IsSystem() {
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
				return filename
			}
		}
	} else {
		versionDir := config.VersionDir(currentVersion.Name)
		if !versionDir.Exists() {
			cli.Errorf("version `%s' is not installed\n", currentVersion.Name)
			cli.Exit(1)
		}
		filename := versionDir.Join("bin", exeName)
		if filename.IsExecutable() {
			return filename
		}
	}

	return utils.NewPathname("")
}

func init() {
	cli.Register("which", whichCmd)
}
