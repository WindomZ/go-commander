package tests

import "testing"

func TestRegexpCommand(t *testing.T) {
	t.Logf("%#v",
		RegexpCommand("new <name>"),
	)
	t.Logf("%#v",
		RegexpCommand("ship <name> move <x> <y>"),
	)
	t.Logf("%#v",
		RegexpCommand("(set|remove) <x> <y> [--moored|--drifting]"),
	)
}

func TestRegexpArgument(t *testing.T) {
	t.Logf("%#v",
		RegexpArgument("new <name>"),
	)
	t.Logf("%#v",
		RegexpArgument("ship <name> move <x> <y>"),
	)
	t.Logf("%#v",
		RegexpArgument("(set|remove) <x> <y> [--moored|--drifting]"),
	)
}

func TestRegexpOption(t *testing.T) {
	t.Logf("%#v",
		RegexpOption("new <name>"),
	)
	t.Logf("%#v",
		RegexpOption("-p <x-y>"),
	)
	t.Logf("%#v",
		RegexpOption("-p"),
	)
	t.Logf("%#v",
		RegexpOption("-p, --pepper"),
	)
	t.Logf("%#v",
		RegexpOption("--pepper"),
	)
	t.Logf("%#v",
		RegexpOption("(set|remove) <x> <y> [--not-ss | -a | --moored|--drifting]"),
	)
}
