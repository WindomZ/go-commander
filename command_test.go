package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestCommand_1(t *testing.T) {
	c := newCommand(true).
		Usage("cmd <x>", "this is description").
		Option("-c, --config", "config description").
		Option("-d, --drop", "drop description")

	c.Command("cmd2").
		Option("-a, --about", "cmd2 about description").
		Option("-t, --test", "cmd2 test description")

	c.Command("cmd3 [y]").
		Option("-b=<kn>, --bold=<kn>", "cmd3 bold description").
		Option("-c, --count", "cmd3 count description")

	assert.Equal(t, c.UsagesString(),
		[]string{
			"cmd <x> [-c|--config] [-d|--drop]",
			"cmd cmd2 [-a|--about] [-t|--test]",
			"cmd cmd3 [y] [-b=<kn>|--bold=<kn>] [-c|--count]",
			"cmd -h | --help",
			"cmd --version",
		})
	assert.Equal(t, c.OptionsString(),
		[]string{
			"-c --config   config description",
			"-d --drop     drop description",
			"-a --about    cmd2 about description",
			"-t --test     cmd2 test description",
			"-b=<kn> --bold=<kn>  cmd3 bold description",
			"-c --count    cmd3 count description",
		})

	assert.Equal(t, c.HelpMessage(),
		`this is description

Usage:
  cmd <x> [-c|--config] [-d|--drop]
  cmd cmd2 [-a|--about] [-t|--test]
  cmd cmd3 [y] [-b=<kn>|--bold=<kn>] [-c|--count]
  cmd -h | --help
  cmd --version

Options:
  -c --config   config description
  -d --drop     drop description
  -a --about    cmd2 about description
  -t --test     cmd2 test description
  -b=<kn> --bold=<kn>  cmd3 bold description
  -c --count    cmd3 count description
`)
}
