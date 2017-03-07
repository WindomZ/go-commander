package commander

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

type ICommand interface {
	Version(ver string) ICommand
	Description(desc string) ICommand
	Command(usage string, args ...interface{}) ICommand
	Option(flags string, args ...interface{}) ICommand
	UsagesString() []string
	OptionsString() []string
	GetHelpMessage() string
	Parse() (DocoptMap, error)
}

type Command struct {
	usage     string
	names     []string
	root      bool
	version   string
	desc      string
	arguments Arguments
	commands  Commands
	options   Options
	execFunc  ExecFunc
	errFunc   ErrFunc
}

func newCommand(usage string, root bool, args ...interface{}) *Command {
	c := &Command{
		usage: strings.TrimSpace(usage),
		root:  root,
		errFunc: func(err error, obj interface{}) {
			fmt.Printf("  err: %v\n  object: %#v\n", err, obj)
		},
	}
	c.regexpNames()
	c.regexpArguments()
	if len(args) >= 1 {
		c.desc, _ = args[0].(string)
	}
	if len(args) >= 2 {
		c.execFunc, _ = args[1].(ExecFunc)
	}
	if len(args) >= 3 {
		c.errFunc, _ = args[2].(ErrFunc)
	}
	return c
}

func (c *Command) regexpNames() {
	c.names = regexp.MustCompile(`[A-Za-z0-9_-]+`).FindAllString(
		regexp.MustCompile(`^[A-Za-z0-9_|\(\)\s-]+`).FindString(c.usage), -1)
}

func (c *Command) regexpArguments() {
	c.arguments.Set(c.usage)
}

func (c Command) Valid() bool {
	return len(c.names) != 0 && len(c.usage) != 0
}

func (c Command) Names() []string {
	return c.names
}

func (c Command) Name() string {
	name := c.names[0]
	if len(c.names) > 1 {
		name = fmt.Sprintf("(%s)", strings.Join(c.names, "|"))
	}
	return name
}

func (c *Command) Version(ver string) ICommand {
	c.version = ver
	return c
}

func (c *Command) Description(desc string) ICommand {
	c.desc = desc
	return c
}

func (c *Command) Command(usage string, args ...interface{}) ICommand {
	cmd := newCommand(usage, false, args...)
	if cmd.Valid() {
		c.commands = append(c.commands, cmd)
	} else if c.errFunc != nil {
		c.errFunc(ErrCommand, cmd)
	}
	return cmd
}

func (c *Command) Option(usage string, args ...interface{}) ICommand {
	if opt := newOption(usage, args...); opt.Valid() {
		c.options = append(c.options, opt)
	} else if c.errFunc != nil {
		c.errFunc(ErrOption, opt)
	}
	return c
}

func (c Command) UsagesString() (r []string) {
	if !c.Valid() {
		return
	}
	str := c.Name()
	if len(c.arguments) != 0 {
		uStrs := c.arguments.UsagesString()
		for _, uStr := range uStrs {
			str = fmt.Sprintf("%s %s", str, uStr)
		}
	}
	if len(c.options) != 0 {
		uStrs := c.options.UsagesString()
		for _, uStr := range uStrs {
			str = fmt.Sprintf("%s %s", str, uStr)
		}
	}
	r = append(r, str)
	for _, cmd := range c.commands {
		uStrs := cmd.UsagesString()
		for _, uStr := range uStrs {
			r = append(r, fmt.Sprintf("%s %s", c.Name(), uStr))
		}
	}
	if c.root {
		r = append(r, fmt.Sprintf("%s -h | --help", c.Name()))
		r = append(r, fmt.Sprintf("%s --version", c.Name()))
	}
	return
}

func (c Command) OptionsString() (r []string) {
	if !c.Valid() {
		return
	}
	r = append(r, c.options.OptionsString()...)
	//r = append(r, c.commands.OptionsString()...)
	return
}

func (c Command) GetHelpMessage() string {
	var bb bytes.Buffer

	if len(c.desc) != 0 {
		bb.WriteString(c.desc + "\n\n")
	}

	if strs := c.UsagesString(); len(strs) != 0 {
		bb.WriteString("Usage:\n")
		for _, str := range strs {
			bb.WriteString(fmt.Sprintf("  %s\n", str))
		}
	}

	if strs := c.OptionsString(); len(strs) != 0 {
		bb.WriteString("\nOptions:\n")
		strs = c.OptionsString()
		for _, str := range strs {
			bb.WriteString(fmt.Sprintf("  %s\n", str))
		}
	}

	return bb.String()
}

func (c Command) Parse() (DocoptMap, error) {
	return Parse(c.GetHelpMessage(), nil, true, c.version, false)
}
