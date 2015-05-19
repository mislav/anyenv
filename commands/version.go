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

type SelectedVersion struct {
	Name   string
	Origin string
}

func (ver SelectedVersion) IsSystem() bool {
	return "system" == ver.Name
}

func versionCmd(args cli.Args) {
	currentVersion := detectVersion()
	fmt.Printf("%s (set by %s)\n", currentVersion.Name, currentVersion.Origin)
}

func detectVersion() SelectedVersion {
	version := config.VersionEnv()
	origin := config.VersionEnvName

	if version == "" {
		pwd := utils.Getwd()
		versionFile := findVersionFile(pwd)
		if !versionFile.IsBlank() {
			version, _ = readVersionFile(versionFile)
			origin = versionFile.String()
		}
	}

	if version == "" {
		var err error
		globalVersionFile := utils.NewPathname(config.Root, "version")
		origin = globalVersionFile.String()
		version, err = readVersionFile(globalVersionFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	return SelectedVersion{version, origin}
}

func findVersionFile(dir utils.Pathname) (versionFile utils.Pathname) {
	for {
		versionFile = dir.Join(config.VersionFilename)
		if versionFile.Exists() {
			return
		}
		if dir.IsRoot() {
			break
		}
		dir = dir.Dir()
	}
	return utils.NewPathname("")
}

func readVersionFile(filename utils.Pathname) (value string, err error) {
	file, err := os.Open(filename.String())
	if err != nil {
		return
	}

	defer file.Close()
	data := make([]byte, 1024)
	count, err := file.Read(data)
	if err != nil {
		return
	}

	value = strings.TrimRight(string(data[:count]), " \r\n")
	return
}

func init() {
	cli.Register("version", versionCmd)
}
