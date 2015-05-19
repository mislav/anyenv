package commands

import (
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
)

func versionsCmd(args cli.Args) {
	bare := args.HasFlag("--bare")

	versionsDir := config.VersionsDir()
	versionPaths := versionsDir.Entries()

	if bare {
		for _, versionPath := range versionPaths {
			cli.Printf("%s\n", versionPath.Base())
		}
	} else {
		currentVersion := detectVersion()

		for _, versionPath := range versionPaths {
			if versionPath.Base() == currentVersion.Name {
				cli.Printf("* %s (set by %s)\n", currentVersion.Name, currentVersion.Origin)
			} else {
				cli.Printf("  %s\n", versionPath.Base())
			}
		}
	}
}

func init() {
	cli.Register("versions", versionsCmd)
}
