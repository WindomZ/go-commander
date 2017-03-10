package commander

import "github.com/docopt/docopt-go"

// Commander Command line implementation
type Commander interface {
	Version(ver string) Commander
	Description(desc string) Commander
	Annotation(title string, contents []string) Commander
	Action(action interface{}, keys ...[]string) Commander
	Command(usage string, args ...interface{}) Commander
	Option(usage string, args ...interface{}) Commander
	LineArgument(usage string, args ...interface{}) Commander
	LineOption(usage string, args ...interface{}) Commander
	UsagesString() []string
	OptionsString() []string
	HelpMessage() string
	ShowHelpMessage() string
	Parse(argv ...[]string) (*Context, error)
}

func newCommander() Commander {
	return newCommand(true)
}

func Parse(doc string, argv []string, help bool, version string,
	optionsFirst bool, exit ...bool) (DocoptMap, error) {
	if argv != nil && len(argv) > 1 {
		argv = argv[1:]
	}
	m, err := docopt.Parse(doc, argv, help, version, optionsFirst, exit...)
	return newDocoptMap(m), err
}
