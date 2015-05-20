package commands

import (
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
	"github.com/mislav/everyenv/utils"
	"os"
	"path/filepath"
	"strings"
)

var rehashHelp = `
Usage: $ProgramName rehash

Rebuild $ProgramName shims. Run this after installing executables.
`

var shimTemplate = `
#!/usr/bin/env bash
set -e

program="${0##*/}"
if [ "$program" = "$MainExecutable" ]; then
  for arg; do
    case "$arg" in
    -e* | -- ) break ;;
    */* )
      if [ -f "$arg" ]; then
        export $DirEnvName="${arg%/*}"
        break
      fi
      ;;
    esac
  done
fi

export $RootEnvName="$Root"
exec "$FullProgramName" exec "$program" "$@"
`

func rehashCmd(args cli.Args) {
	shimsDir := config.ShimsDir()
	shimFile := shimsDir.Join("." + args.ProgramName() + "-shim")

	shimScript := os.Expand(strings.Trim(shimTemplate, " \n"), func(name string) string {
		switch name {
		case "FullProgramName":
			return args.FullProgramName()
		case "MainExecutable":
			return config.MainExecutable
		case "DirEnvName":
			return config.DirEnvName
		case "RootEnvName":
			return config.RootEnvName
		case "Root":
			return config.Root
		}
		return "${" + name + "}"
	})

	err := os.MkdirAll(shimsDir.String(), 0755)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(shimFile.String(), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0755)
	if err != nil {
		cli.Errorf("%s: cannot rehash: %s\n", args.ProgramName(), err)
		cli.Exit(1)
	}

	_, err = file.WriteString(shimScript)
	if err != nil {
		panic(err)
	}
	file.Close()

	versionsDir := config.VersionsDir()
	matches, err := filepath.Glob(versionsDir.String() + "/*/bin/*")
	if err != nil {
		panic(err)
	}
	executableMap := make(map[string]bool)

	for _, match := range matches {
		foundPath := utils.NewPathname(match)
		if foundPath.IsExecutable() {
			executableMap[foundPath.Base()] = true
		}
	}

	for entry, _ := range executableMap {
		newShim := shimsDir.Join(entry)
		if newShim.Exists() {
			os.Remove(newShim.String())
		}
		os.Link(shimFile.String(), newShim.String())
	}

	for _, existingShim := range shimsDir.Entries() {
		if executableMap[existingShim.Base()] != true {
			os.Remove(existingShim.String())
		}
	}
}

func init() {
	cli.Register("rehash", rehashCmd, rehashHelp)
}
