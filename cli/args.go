package cli

type Args struct {
	List []string
}

func (a Args) HasFlag(flag string) bool {
	for _, arg := range a.List {
		if arg == flag {
			return true
		}
	}
	return false
}
