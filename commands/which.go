package commands

import (
	"fmt"
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
	"github.com/mislav/everyenv/utils"
	"log"
	"os"
	"strings"
)

func whichCmd(args cli.Args) {
	exePath := utils.NewPathname("")
	exeName := args.List[0]
	currentVersion := detectVersion()

	if currentVersion.IsSystem() {
		shimsDir := utils.NewPathname(config.Root, "shims")
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
		filename := utils.NewPathname(config.Root, "versions", currentVersion.Name, "bin", exeName)
		if filename.IsExecutable() {
			exePath = filename
		}
	}

	if exePath.IsBlank() {
		log.Fatalf("command not found: `%s`\n", exeName)
	} else {
		fmt.Println(exePath)
	}
}

func init() {
	cli.Register("which", whichCmd)
}
