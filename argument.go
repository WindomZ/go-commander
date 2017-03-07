package commander

import "strings"

type Argument struct {
	name string
}

func newArgument(name string) *Argument {
	return &Argument{
		name: strings.TrimSpace(name),
	}
}

func (a Argument) Name() string {
	return a.name
}

func (a Argument) UsageString() string {
	return a.name
}

func (a Argument) OptionString() string {
	return a.name
}
