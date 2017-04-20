package commander

import (
	"github.com/WindomZ/testify/assert"
	"strings"
	"testing"
)

func TestExec_ExecPipeCommand(t *testing.T) {
	stdout, stderr, err := ExecPipeCommand("")
	assert.Error(t, err)
	assert.Empty(t, stderr)
	assert.Empty(t, stdout)

	stdout, stderr, err = ExecPipeCommand("xxx", "xxx")
	assert.Error(t, err)
	assert.Empty(t, stderr)
	assert.Empty(t, stdout)

	stdout, stderr, err = ExecPipeCommand("ls")
	assert.NoError(t, err)
	assert.Empty(t, stderr)
	assert.NotEmpty(t, stdout)
	assert.True(t, strings.Contains(stdout, ".go"))
}

func TestExec_ExecStdCommand(t *testing.T) {
	stdout, err := ExecStdCommand("")
	assert.Error(t, err)
	assert.Empty(t, stdout)

	stdout, err = ExecStdCommand("xxx", "xxx")
	assert.Error(t, err)
	assert.Empty(t, stdout)

	stdout, err = ExecStdCommand("ls", "-d")
	assert.NoError(t, err)
	assert.NotEmpty(t, stdout)
	assert.True(t, strings.Contains(stdout, "."))
}

func TestExec_ExecPipeStatementCommand(t *testing.T) {
	stdout, stderr, err := ExecPipeStatementCommand("ls -d")
	assert.NoError(t, err)
	assert.Empty(t, stderr)
	assert.NotEmpty(t, stdout)
	assert.True(t, strings.Contains(stdout, "."))
}

func TestExec_ExecStdStatementCommand(t *testing.T) {
	stdout, err := ExecStdStatementCommand("ls -d")
	assert.NoError(t, err)
	assert.NotEmpty(t, stdout)
	assert.True(t, strings.Contains(stdout, "."))
}

func TestExec_SplitCommandStatement(t *testing.T) {
	strs := SplitCommandStatement(`rm -rf "$HOME/user/Download"`)
	assert.Equal(t, strs, []string{
		"rm",
		"-rf",
		"\"$HOME/user/Download\"",
	})
}
