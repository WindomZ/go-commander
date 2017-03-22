package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func Test_Argv_GetArg(t *testing.T) {
	var argv _Argv = newArgv([]string{
		"aaa",
		"bbb",
		"ccc",
	})
	assert.Equal(t, argv.GetArg(0), "bbb")
	assert.Equal(t, argv.GetArg(1), "ccc")
	assert.Equal(t, argv.GetArg(2), "")
}

func Test_Argv_GetArgs(t *testing.T) {
	var argv _Argv = newArgv([]string{
		"aaa",
		"bbb",
		"ccc",
	})
	assert.Equal(t, argv.GetArgs(), []string{"bbb", "ccc"})
	assert.Equal(t, argv.GetArgs(1), []string{"ccc"})
	assert.Equal(t, argv.GetArgs(2), []string{})
}

func Test_Argv_ArgsString(t *testing.T) {
	var argv _Argv = newArgv([]string{
		"aaa",
		"bbb",
		"ccc",
	})
	assert.Equal(t, argv.ArgsString(), "bbb ccc")
}

func Test_Argv_ArgsStringSeparator(t *testing.T) {
	var argv _Argv = newArgv([]string{
		"aaa",
		"bbb",
		"ccc",
	})
	assert.Equal(t, argv.ArgsStringSeparator(", "), "bbb, ccc")
	assert.Equal(t, argv.ArgsStringSeparator("|"), "bbb|ccc")
	assert.Equal(t, argv.ArgsStringSeparator("$"), "bbb$ccc")
	assert.Equal(t, argv.ArgsStringSeparator("$", 1), "ccc")
}
