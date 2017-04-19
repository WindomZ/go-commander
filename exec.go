package commander

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func ExecPipeCommand(name string, arg ...string) (string, string, error) {
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

func ExecStdCommand(name string, arg ...string) (string, error) {
	stdout, stderr, err := ExecPipeCommand(name, arg...)
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

func ExecPipeStatementCommand(statement string) (string, string, error) {
	return ExecPipeCommand(
		"/bin/bash",
		"-c",
		statement,
	)
}

func ExecStdStatementCommand(statement string) (string, error) {
	return ExecStdCommand(
		"/bin/bash",
		"-c",
		statement,
	)
}

func SplitCommandStatement(statement string) (result []string) {
	result = regexp.MustCompile(`([^\\]"[\w\s\S]+?[^\\]")|([^\s\\]+)`).
		FindAllString(statement, -1)
	for i, str := range result {
		result[i] = strings.TrimSpace(str)
	}
	return
}
