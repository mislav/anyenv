package commands

import (
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
	"os"
	"strings"
	"syscall"
)

func execCmd(args cli.Args) {
	currentVersion := detectVersion()
	exeName := args.List[0]
	exePath := findExecutable(exeName, currentVersion)
	env := os.Environ()

	if !currentVersion.IsSystem() {
		for i, value := range env {
			if strings.Index(value, "PATH=") == 0 {
				pair := strings.SplitN(value, "=", 2)
				versionBindir := config.VersionDir(currentVersion.Name).Join("bin")
				env[i] = "PATH=" + versionBindir.String() + ":" + pair[1]
			}
		}
	}

	argv := []string{exeName}
	argv = append(argv, args.List[1:]...)
	syscall.Exec(exePath.String(), argv, env)
}

func init() {
	cli.Register("exec", execCmd)
}
