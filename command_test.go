package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestCommand_CommandsString(t *testing.T) {
	c := newCommand(true)

	c.Command("test").
		Version("0.0.1").
		Description("this is a test cli.")

	c.Command("add <x> <y>").
		Description("this addition operation")

	c.Command("sub <x> <y>").
		Description("this subtraction operation").
		Command("again").
		Description("this subtraction operation again")

	assert.Equal(t, c.CommandsString(""),
		[]string{
			"add           this addition operation",
			"sub           this subtraction operation",
			"sub again     this subtraction operation again",
		})
}

func TestCommand_HelpMessage(t *testing.T) {
	c := newCommand(true)
	assert.Equal(t, c.HelpMessage(), "")

	c.Usage("cmd <x>", "this is description").
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
		})
	opts := c.OptionsString()
	assert.Equal(t, opts["-c|--config"], "-c --config   config description")
	assert.Equal(t, opts["-d|--drop"], "-d --drop     drop description")
	assert.Equal(t, opts["-a|--about"], "-a --about    cmd2 about description")
	assert.Equal(t, opts["-t|--test"], "-t --test     cmd2 test description")
	assert.Equal(t, opts["-b|--bold"], "-b=<kn> --bold=<kn>\n              cmd3 bold description")
	assert.Equal(t, opts["-c|--count"], "-c --count    cmd3 count description")

	assert.Equal(t, c.HelpMessage(),
		`  this is description

  Usage:
    cmd <x> [-c|--config] [-d|--drop]
    cmd cmd2 [-a|--about] [-t|--test]
    cmd cmd3 [y] [-b=<kn>|--bold=<kn>] [-c|--count]

  Options:
    -a --about    cmd2 about description
    -b=<kn> --bold=<kn>
                  cmd3 bold description
    -c --config   config description
    -c --count    cmd3 count description
    -d --drop     drop description
    -t --test     cmd2 test description
`)

	c.Annotation("Example", []string{
		"cmd xxx -c",
		"cmd cmd2 -a",
		"cmd cmd3 y -b",
	})
	assert.Equal(t, c.HelpMessage(),
		`  this is description

  Usage:
    cmd <x> [-c|--config] [-d|--drop]
    cmd cmd2 [-a|--about] [-t|--test]
    cmd cmd3 [y] [-b=<kn>|--bold=<kn>] [-c|--count]

  Options:
    -a --about    cmd2 about description
    -b=<kn> --bold=<kn>
                  cmd3 bold description
    -c --config   config description
    -c --count    cmd3 count description
    -d --drop     drop description
    -t --test     cmd2 test description

  Example:
    cmd xxx -c
    cmd cmd2 -a
    cmd cmd3 y -b
`)

	c.Doc(`  Usage:
    cmd <x> [-c|--config] [-d|--drop]
    cmd cmd2 [-a|--about] [-t|--test]
    cmd cmd3 [y] [-b=<kn>|--bold=<kn>] [-c|--count]
`)

	assert.Equal(t, c.HelpMessage(), `  Usage:
    cmd <x> [-c|--config] [-d|--drop]
    cmd cmd2 [-a|--about] [-t|--test]
    cmd cmd3 [y] [-b=<kn>|--bold=<kn>] [-c|--count]
`)
}

func TestCommand_Aliases(t *testing.T) {
	c := newCommand(true)
	c.Command("cmd").
		Aliases([]string{"c0", "cmd0"}).
		Option("-a, --about", "cmd about description").
		Option("-t, --test", "cmd test description")
	assert.Equal(t, c.HelpMessage(), `  Usage:
    (cmd|c0|cmd0) [-a|--about] [-t|--test]

  Options:
    -a --about    cmd about description
    -t --test     cmd test description
`)
}

func TestCommand_ShowVersion(t *testing.T) {
	c := newCommand(true)
	c.Version("v0.0.1")

	assert.Empty(t, c.ShowVersion())
}

func TestCommand_ShowHelpMessage(t *testing.T) {
	c := newCommand(true)
	c.Version("v0.0.1").Description("this is a test cli.")

	assert.NotEmpty(t, c.ShowHelpMessage())

	assert.Equal(t, c.HelpMessage(),
		`  this is a test cli.

`)
}

func TestCommand_ErrorHandling(t *testing.T) {
	c := newCommand(true)
	c.ErrorHandling(func(err error) {
		assert.NoError(t, err)
	})

	c.Command("test").Action(func() error {
		return newError("this is a test error")
	})

	if _, err := c.Parse([]string{"go-commander", "test"}); err != nil {
		assert.Error(t, err)
	}
}
