package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestRegexpCommand(t *testing.T) {
	assert.Equal(t,
		RegexpCommand("new <name>"),
		[]string{"new"},
	)
	assert.Equal(t,
		RegexpCommand("ship <name> move <x> <y>"),
		[]string{"ship"},
	)
	assert.Equal(t,
		RegexpCommand("(set|remove) <x> <y> [--moored|--drifting]"),
		[]string{"set", "remove"},
	)
}

func TestRegexpArgument(t *testing.T) {
	assert.Equal(t,
		RegexpArgument("new <name>"),
		[]string{"<name>"},
	)
	assert.Equal(t,
		RegexpArgument("ship <name> move <x> <y>"),
		[]string{"<name>", "<x>", "<y>"},
	)
	assert.Equal(t,
		RegexpArgument("(set|remove) <x> <y> [--moored|--drifting]"),
		[]string{"<x>", "<y>"},
	)
}

func TestRegexpOption(t *testing.T) {
	assert.Equal(t,
		RegexpOption("new <name>"),
		[]string(nil),
	)
	assert.Equal(t,
		RegexpOption("-p <x-y>"),
		[]string{"-p"},
	)
	assert.Equal(t,
		RegexpOption("-p"),
		[]string{"-p"},
	)
	assert.Equal(t,
		RegexpOption("-p, --pepper"),
		[]string{"-p", "--pepper"},
	)
	assert.Equal(t,
		RegexpOption("--pepper"),
		[]string{"--pepper"},
	)
	assert.Equal(t,
		RegexpOption("(set|remove) <x> <y> [--not-ss | -a | --moored|--drifting]"),
		[]string{"--not-ss", "-a", "--moored", "--drifting"},
	)
}
