package utils

import (
	"sort"
)

type Set struct {
	Data map[string]bool
}

func (s Set) Add(value string) {
	s.Data[value] = true
}

func (s Set) Array() []string {
	values := make([]string, len(s.Data))
	i := 0
	for value, _ := range s.Data {
		values[i] = value
		i += 1
	}
	return values
}

func (s Set) Sort() []string {
	values := s.Array()
	sort.Strings(values)
	return values
}

func NewSet() Set {
	return Set{make(map[string]bool)}
}

func NewSetFromSlice(s []string) Set {
	set := NewSet()
	for _, value := range s {
		set.Add(value)
	}
	return set
}
