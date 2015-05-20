package commands

import (
	"github.com/mislav/everyenv/cli"
	"github.com/mislav/everyenv/config"
	"strings"
)

var initHelp = `
Usage: eval "$($ProgramName init - [--no-rehash] [<shell>])"

Configure the shell environment for $ProgramName
`

func initCmd(args cli.Args) {
	evalMode := args.At(0) == "-"
	skipRehash := args.HasFlag("--no-rehash")
	shell := cli.DetectShell(args.Word(0))
	programName := args.ProgramName()

	if !evalMode {
		cli.Errorf("# Load %s automatically by adding\n", programName)
		cli.Errorf("# the following to %s:\n\n", profileName(shell))
		if shell == "fish" {
			cli.Errorf("status --is-interactive; and . (%s init -|psub)\n", programName)
		} else {
			cli.Errorf("eval \"$(%s init -)\"\n", programName)
		}
		cli.Exit(1)
	} else {
		if shell == "fish" {
			cli.Printf("setenv PATH '%s' $PATH\n", config.ShimsDir().String())
			cli.Printf("setenv %s %s\n", config.ShellEnvName, shell)
		} else {
			cli.Printf("export PATH=\"%s:${PATH}\"\n", config.ShimsDir().String())
			cli.Printf("export %s=%s\n", config.ShellEnvName, shell)
		}

		if !skipRehash {
			cli.Printf("%s rehash 2>/dev/null\n", programName)
		}

		shCommands := []string{}
		for _, cmd := range findAvailableCommands(args.ProgramName()).Array() {
			if strings.HasPrefix(cmd, "sh-") {
				shCommands = append(shCommands, strings.TrimPrefix(cmd, "sh-"))
			}
		}

		if shell == "fish" {
			cli.Printf(`function %s
  set command $argv[1]
  set -e argv[1]

  switch "$command"
  case %s
    eval (%s "sh-$command" $argv)
  case '*'
    command %s "$command" $argv
  end
end
`, programName, strings.Join(shCommands, " "), programName, programName)
		} else {
			if shell == "ksh" {
				cli.Printf("function %s {\n typeset command\n", programName)
			} else {
				cli.Printf("%s() {\n local command\n", programName)
			}

			cli.Printf(` command="$1"
  if [ "$#" -gt 0 ]; then
    shift
  fi

  case "$command" in
  %s)
    eval "%s";;
  *)
    command %s "$command" "$@";;
  esac
}
`, strings.Join(shCommands, "|"), "`"+programName+" \"sh-$command\" \"$@\"`", programName)
		}
	}
}

func profileName(shell string) string {
	switch shell {
	case "bash":
		return "~/.bash_profile"
	case "zsh":
		return "~/.zshrc"
	case "ksh":
		return "~/.profile"
	case "fish":
		return "~/.config/fish/config.fish"
	default:
		return "your profile"
	}
}

func init() {
	cli.Register("init", initCmd, initHelp)
}
