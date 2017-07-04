package commander

import "strings"

// _Argv the collection of command args strings.
type _Argv []string

// newArgv returns new instance of _Argv.
func newArgv(args []string) _Argv {
	if args != nil && len(args) > 1 {
		return _Argv(args[1:])
	}
	return _Argv([]string{})
}

// GetArg returns a arg string by index if that less than size of _Argv, otherwise returns empty string.
func (a _Argv) GetArg(index int) string {
	if index >= 0 && index < len(a) {
		return a[index]
	}
	return ""
}

// GetArgs returns the arg strings offset by offsets which if not empty.
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

// ArgsString Similar to "toString" and join with a space " ".
func (a _Argv) ArgsString() string {
	return strings.Join(a, " ")
}

// ArgsStringSeparator Similar to "ArgsString", but defines separator string sep.
// offsets used for offset _Argv.
func (a _Argv) ArgsStringSeparator(sep string, offsets ...int) string {
	return strings.Join(a.GetArgs(offsets...), sep)
}
