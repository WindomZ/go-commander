package commander

import (
	"fmt"
	"regexp"
	"strings"
)

type Command struct {
	cmd      string
	name     string
	args     string
	alias    []string
	desc     string
	version  Version
	usage    Usage
	commands Commands
	options  Options
	exec     Exec
	errFunc  ErrFunc
	root     bool
}

func newCommand(cmd string, root bool) *Command {
	c := &Command{
		cmd:  strings.TrimSpace(cmd),
		desc: "",
		errFunc: func(err error, obj interface{}) {
			fmt.Printf("  err: %v\n  object: %#v\n", err, obj)
		},
		root: root,
	}
	return c.initCommand()
}

func (c *Command) initCommand() *Command {
	if strs := regexp.MustCompile(`^[A-Za-z0-9_-]+`).
		FindAllString(c.cmd, -1); len(strs) != 0 {
		c.name = strs[0]
	}
	if strs := regexp.MustCompile(`(?i:<|\[)[A-Za-z0-9_\[\]<>-]+(?i:>|])`).
		FindAllString(c.cmd, -1); len(strs) != 0 {
		c.args = strings.Join(strs, " ")
	}
	return c
}

func (c Command) Name() string {
	if alias := strings.Join(c.alias, "|"); len(alias) != 0 {
		return fmt.Sprintf("(%s|%s)", c.name, alias)
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
	if strs := regexp.MustCompile(`^[A-Za-z0-9_-]+`).
		FindAllString(strings.TrimSpace(alias), -1); len(strs) != 0 {
		c.alias = append(c.alias, strs[0])
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

func (c Command) UsagesString() (r []string) {
	if !c.Valid() {
		return
	}
	str := c.Name()
	if len(c.args) != 0 {
		str = fmt.Sprintf("%s %s", str, c.args)
	}
	if len(c.options) != 0 {
		// TODO: 未完成
	}
	r = append(r, str)
	for _, cmd := range c.commands {
		usages := cmd.UsagesString()
		for _, str := range usages {
			r = append(r, fmt.Sprintf("%s %s", c.Name(), str))
		}
	}
	if c.root || c.usage.Valid() {
		r = append(r, fmt.Sprintf("%s -h | --help", c.Name()))
	}
	if c.root || c.version.Valid() {
		r = append(r, fmt.Sprintf("%s --version", c.Name()))
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
