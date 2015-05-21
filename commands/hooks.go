package commands

import (
	"github.com/mislav/anyenv/cli"
	"github.com/mislav/anyenv/config"
	"os"
	"path/filepath"
	"strings"
)

var hooksHelp = `
Usage: $ProgramName hooks <command>

List hook scripts for the specified $ProgramName command.

The hook scripts should be named '<command>/*.bash' and are searched for under
these paths in order:

  $HookEnvName  (colon-separated list)
  $Root/$ProgramName.d
  /usr/local/etc/$ProgramName.d
  /etc/$ProgramName.d
  /usr/lib/$ProgramName/hooks
  $PluginsDir/*/etc/$ProgramName.d
`

func hooksCmd(args cli.Args) {
	commandName := args.Required(0)

	for _, script := range findHookScripts(commandName) {
		cli.Println(script)
	}
}

func findHookScripts(commandName string) []string {
	results := []string{}

	paths := append(strings.Split(os.Getenv(config.HookEnvName), ":"),
		filepath.Join(config.Root, "rbenv.d"),
		"/usr/local/etc/rbenv.d",
		"/etc/rbenv.d",
		"/usr/lib/rbenv/hooks")

	pluginsGlob := config.PluginsDir().Join("*/etc/rbenv.d")
	pluginPaths, err := filepath.Glob(pluginsGlob.String())
	if err != nil {
		panic(err)
	}
	paths = append(paths, pluginPaths...)

	for _, path := range paths {
		if path == "" {
			continue
		}
		hookScripts, err := filepath.Glob(filepath.Join(path, commandName, "*.bash"))
		if err != nil {
			panic(err)
		}
		for _, script := range hookScripts {
			results = append(results, script)
		}
	}

	return results
}

func init() {
	cli.Register("hooks", hooksCmd, hooksHelp)
}
