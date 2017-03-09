package commander

import "strings"

type ContextArgs []string

func newContextArgs(args []string) ContextArgs {
	if args != nil && len(args) > 1 {
		return ContextArgs(args[1:])
	}
	return ContextArgs([]string{})
}

func (c ContextArgs) Get(index int) string {
	if index >= 0 && index < len(c) {
		return c[index]
	}
	return ""
}

func (c ContextArgs) String(offsets ...int) string {
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
	return strings.Join(c[offset:], " | ")
}
