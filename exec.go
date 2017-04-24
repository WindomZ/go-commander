package commander

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var Exec Execute = newExec()

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

func (e _Exec) ExecPipeStatementCommand(statement string) (string, string, error) {
	return e.ExecPipeCommand(
		"/bin/bash",
		"-c",
		statement,
	)
}

func (e _Exec) ExecStdStatementCommand(statement string) (string, error) {
	return e.ExecStdCommand(
		"/bin/bash",
		"-c",
		statement,
	)
}

func (e _Exec) SplitCommandStatement(statement string) (result []string) {
	result = regexp.MustCompile(`([^\\]"[\w\s\S]+?[^\\]")|([^\s\\]+)`).
		FindAllString(statement, -1)
	for i, str := range result {
		result[i] = strings.TrimSpace(str)
	}
	return
}
