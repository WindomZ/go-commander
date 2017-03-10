package commander

import "strings"

type _ContextArguments []string

func newContextArguments(args []string) _ContextArguments {
	if args != nil && len(args) > 1 {
		return _ContextArguments(args[1:])
	}
	return _ContextArguments([]string{})
}

func (c _ContextArguments) GetArg(index int) string {
	if index >= 0 && index < len(c) {
		return c[index]
	}
	return ""
}

func (c _ContextArguments) ArgsString() string {
	return strings.Join(c, " ")
}

func (c _ContextArguments) ArgsStringSeparator(sep string, offsets ...int) string {
	var offset int = 0
	if len(offsets) != 0 {
		offset = offsets[0]
	}
	if len(c) <= offset {
		offset = len(c) - 1
	}
	if offset < 0 {
		offset = 0
	}
	return strings.Join(c[offset:], sep)
}
