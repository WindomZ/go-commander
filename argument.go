package commander

import "strings"

// _Argument command line parameter implementation.
type _Argument struct {
	name string
}

// newArgument returns new instance of _Argument.
func newArgument(name string) *_Argument {
	return &_Argument{
		name: strings.TrimSpace(name),
	}
}

//func (a _Argument) Name() string {
//	return a.name
//}
//
//func (a _Argument) UsageString() string {
//	return a.name
//}
//
//func (a _Argument) OptionString() string {
//	return a.name
//}
