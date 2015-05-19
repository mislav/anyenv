package commands

import (
	"fmt"
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
	"os"
	"path"
)

func versionsCmd(args []string) {
	var versions []string
	bare := cli.HasFlag("--bare", args)

	versionsDir := path.Join(config.Root, "versions")
	dir, err := os.Open(versionsDir)
	if err == nil {
		versions, _ = dir.Readdirnames(0)
	} else {
		versions = []string{}
	}

	if bare {
		for _, version := range versions {
			fmt.Printf("%s\n", version)
		}
	} else {
		currentVersion := detectVersion()

		for _, version := range versions {
			if version == currentVersion.Name {
				fmt.Printf("* %s (set by %s)\n", version, currentVersion.Origin)
			} else {
				fmt.Printf("  %s\n", version)
			}
		}
	}
}

func init() {
	cli.Register("versions", versionsCmd)
}
