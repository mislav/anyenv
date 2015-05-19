package cli

import (
	"path"
)

type Args struct {
	ARGV []string
}

func (a Args) ProgramName() string {
	return path.Base(a.ARGV[0])
}

func (a Args) CommandName() string {
	return a.At(-1)
}

func (a Args) At(n int) string {
	n += 2
	if len(a.ARGV) > n {
		return a.ARGV[n]
	} else {
		return ""
	}
}

func (a Args) Required(n int) string {
	value := a.At(n)
	if value == "" {
		Errorln(HelpText(a.CommandName(), a.ProgramName()))
		Exit(1)
	}
	return value
}

func (a Args) HasFlag(flag string) bool {
	for _, arg := range a.ARGV {
		if arg == flag {
			return true
		}
	}
	return false
}
