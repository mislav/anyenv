package cli

import (
	"github.com/mislav/anyenv/config"
	"os"
	"strings"
)

var commands = make(map[string]func(Args))
var helpText = make(map[string]string)

func Lookup(cmdName string) func(Args) {
	return commands[cmdName]
}

func HelpText(cmdName string, programName string) string {
	text := helpText[cmdName]
	if text == "" && !strings.HasPrefix(cmdName, "sh-") {
		text = helpText["sh-"+cmdName]
	}

	text = strings.Trim(text, " \n")

	return os.Expand(text, func(name string) string {
		switch name {
		case "ProgramName":
			return programName
		case "VersionsDir":
			return strings.Replace(config.VersionsDir().String(),
				config.Root,
				"$"+config.RootEnvName,
				1)
		case "VersionFilename":
			return config.VersionFilename
		case "VersionEnvName":
			return config.VersionEnvName
		}
		return "$" + name
	})
}

func Register(cmdName string, fn func(Args), help string) {
	commands[cmdName] = fn
	helpText[cmdName] = help
}

func CommandNames() []string {
	names := make([]string, len(commands))
	i := 0
	for name, _ := range commands {
		names[i] = name
		i += 1
	}
	return names
}
