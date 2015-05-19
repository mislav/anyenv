package cli

var commands = make(map[string]func(Args))

func Lookup(cmdName string) func(Args) {
	return commands[cmdName]
}

func Register(cmdName string, fn func(Args)) {
	commands[cmdName] = fn
}
