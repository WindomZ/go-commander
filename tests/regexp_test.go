package tests

import (
	"regexp"
	"testing"
)

func RegexpCommand(str string) []string {
	return regexp.MustCompile(`[A-Za-z0-9_-]+\b`).FindAllString(
		regexp.MustCompile(`^[A-Za-z0-9_|\(\)\s-]+`).FindString(str), -1)
}

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

func RegexpArgument(str string) []string {
	return regexp.MustCompile(`(?i:<|\[)[A-Za-z0-9_\[\]<>-]+\b(?i:>|])`).
		FindAllString(str, -1)
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

func RegexpOption(str string) []string {
	return regexp.MustCompile(`-{1,2}[A-Za-z0-9_-]+\b`).
		FindAllString(regexp.MustCompile(`(<|\[)[A-Za-z0-9_\[\]<>-]+\b(>|])`).
			ReplaceAllString(str, ""), -1)
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
