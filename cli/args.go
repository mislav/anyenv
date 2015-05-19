package cli

func HasFlag(flag string, args []string) bool {
	for _, arg := range args {
		if arg == flag {
			return true
		}
	}
	return false
}
