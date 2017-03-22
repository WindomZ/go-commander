package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestRegexpCommand(t *testing.T) {
	assert.Equal(t,
		regexpCommand("new <name>"),
		[]string{"new"},
	)
	assert.Equal(t,
		regexpCommand("ship <name> move <x> <y>"),
		[]string{"ship"},
	)
	assert.Equal(t,
		regexpCommand("(set|remove) <x> <y> [--moored|--drifting]"),
		[]string{"set", "remove"},
	)
}

func TestRegexpArgument(t *testing.T) {
	assert.Equal(t,
		regexpArgument("new <name>"),
		[]string{"<name>"},
	)
	assert.Equal(t,
		regexpArgument("ship <name> move <x> <y>"),
		[]string{"<name>", "<x>", "<y>"},
	)
	assert.Equal(t,
		regexpArgument("(set|remove) <x> <y> [--moored|--drifting]"),
		[]string{"<x>", "<y>"},
	)
}

func TestRegexpOption(t *testing.T) {
	assert.Equal(t,
		regexpOption("new <name>"),
		[]string(nil),
	)
	assert.Equal(t,
		regexpOption("-p <x-y>"),
		[]string{"-p"},
	)
	assert.Equal(t,
		regexpOption("-p"),
		[]string{"-p"},
	)
	assert.Equal(t,
		regexpOption("-p, --pepper"),
		[]string{"-p", "--pepper"},
	)
	assert.Equal(t,
		regexpOption("--pepper"),
		[]string{"--pepper"},
	)
	assert.Equal(t,
		regexpOption("(set|remove) <x> <y> [--not-ss | -a | --moored|--drifting]"),
		[]string{"--not-ss", "-a", "--moored", "--drifting"},
	)
}

func TestReplaceCommand(t *testing.T) {
	assert.Equal(t,
		replaceCommand("new <name>", "new", "(new|n)"),
		"(new|n) <name>",
	)
	assert.Equal(t,
		replaceCommand("(new|n) <name>", "(new|n)", "(new|n|add)"),
		"(new|n|add) <name>",
	)
}
