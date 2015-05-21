package commands

import (
	"github.com/mislav/anyenv/cli"
	"github.com/mislav/anyenv/config"
	"github.com/mislav/anyenv/utils"
	"os/exec"
	"strings"
	"syscall"
)

var execHelp = `
Usage: $program_name exec <command> [arg1 arg2...]

Runs an executable by first preparing PATH so that the selected version's
'bin' directory is directly in the front.
`

func execCmd(args cli.Args) {
	currentVersion := detectVersion()
	exeName := args.Required(0)
	exePath, err := findExecutable(exeName, currentVersion)
	if err != nil {
		cli.Errorf("%s: %s\n", args.ProgramName(), err)
		cli.Exit(1)
	}

	env := utils.EnvFromEnviron()

	hooks := findHookScripts("exec")
	if len(hooks) > 0 {
		hookArgs := append([]string{"-c", `
			for file; do source "$file"; done
			export
		`, "--"}, hooks...)
		hookCmd := exec.Command("bash", hookArgs...)
		hookOut, err := hookCmd.Output()
		if err != nil {
			panic(err)
		}
		lines := strings.Split(string(hookOut), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "declare -x ") {
				pair := strings.SplitN(strings.TrimPrefix(line, "declare -x "), "=", 2)
				if len(pair) > 1 {
					env.Set(pair[0], strings.Trim(pair[1], "\""))
				}
			}
		}
	}

	if !currentVersion.IsSystem() {
		versionBindir := config.VersionDir(currentVersion.Name).Join("bin")
		env.Set("PATH", versionBindir.String()+":"+env.Get("PATH"))
	}

	argv := []string{exeName}
	argv = append(argv, args.ARGV[3:]...)

	err = syscall.Exec(exePath.String(), argv, env.Environ())
	if err != nil {
		cli.Errorf("%s: %s\n", args.ProgramName(), err)
		cli.Exit(1)
	}
}

func init() {
	cli.Register("exec", execCmd, execHelp)
}
