package commands

import (
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
	"github.com/mislav/everyenv/utils"
)

var whenceHelp = `
Usage: $ProgramName whence [--path] <command>

List all versions that contain the given executable
`

func whenceCmd(args cli.Args) {
	exeName := args.Required(0)
	versions := whence(exeName)

	for _, version := range versions {
		cli.Println(version)
	}
}

func whence(exeName string) []string {
	var exeFile utils.Pathname
	results := []string{}

	versionsDir := config.VersionsDir()
	versionPaths := versionsDir.Entries()

	for _, versionPath := range versionPaths {
		exeFile = versionPath.Join("bin", exeName)
		if exeFile.IsExecutable() {
			results = append(results, versionPath.Base())
		}
	}

	return results
}

func init() {
	cli.Register("whence", whenceCmd, whenceHelp)
}
