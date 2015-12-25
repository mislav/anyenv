package utils

import (
	"os"
	"sort"
	"strings"
	"testing"
)

func assertEqual(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Fatalf("expected %#v; got %#v\n", expected, actual)
	}
}

func resetEnv() {
	for _, pair := range os.Environ() {
		if strings.HasPrefix(pair, "ANYENV_TESTING") {
			parts := strings.SplitN(pair, "=", 2)
			os.Unsetenv(parts[0])
		}
	}
}

func TestClone(t *testing.T) {
	os.Setenv("ANYENV_TESTING", "hello")
	defer resetEnv()
	env1 := EnvFromEnviron()
	env2 := env1.Clone()

	assertEqual(t, "hello", env1.Get("ANYENV_TESTING"))
	assertEqual(t, "hello", env2.Get("ANYENV_TESTING"))

	env2.Set("ANYENV_TESTING", "world")
	assertEqual(t, "hello", env1.Get("ANYENV_TESTING"))
	assertEqual(t, "world", env2.Get("ANYENV_TESTING"))
}

func TestEnviron(t *testing.T) {
	os.Setenv("ANYENV_TESTING_1", "hello")
	os.Setenv("ANYENV_TESTING_2", "world")
	defer resetEnv()

	env := EnvFromEnviron()

	filtered := []string{}
	for _, pair := range env.Environ() {
		if strings.HasPrefix(pair, "ANYENV_TESTING") {
			filtered = append(filtered, pair)
		}
	}

	sort.Strings(filtered)
	assertEqual(t, 2, len(filtered))
	assertEqual(t, "ANYENV_TESTING_1=hello", filtered[0])
	assertEqual(t, "ANYENV_TESTING_2=world", filtered[1])
}
