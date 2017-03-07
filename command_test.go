package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestCommand_1(t *testing.T) {
	c := newCommand("cmd <x>", true).
		Version("v1.0.1").
		Description("this is description").
		Option("-c, --config", "config description").
		Option("-d, --drop", "drop description")

	c.Command("cmd2").
		Option("-a, --about", "cmd2 about description").
		Option("-t, --test", "cmd2 test description")

	c.Command("cmd3 [y]").
		Option("-b, --bold=<kn>", "cmd3 bold description").
		Option("-c, --count", "cmd3 count description")

	assert.Equal(t, c.UsagesString(),
		[]string([]string{
			"cmd <x> [-c|--config] [-d|--drop]",
			"cmd cmd2 [-a|--about] [-t|--test]",
			"cmd cmd3 [y] [(-b|--bold)=<kn>] [-c|--count]",
			"cmd -h | --help",
			"cmd --version",
		}))
	assert.Equal(t, c.OptionsString(),
		[]string([]string{
			"-c, --config  config description",
			"-d, --drop    drop description",
		}))

	assert.Equal(t, c.GetHelpMessage(),
		`this is description

Usage:
  cmd <x> [-c|--config] [-d|--drop]
  cmd cmd2 [-a|--about] [-t|--test]
  cmd cmd3 [y] [(-b|--bold)=<kn>] [-c|--count]
  cmd -h | --help
  cmd --version

Options:
  -c, --config  config description
  -d, --drop    drop description
`)
}
