package commands

import (
	"github.com/mislav/anyenv/cli"
	"github.com/mislav/anyenv/config"
	"github.com/mislav/anyenv/utils"
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

type VersionNotFound struct {
	Version string
}

func (err VersionNotFound) Error() string {
	return "version `" + err.Version + "' is not installed"
}

var versionHelp = `
Usage: $ProgramName version

Shows the currently selected version and how it was
selected. To obtain only the version string, use
'$ProgramName version-name'.
`

func versionCmd(args cli.Args) {
	currentVersion := detectVersion()
	cli.Printf("%s (set by %s)\n", currentVersion.Name, currentVersion.Origin)
}

func detectVersion() SelectedVersion {
	version := config.VersionEnv()
	origin := config.VersionEnvName + " environment variable"

	if version == "" {
		versionFile := detectVersionFile()
		origin = versionFile.String()
		var err error
		version, err = readVersionFile(versionFile)
		if err != nil {
			version = "system"
		}
	}

	versionPrefix := config.MainExecutable + "-"
	if strings.HasPrefix(version, versionPrefix) {
		prefixedDir := config.VersionDir(version)
		unprefixedDir := config.VersionDir(strings.TrimPrefix(version, versionPrefix))
		if !prefixedDir.Exists() && unprefixedDir.Exists() {
			version = unprefixedDir.Base()
		}
	}

	return SelectedVersion{version, origin}
}

func detectVersionFile() utils.Pathname {
	dirEnv := utils.NewPathname(config.DirEnv())
	pwd := utils.Getwd()

	if !dirEnv.IsBlank() {
		versionFile := findVersionFile(dirEnv)
		if !versionFile.IsBlank() {
			return versionFile
		}
	}

	if !pwd.Equal(dirEnv) {
		versionFile := findVersionFile(pwd)
		if !versionFile.IsBlank() {
			return versionFile
		}
	}

	return config.GlobalVersionFile()
}

func findVersionFile(dir utils.Pathname) (versionFile utils.Pathname) {
	for {
		versionFile = dir.Join(config.VersionFilename)
		if versionFile.IsFile() {
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
	cli.Register("version", versionCmd, versionHelp)
}
