package commander

import "github.com/docopt/docopt-go"

type Commander interface {
	Version(ver string) Commander
	Description(desc string) Commander
	Action(action Action) Commander
	Command(usage string, args ...interface{}) Commander
	Option(usage string, args ...interface{}) Commander
	LineArgument(usage string, args ...interface{}) Commander
	LineOption(usage string, args ...interface{}) Commander
	UsagesString() []string
	OptionsString() []string
	GetHelpMessage() string
	ShowHelpMessage() string
	Parse(argv ...[]string) (*Context, error)
}

func NewCommander(usage string, args ...interface{}) Commander {
	return newCommand(usage, true, args...)
}

func Parse(doc string, argv []string, help bool, version string,
	optionsFirst bool, exit ...bool) (DocoptMap, error) {
	if argv != nil && len(argv) > 1 {
		argv = argv[1:]
	}
	m, err := docopt.Parse(doc, argv, help, version, optionsFirst, exit...)
	return DocoptMap(m), err
}
