package commands

import (
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
	"github.com/mislav/everyenv/utils"
	"os"
	"strings"
)

func whichCmd(args cli.Args) {
	exePath := utils.NewPathname("")
	exeName := args.List[0]
	currentVersion := detectVersion()

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
				exePath = filename
				break
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
			exePath = filename
		}
	}

	if exePath.IsBlank() {
		cli.Errorf("%s: command not found\n", exeName)
		cli.Exit(127)
	} else {
		cli.Println(exePath)
	}
}

func init() {
	cli.Register("which", whichCmd)
}
