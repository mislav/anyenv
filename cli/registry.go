package cli

var commands = make(map[string]func([]string))

func Lookup(cmdName string) func([]string) {
	return commands[cmdName]
}

func Register(cmdName string, fn func([]string)) {
	commands[cmdName] = fn
}
