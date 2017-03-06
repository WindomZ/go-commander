package commander

import (
	"bytes"
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

func (c *Command) Command(usage string, args ...interface{}) Commander {
	cmd := newCommand(usage, false)
	if len(args) >= 1 {
		// TODO: action
	}
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

func (c *Command) Option(flags string, args ...interface{}) Commander {
	var desc string
	if len(args) >= 1 {
		desc, _ = args[0].(string)
	}
	if len(args) >= 2 {
		// TODO: action
	}
	if len(args) >= 3 {
		descDef := make([]string, 0, len(args)-2)
		for _, d := range args[2:] {
			descDef = append(descDef, fmt.Sprintf("%v", d))
		}
		desc += fmt.Sprintf(" [default: %s]", strings.Join(descDef, ","))
	}
	if opt := newOption(strings.TrimSpace(flags), strings.TrimSpace(desc)); opt.Valid() {
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
	if c.root || c.usage.Valid() {
		r = append(r, fmt.Sprintf("%s -h | --help", c.Name()))
	}
	if c.root && c.version.Valid() {
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

func (c Command) GetUsage() string {
	if c.usage.Valid() {
		return c.usage.Get()
	}
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

func (c Command) Parse() (map[string]interface{}, error) {
	return Parse(c.GetUsage(), nil, true, c.version.Get(), false)
}
