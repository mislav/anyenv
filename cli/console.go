package cli

import (
	"fmt"
	"github.com/mislav/everyenv/config"
	"os"
	"path"
	"strings"
)

var (
	Stdout = os.Stdout
	Stderr = os.Stderr
)

func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(Stdout, format, a...)
}

func Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(Stdout, a...)
}

func Errorf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(Stderr, format, a...)
}

func Errorln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(Stderr, a...)
}

func Exit(code int) {
	os.Exit(code)
}

func DetectShell(shell string) string {
	if shell == "" {
		shell = os.Getenv(config.ShellEnvName)
	}

	if shell == "" {
		shell = strings.TrimLeft(path.Base(os.Getenv("SHELL")), "-")
	}

	if shell == "" {
		shell = "bash"
	}

	return shell
}
