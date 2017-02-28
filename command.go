package commander

import "fmt"

type Command struct {
	name     string
	alias    string
	desc     string
	version  Version
	usage    Usage
	commands []Command
	options  []Option
	exec     Exec
	errFunc  ErrFunc
}

func newCommand(name string) *Command {
	return &Command{
		name: name,
		desc: "",
		errFunc: func(err error, obj interface{}) {
			fmt.Printf("  err: %v\n  object: %#v\n", err, obj)
		},
	}
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
	c.alias = alias
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
