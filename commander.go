package commander

import "github.com/docopt/docopt-go"

type Commander interface {
	ICommand
}

func NewCommander(usage string, args ...interface{}) Commander {
	return newCommand(usage, true, args...)
}

func Parse(doc string, argv []string, help bool, version string,
	optionsFirst bool, exit ...bool) (DocoptMap, error) {
	m, err := docopt.Parse(doc, argv, help, version, optionsFirst, exit...)
	return DocoptMap(m), err
}
