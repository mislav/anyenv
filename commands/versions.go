package commands

import (
	"fmt"
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
	"github.com/mislav/everyenv/utils"
)

func versionsCmd(args cli.Args) {
	bare := args.HasFlag("--bare")

	versionsPath := utils.NewPathname(config.Root, "versions")
	versionPaths := versionsPath.Entries()

	if bare {
		for _, versionPath := range versionPaths {
			fmt.Printf("%s\n", versionPath.Base())
		}
	} else {
		currentVersion := detectVersion()

		for _, versionPath := range versionPaths {
			if versionPath.Base() == currentVersion.Name {
				fmt.Printf("* %s (set by %s)\n", currentVersion.Name, currentVersion.Origin)
			} else {
				fmt.Printf("  %s\n", versionPath.Base())
			}
		}
	}
}

func init() {
	cli.Register("versions", versionsCmd)
}
