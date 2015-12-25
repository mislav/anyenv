package utils

import (
	"os"
	"strings"
)

type Env struct {
	Data map[string]string
}

func (e *Env) Get(key string) string {
	return e.Data[key]
}

func (e *Env) Set(key string, value string) {
	e.Data[key] = value
}

func (e *Env) Unset(key string) {
	delete(e.Data, key)
}

func (e *Env) Clone() *Env {
	env := Env{make(map[string]string)}
	for key, value := range e.Data {
		env.Set(key, value)
	}
	return &env
}

func (e *Env) Environ() []string {
	result := make([]string, len(e.Data))
	i := 0
	for key, value := range e.Data {
		result[i] = key + "=" + value
		i += 1
	}
	return result
}

func EnvFromEnviron() *Env {
	env := Env{make(map[string]string)}
	for _, line := range os.Environ() {
		pair := strings.SplitN(line, "=", 2)
		env.Set(pair[0], pair[1])
	}
	return &env
}
