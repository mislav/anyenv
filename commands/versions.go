package commands

import (
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
)

var versionsHelp = `
Usage: $ProgramName versions [--bare]

Lists all versions found in '$VersionsDir'.
`

func versionsCmd(args cli.Args) {
	bare := args.HasFlag("--bare")

	versionsDir := config.VersionsDir()
	versions := versionsDir.BareEntries()

	systemExecutable := findInPath(config.MainExecutable)
	if !systemExecutable.IsBlank() {
		versions = append([]string{"system"}, versions...)
	}

	if bare {
		for _, version := range versions {
			cli.Printf("%s\n", version)
		}
	} else {
		currentVersion := detectVersion()

		for _, version := range versions {
			if version == currentVersion.Name {
				cli.Printf("* %s (set by %s)\n", currentVersion.Name, currentVersion.Origin)
			} else {
				cli.Printf("  %s\n", version)
			}
		}
	}
}

func init() {
	cli.Register("versions", versionsCmd, versionsHelp)
}
