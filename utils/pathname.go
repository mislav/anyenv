package utils

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Pathname struct {
	Path string
}

func NewPathname(p ...string) Pathname {
	fullpath := path.Join(p...)
	return Pathname{fullpath}
}

func Getwd() Pathname {
	pwd, _ := os.Getwd()
	return NewPathname(pwd)
}

func (p Pathname) String() string {
	return p.Path
}

func (p Pathname) Dir() Pathname {
	return NewPathname(path.Dir(p.Path))
}

func (p Pathname) Base() string {
	return path.Base(p.Path)
}

func (p Pathname) Abs() Pathname {
	abs, err := filepath.Abs(p.Path)
	if err == nil {
		return NewPathname(abs)
	} else {
		return p
	}
}

func (p Pathname) Join(names ...string) Pathname {
	components := []string{p.Path}
	components = append(components, names...)
	return NewPathname(path.Join(components...))
}

func (p Pathname) IsBlank() bool {
	return p.Path == ""
}

func (p Pathname) IsRoot() bool {
	return p.Path == "/"
}

func (p Pathname) IsExecutable() bool {
	fileInfo, err := os.Stat(p.Path)
	return err == nil && (fileInfo.Mode()&0111) != 0
}

func (p Pathname) IsFile() bool {
	fileInfo, err := os.Stat(p.Path)
	return err == nil && !fileInfo.IsDir()
}

func (p Pathname) Exists() bool {
	_, err := os.Stat(p.Path)
	return err == nil
}

func (p Pathname) Equal(other Pathname) bool {
	return p.Path == other.Path
}

func (p Pathname) Entries() []Pathname {
	entries := p.BareEntries()
	results := make([]Pathname, len(entries))
	for i, entry := range entries {
		results[i] = p.Join(entry)
	}
	return results
}

func (p Pathname) EntriesMatching(pattern string) []Pathname {
	entries, _ := filepath.Glob(p.Path + "/" + pattern)
	results := make([]Pathname, len(entries))
	for i, entry := range entries {
		results[i] = NewPathname(entry)
	}
	return results
}

func (p Pathname) BareEntries() []string {
	file, err := os.Open(p.Path)
	if err == nil {
		entries, err := file.Readdirnames(0)
		if err == nil {
			return entries
		}
	}
	return []string{}
}

func SearchInPath(pattern string) []Pathname {
	dirs := strings.Split(os.Getenv("PATH"), ":")
	results := []Pathname{}

	for _, p := range dirs {
		dir := NewPathname(p)
		if dir.IsBlank() {
			continue
		}
		for _, cmd := range dir.EntriesMatching(pattern) {
			if cmd.IsExecutable() {
				results = append(results, cmd.Abs())
			}
		}
	}

	return results
}
