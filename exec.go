package commander

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// Execute default Execute
var Exec Execute = newExec()

// Execute the struct has the command line function set implementation
type Execute interface {
	ExecPipeCommand(name string, arg ...string) (string, string, error)
	ExecStdCommand(name string, arg ...string) (string, error)
	ExecPipeStatementCommand(statement string) (string, string, error)
	ExecStdStatementCommand(statement string) (string, error)
	SplitCommandStatement(statement string) (result []string)
}

type _Exec struct {
}

func newExec() *_Exec {
	return &_Exec{}
}

// ExecPipeCommand new a Cmd struct to execute the named program with the given arguments.
//
// If name contains no path separators, Command uses LookPath to
// resolve name to a complete path if possible. Otherwise it uses name
// directly as Path.
//
// Cmd's Args field is constructed from the command name
// followed by the elements of arg, so arg should not include the
// command name itself. For example, Command("echo", "hello").
// Args[0] is always name, not the possibly resolved Path.
//
// First returns a string that collected from Cmd's Stdout.
// Second returns a string that collected from Cmd's Stderr.
func (e _Exec) ExecPipeCommand(name string, arg ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(name, arg...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		return stdout.String(), stderr.String(), err
	}
	if err := cmd.Wait(); err != nil {
		return stdout.String(), stderr.String(), err
	}
	return stdout.String(), stderr.String(), nil
}

// ExecStdCommand usage is similar to 'ExecStdCommand'
// But returns a string collected from Cmd's Stdout and Stderr.
func (e _Exec) ExecStdCommand(name string, arg ...string) (string, error) {
	stdout, stderr, err := e.ExecPipeCommand(name, arg...)
	os.Stdout.WriteString(stdout)
	os.Stderr.WriteString(stderr)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return strings.TrimSpace(stdout + stderr), err
	} else if len(stderr) != 0 {
		return stdout, errors.New(stderr)
	}
	return stdout, nil
}

// ExecPipeStatementCommand Cmd based bash shell to exec statement script.
// First returns a string that collected from Cmd's Stdout.
// Second returns a string that collected from Cmd's Stderr.
func (e _Exec) ExecPipeStatementCommand(statement string) (string, string, error) {
	return e.ExecPipeCommand(
		"/bin/bash",
		"-c",
		statement,
	)
}

// ExecStdStatementCommand usage is similar to 'ExecPipeStatementCommand'
// But returns a string collected from Cmd's Stdout and Stderr.
func (e _Exec) ExecStdStatementCommand(statement string) (string, error) {
	return e.ExecStdCommand(
		"/bin/bash",
		"-c",
		statement,
	)
}

// SplitCommandStatement returns a slice of string, split the statement to scripts.
func (e _Exec) SplitCommandStatement(statement string) (result []string) {
	result = regexp.MustCompile(`([^\\]"[\w\s\S]+?[^\\]")|([^\s\\]+)`).
		FindAllString(statement, -1)
	for i, str := range result {
		result[i] = strings.TrimSpace(str)
	}
	return
}
