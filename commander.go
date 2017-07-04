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

/*
Parse `argv` based on the command-line interface described in `doc`.
Returns a map of option names to the values parsed from `argv`, and an error or `nil`.
If `argv` is `nil`, `os.Args[1:]` is used.

Set `help` to `false` to disable automatic help messages on `-h` or `--help`.
If `version` is a non-empty string, it will be printed when `--version` is
specified. Set `optionsFirst` to `true` to require that options always come
before positional arguments; otherwise they can overlap.

By default, it calls `os.Exit(0)` if it handled a built-in option such as
`-h` or `--version`. If the user errored with a wrong command or options,
it exits with a return code of 1. To stop it from calling `os.Exit()`
and to handle your own return codes, pass an optional last parameter of `false`
for `exit`.
*/
func Parse(doc string, argv []string, help bool, version string,
	optionsFirst bool, exit ...bool) (DocoptMap, error) {
	if argv != nil && len(argv) > 1 {
		argv = argv[1:]
	}
	m, err := docopt.Parse(doc, argv, help, version, optionsFirst, exit...)
	return newDocoptMap(m), err
}
