package commander

import "github.com/docopt/docopt-go"

type Commander interface {
	Version(ver string, flags ...string) Commander
	Description(desc string) Commander
	Usage(usage string) Commander
	Command(usage string, args ...interface{}) Commander
	Alias(alias string) Commander
	Option(flags string, args ...interface{}) Commander
	UsagesString() []string
	OptionsString() []string
	GetHelpMessage() string
	Parse() (map[string]interface{}, error)
}

func NewCommander(name string) Commander {
	return newCommand(name, true)
}

func Parse(doc string, argv []string, help bool, version string,
	optionsFirst bool, exit ...bool) (map[string]interface{}, error) {
	return docopt.Parse(doc, argv, help, version, optionsFirst, exit...)
}
