package commander

import "testing"

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

	usageStrs := c.UsagesString()
	for _, str := range usageStrs {
		t.Logf("Usage: %s", str)
	}

	optStrs := c.OptionsString()
	for _, str := range optStrs {
		t.Logf("Options: %s", str)
	}

	t.Log(c.GetUsage())
}
