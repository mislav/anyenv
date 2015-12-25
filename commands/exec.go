package commands

import (
	"github.com/mislav/anyenv/cli"
	"github.com/mislav/anyenv/utils"
	"os/exec"
	"strconv"
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

	binPath := exePath.Dir()
	env := utils.EnvFromEnviron()
	argv := args.ARGV[3:]

	hooks := findHookScripts("exec")
	if len(hooks) > 0 {
		hookEnv := env.Clone()
		hookEnv.Set("RBENV_VERSION", currentVersion.Name)
		hookEnv.Set("RBENV_COMMAND", exeName)
		hookEnv.Set("RBENV_COMMAND_PATH", exePath.String())
		hookEnv.Set("RBENV_BIN_PATH", binPath.String())

		hookArgs := []string{"-e", "-c", `
			scripts=$1; shift 1
			while [ $((scripts--)) -gt 0 ]; do source "$1"; shift 1; done
			echo anyenv -x "RBENV_COMMAND=${RBENV_COMMAND}"
			echo anyenv -x "RBENV_COMMAND_PATH=${RBENV_COMMAND_PATH}"
			echo anyenv -x "RBENV_BIN_PATH=${RBENV_BIN_PATH}"
			export
			printf "%s\n" "---"
			for arg; do printf "%s\0" "$arg"; done
		`, "--", strconv.Itoa(len(hooks))}
		hookArgs = append(hookArgs, hooks...)
		hookArgs = append(hookArgs, argv...)

		hookCmd := exec.Command("bash", hookArgs...)
		hookCmd.Env = hookEnv.Environ()
		hookCmd.Stderr = cli.Stderr

		hookOut, err := hookCmd.Output()
		if err != nil {
			cli.Errorf("%s: %s in `exec' hook\n", args.ProgramName(), err)
			cli.Exit(1)
		}

		lines := strings.Split(string(hookOut), "\n")
		positionalLines := []string{}
		readingPositional := false

		for _, line := range lines {
			if readingPositional {
				positionalLines = append(positionalLines, line)
			} else if strings.HasPrefix(line, "declare -x ") {
				pair := strings.SplitN(strings.TrimPrefix(line, "declare -x "), "=", 2)
				if len(pair) > 1 {
					env.Set(pair[0], strings.Trim(pair[1], "\""))
				}
			} else if strings.HasPrefix(line, "anyenv -x ") {
				pair := strings.SplitN(strings.TrimPrefix(line, "anyenv -x "), "=", 2)
				if len(pair) > 1 {
					switch pair[0] {
					case "RBENV_COMMAND":
						exeName = pair[1]
					case "RBENV_COMMAND_PATH":
						exePath = utils.NewPathname(pair[1])
					case "RBENV_BIN_PATH":
						binPath = utils.NewPathname(pair[1])
					}
				}
			} else if line == "---" {
				readingPositional = true
			}
		}

		positionalSerialized := strings.Join(positionalLines, "\n")
		argv = strings.Split(strings.TrimSuffix(positionalSerialized, "\x00"), "\x00")
	}

	if !currentVersion.IsSystem() {
		env.Set("PATH", binPath.String()+":"+env.Get("PATH"))
	}

	argv = append([]string{exeName}, argv...)

	err = syscall.Exec(exePath.String(), argv, env.Environ())
	if err != nil {
		cli.Errorf("%s: %s\n", args.ProgramName(), err)
		cli.Exit(1)
	}
}

func init() {
	cli.Register("exec", execCmd, execHelp)
}
