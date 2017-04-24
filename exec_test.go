package commander

import (
	"github.com/WindomZ/testify/assert"
	"strings"
	"testing"
)

func Test_Exec_ExecPipeCommand(t *testing.T) {
	stdout, stderr, err := Exec.ExecPipeCommand("")
	assert.Error(t, err)
	assert.Empty(t, stderr)
	assert.Empty(t, stdout)

	stdout, stderr, err = Exec.ExecPipeCommand("xxx", "xxx")
	assert.Error(t, err)
	assert.Empty(t, stderr)
	assert.Empty(t, stdout)

	stdout, stderr, err = Exec.ExecPipeCommand("ls")
	assert.NoError(t, err)
	assert.Empty(t, stderr)
	assert.NotEmpty(t, stdout)
	assert.True(t, strings.Contains(stdout, ".go"))
}

func Test_Exec_ExecStdCommand(t *testing.T) {
	stdout, err := Exec.ExecStdCommand("")
	assert.Error(t, err)
	assert.Empty(t, stdout)

	stdout, err = Exec.ExecStdCommand("xxx", "xxx")
	assert.Error(t, err)
	assert.Empty(t, stdout)

	stdout, err = Exec.ExecStdCommand("ls", "-d")
	assert.NoError(t, err)
	assert.NotEmpty(t, stdout)
	assert.True(t, strings.Contains(stdout, "."))
}

func Test_Exec_ExecPipeStatementCommand(t *testing.T) {
	stdout, stderr, err := Exec.ExecPipeStatementCommand("ls -d")
	assert.NoError(t, err)
	assert.Empty(t, stderr)
	assert.NotEmpty(t, stdout)
	assert.True(t, strings.Contains(stdout, "."))
}

func Test_Exec_ExecStdStatementCommand(t *testing.T) {
	stdout, err := Exec.ExecStdStatementCommand("ls -d")
	assert.NoError(t, err)
	assert.NotEmpty(t, stdout)
	assert.True(t, strings.Contains(stdout, "."))
}

func Test_Exec_SplitCommandStatement(t *testing.T) {
	strs := Exec.SplitCommandStatement(`rm -rf "$HOME/user/Download"`)
	assert.Equal(t, strs, []string{
		"rm",
		"-rf",
		"\"$HOME/user/Download\"",
	})
}
