package commander

import "strings"

type _Argv []string

func newArgv(args []string) _Argv {
	if args != nil && len(args) > 1 {
		return _Argv(args[1:])
	}
	return _Argv([]string{})
}

func (a _Argv) GetArg(index int) string {
	if index >= 0 && index < len(a) {
		return a[index]
	}
	return ""
}

func (a _Argv) GetArgs(offsets ...int) []string {
	var offset int = 0
	if len(offsets) != 0 {
		offset = offsets[0]
	}
	if len(a) <= offset {
		return []string{}
	}
	if offset < 0 {
		offset = 0
	}
	return a[offset:]
}

func (a _Argv) ArgsString() string {
	return strings.Join(a, " ")
}

func (a _Argv) ArgsStringSeparator(sep string, offsets ...int) string {
	return strings.Join(a.GetArgs(offsets...), sep)
}
