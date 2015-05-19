package cli

import (
	"fmt"
	"os"
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
