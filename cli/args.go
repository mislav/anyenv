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

func (a Args) HasFlag(flag string) bool {
	for _, arg := range a.ARGV {
		if arg == flag {
			return true
		}
	}
	return false
}
