package commander

import "strings"

// _Argument Implementation of command line parameter
type _Argument struct {
	name string
}

func newArgument(name string) *_Argument {
	return &_Argument{
		name: strings.TrimSpace(name),
	}
}

func (a _Argument) Name() string {
	return a.name
}

func (a _Argument) UsageString() string {
	return a.name
}

func (a _Argument) OptionString() string {
	return a.name
}
