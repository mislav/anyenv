package commands

import (
	"fmt"
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
	"log"
	"os"
	"path"
	"strings"
)

func whichCmd(args cli.Args) {
	var exePath string
	exeName := args.List[0]
	currentVersion := detectVersion()

	if currentVersion.IsSystem() {
		shimsDir := path.Join(config.Root, "shims")
		dirs := strings.Split(os.Getenv("PATH"), ":")
		for _, dir := range dirs {
			if dir == "" || dir == shimsDir {
				continue
			}
			filename := path.Join(dir, exeName)
			if fileExecutable(filename) {
				exePath = filename
				break
			}
		}
	} else {
		filename := path.Join(config.Root, "versions", currentVersion.Name, "bin", exeName)
		if fileExecutable(filename) {
			exePath = filename
		}
	}

	if exePath == "" {
		log.Fatalf("command not found: `%s`\n", exeName)
	} else {
		fmt.Println(exePath)
	}
}

func fileExecutable(filename string) bool {
	fileInfo, err := os.Stat(filename)
	return err == nil && (fileInfo.Mode()&0111) != 0
}

func init() {
	cli.Register("which", whichCmd)
}
