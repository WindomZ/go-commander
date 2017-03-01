package commander

import (
	"fmt"
	"strings"
)

type Command struct {
	name     string
	args     string
	alias    []string
	desc     string
	version  Version
	usage    Usage
	commands []Command
	options  []Option
	exec     Exec
	errFunc  ErrFunc
}

func newCommand(cmd string) *Command {
	return &Command{
		name: commandName(cmd),
		args: commandArguments(cmd),
		desc: "",
		errFunc: func(err error, obj interface{}) {
			fmt.Printf("  err: %v\n  object: %#v\n", err, obj)
		},
	}
}

func commandName(cmd string) string {
	cmd = strings.TrimSpace(cmd)
	if i := strings.Index(cmd, " "); i >= 0 {
		cmd = cmd[:i]
	}
	return cmd
}

func commandArguments(cmd string) string {
	cmd = strings.TrimSpace(cmd)
	if i := strings.Index(cmd, " "); i >= 0 {
		cmd = cmd[i+1:]
	}
	return cmd
}

func (c Command) Name() string {
	if alias := strings.Join(c.alias, "|"); len(alias) != 0 {
		return fmt.Sprintf("(%v|%v)", c.name, alias)
	}
	return c.name
}

func (c Command) hasExec() bool {
	return c.exec != nil
}

func (c *Command) Version(ver string, flags ...string) Commander {
	c.version.Set(ver)
	return c
}

func (c *Command) Description(desc string) Commander {
	c.desc = desc
	return c
}

func (c *Command) Usage(usage string) Commander {
	c.usage.Set(usage)
	return c
}

func (c *Command) Command(name string) Commander {
	cmd := newCommand(name)
	if cmd.Valid() {
		c.commands = append(c.commands, *cmd)
	} else if c.errFunc != nil {
		c.errFunc(ErrCommand, cmd)
	}
	return cmd
}

func (c *Command) Alias(alias string) Commander {
	if alias = commandName(alias); len(alias) != 0 {
		c.alias = append(c.alias, alias)
	}
	return c
}

func (c *Command) Option(flags, desc string) Commander {
	if opt := newOption(flags, desc); opt.Valid() {
		c.options = append(c.options, *opt)
	} else if c.errFunc != nil {
		c.errFunc(ErrOption, opt)
	}
	return c
}

func (c Command) Valid() bool {
	return len(c.name) != 0
}

func (c Command) UsageString() (r []string) {
	if !c.Valid() {
		return
	}
	r = append(r, fmt.Sprintf("%v %v", c.Name(), c.args))
	// TODO: Not finish
	r = append(r,
		fmt.Sprintf("%v -h | --help", c.Name()),
		fmt.Sprintf("%v --version", c.Name()),
	)
	return
}

func (c Command) OptionsString() (r []string) {
	if !c.Valid() {
		return
	}
	for _, opt := range c.options {
		r = append(r, opt.OptionString())
	}
	for _, cmd := range c.commands {
		r = append(r, cmd.OptionsString()...)
	}
	return
}
