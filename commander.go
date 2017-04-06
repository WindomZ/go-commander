package commander

import "github.com/docopt/docopt-go"

// Commander Command line implementation
type Commander interface {
	Doc(doc string) Commander
	Version(ver string) Commander
	ShowVersion() string
	Description(desc string) Commander
	Annotation(title string, contents []string) Commander
	Command(usage string, args ...interface{}) Commander
	Aliases(aliases []string) Commander
	Option(usage string, args ...interface{}) Commander
	Action(action interface{}, keys ...[]string) Commander
	HelpMessage() string
	ShowHelpMessage() string
	Parse(argv ...[]string) (Context, error)
	ErrorHandling(func(error)) Commander
}

func Parse(doc string, argv []string, help bool, version string,
	optionsFirst bool, exit ...bool) (DocoptMap, error) {
	if argv != nil && len(argv) > 1 {
		argv = argv[1:]
	}
	m, err := docopt.Parse(doc, argv, help, version, optionsFirst, exit...)
	return newDocoptMap(m), err
}
