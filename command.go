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
	commands []*Command
	options  []*Option
	exec     Exec
	errFunc  ErrFunc
	root     bool
}

func newCommand(cmd string, root bool) *Command {
	return &Command{
		name: commandName(cmd),
		args: commandArguments(cmd),
		desc: "",
		errFunc: func(err error, obj interface{}) {
			fmt.Printf("  err: %v\n  object: %#v\n", err, obj)
		},
		root: root,
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
		return cmd
	}
	return ""
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
	cmd := newCommand(name, false)
	if cmd.Valid() {
		c.commands = append(c.commands, cmd)
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
		c.options = append(c.options, opt)
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
	if len(c.args) != 0 {
		r = append(r, fmt.Sprintf("%v %v", c.Name(), c.args))
	} else {
		r = append(r, c.Name())
	}
	for _, cmd := range c.commands {
		usages := cmd.UsageString()
		for _, str := range usages {
			r = append(r, fmt.Sprintf("%v %v", c.Name(), str))
		}
	}
	if c.root || c.usage.Valid() {
		r = append(r, fmt.Sprintf("%v -h | --help", c.Name()))
	}
	if c.root || c.version.Valid() {
		r = append(r, fmt.Sprintf("%v --version", c.Name()))
	}
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

func (c Command) GetUsage() string {
	return ""
}
