package commander

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

type Command struct {
	actor
	usage      string
	root       bool
	clone      bool
	version    string
	desc       string
	annotation map[string][]string
	arguments  Arguments
	commands   Commands
	options    Options
	errFunc    ErrFunc
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
		c.action, _ = args[1].(Action)
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

func (c *Command) Clone() *Command {
	c.clone = true
	return c
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

func (c *Command) Version(ver string) Commander {
	c.version = ver
	return c
}

func (c *Command) Description(desc string) Commander {
	c.desc = desc
	return c
}

func (c *Command) Annotation(title string, contents []string) Commander {
	if c.annotation == nil {
		c.annotation = make(map[string][]string)
	}
	c.annotation[title] = contents
	return c
}

func (c *Command) Action(action Action) Commander {
	c.actor.Action(action)
	return c
}

func (c *Command) Command(usage string, args ...interface{}) Commander {
	if c.clone {
		c.usage += " " + usage
		c.regexpNames()
		c.regexpArguments()
		return c
	}
	cmd := newCommand(usage, false, args...)
	if cmd.Valid() {
		c.commands = append(c.commands, cmd)
	} else if c.errFunc != nil {
		c.errFunc(ErrCommand, cmd)
	}
	return cmd
}

func (c *Command) Option(usage string, args ...interface{}) Commander {
	if opt := newOption(usage, args...); opt.Valid() {
		c.options = append(c.options, opt)
	} else if c.errFunc != nil {
		c.errFunc(ErrOption, opt)
	}
	return c
}

func (c *Command) line(usage string, args ...interface{}) *Command {
	return newCommand(usage, c.root, args...).Clone()
}

func (c *Command) LineArgument(usage string, args ...interface{}) Commander {
	usage = c.Name() + " " + usage
	cmd := c.line(usage, args...)
	if cmd.arguments.IsEmpty() {
		return cmd
	}
	c.commands = append(c.commands, cmd)
	return cmd
}

func (c *Command) LineOption(usage string, args ...interface{}) Commander {
	cmd := c.line(c.usage, args...)
	cmd.Option(usage, args...)
	if cmd.options.IsEmpty() {
		return cmd
	}
	c.commands = append(c.commands, cmd)
	return cmd
}

func (c Command) UsagesString() (r []string) {
	if !c.Valid() {
		return
	}
	str := c.usage
	if len(c.options) != 0 {
		uStrs := c.options.UsagesString(c.clone && c.arguments.IsEmpty())
		for _, uStr := range uStrs {
			str += " " + uStr
		}
	}
	name := c.Name()
	if str != name {
		r = append(r, str)
	}
	name += " "
	for _, cmd := range c.commands {
		uStrs := cmd.UsagesString()
		for _, uStr := range uStrs {
			if strings.HasPrefix(uStr, name) {
				r = append(r, uStr)
			} else {
				r = append(r, name+uStr)
			}
		}
	}
	if c.root && !c.clone {
		r = append(r, fmt.Sprintf("%s-h | --help", name))
		r = append(r, fmt.Sprintf("%s--version", name))
	}
	return
}

func (c Command) OptionsString() (r []string) {
	if !c.Valid() {
		return
	}
	r = append(r, c.options.OptionsString()...)
	r = append(r, c.commands.OptionsString()...)
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

	if c.annotation != nil {
		for title, contents := range c.annotation {
			bb.WriteString(fmt.Sprintf("\n%s:\n", title))
			for _, content := range contents {
				bb.WriteString(fmt.Sprintf("  %s\n", content))
			}
		}
	}

	return bb.String()
}

func (c Command) ShowHelpMessage() string {
	s := c.GetHelpMessage()
	fmt.Println(s)
	return s
}

func (c Command) Parse(args ...[]string) (*Context, error) {
	var argv []string = nil
	if len(args) != 0 {
		argv = args[0]
	}
	d, err := Parse(c.GetHelpMessage(), argv, true, c.version, false)
	if err != nil {
		return nil, err
	}
	context := newContext(d)
	if r := c.run(context); r != nil {
		if r.Break() {
			return context, r.Error()
		}
	} else {
		// TODO: Should be print help message, but docopt auto to do this.
	}
	return context, nil
}

func (c Command) run(context *Context) Result {
	if c.root || c.allow(context) {
		if r := c.options.run(context); r != nil && r.Break() {
			return r
		}
		if r := c.actor.run(context); r != nil && r.Break() {
			return r
		}
		return c.commands.run(context)
	}
	return nil
}
