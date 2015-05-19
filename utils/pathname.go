package utils

import (
	"os"
	"path"
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

func (p Pathname) Join(name string) Pathname {
	return NewPathname(path.Join(p.Path, name))
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

func (p Pathname) Exists() bool {
	fileInfo, err := os.Stat(p.Path)
	return err == nil && !fileInfo.IsDir()
}

func (p Pathname) Equal(other Pathname) bool {
	return p.Path == other.Path
}

func (p Pathname) Entries() []Pathname {
	file, err := os.Open(p.Path)
	if err == nil {
		entries, err := file.Readdirnames(0)
		if err == nil {
			results := make([]Pathname, len(entries))
			for i, entry := range entries {
				results[i] = p.Join(entry)
			}
			return results
		}
	}
	return []Pathname{}
}
