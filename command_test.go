package commander

import "testing"

func TestCommand_1(t *testing.T) {
	c := newCommand("cmd <x>").
		Version("v1.0.1").
		Description("this is description").
		Option("-c, --config", "config description").
		Option("-d, --drop", "drop description")

	usageStrs := c.UsageString()
	for _, str := range usageStrs {
		t.Logf("Usage: %s", str)
	}

	optStrs := c.OptionsString()
	for _, str := range optStrs {
		t.Logf("Options: %s", str)
	}
}
