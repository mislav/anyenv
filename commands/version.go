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

type SelectedVersion struct {
	Name   string
	Origin string
}

func versionCmd(args []string) {
	currentVersion := detectVersion()
	fmt.Printf("%s (set by %s)\n", currentVersion.Name, currentVersion.Origin)
}

func detectVersion() SelectedVersion {
	version := config.VersionEnv()
	origin := config.VersionEnvName

	if version == "" {
		pwd, _ := os.Getwd()
		origin = findVersionFile(pwd)
		version, _ = readVersionFile(origin)
	}

	if version == "" {
		var err error
		origin = path.Join(config.Root, "version")
		version, err = readVersionFile(origin)
		if err != nil {
			log.Fatal(err)
		}
	}

	return SelectedVersion{version, origin}
}

func fileExists(filename string) bool {
	fileInfo, err := os.Stat(filename)
	return err == nil && !fileInfo.IsDir()
}

func findVersionFile(dir string) (filename string) {
	var parentDir string
	for {
		filename = path.Join(dir, config.VersionFilename)
		if fileExists(filename) {
			return
		}
		parentDir = path.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}
	return ""
}

func readVersionFile(filename string) (value string, err error) {
	if filename == "" {
		return
	}

	file, err := os.Open(filename)
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
