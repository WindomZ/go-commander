package commander

import (
	"fmt"
	"strings"
)

// _Command Command line command implementation
type _Command struct {
	actor
	usage      string              // api set usage
	root       bool                // root command
	clone      bool                // clone command for new help message line
	desc       string              // description
	annotation map[string][]string // annotation, like 'try', 'examples', etc.
	arguments  _Arguments          // parse arguments from usage
	commands   _Commands           // api set subcommands
	options    _Options            // api set options
	errFunc    errFunc             // error function // TODO: not finish this
}

func newCommand(root bool) *_Command {
	c := &_Command{
		root: root,
		errFunc: func(err error, obj interface{}) {
			fmt.Printf("  err: %v\n  object: %#v\n", err, obj)
		},
	}
	return c
}

func (c *_Command) Usage(usage string, args ...interface{}) Commander {
	c.usage = strings.TrimSpace(usage)
	c.regexpNames()
	c.regexpArguments()
	if len(args) >= 1 {
		c.desc, _ = args[0].(string)
	}
	if len(args) >= 2 {
		c.setAction(args[1])
	}
	if len(args) >= 3 {
		c.errFunc, _ = args[2].(errFunc)
	}
	return c
}

func (c *_Command) regexpNames() {
	c.names = regexpCommand(c.usage)
}

func (c *_Command) regexpArguments() {
	c.arguments.Set(c.usage)
}

func (c _Command) Valid() bool {
	return len(c.names) != 0 && len(c.usage) != 0
}

func (c _Command) Name() string {
	if len(c.names) == 0 {
		return ""
	}
	name := c.names[0]
	if len(c.names) > 1 {
		name = fmt.Sprintf("(%s)", strings.Join(c.names, "|"))
	}
	return name
}

func (c *_Command) Version(ver string) Commander {
	return c
}

func (c *_Command) Description(desc string) Commander {
	c.desc = desc
	return c
}

func (c *_Command) Annotation(title string, contents []string) Commander {
	if c.annotation == nil {
		c.annotation = make(map[string][]string)
	}
	c.annotation[title] = contents
	return c
}

func (c *_Command) Aliases(aliases []string) Commander {
	name := c.Name()
	c.names = append(c.names, aliases...)
	c.usage = replaceCommand(c.usage, name, c.Name())
	return c
}

func (c *_Command) Action(action interface{}, keys ...[]string) Commander {
	c.actor.Action(action, keys...)
	return c
}

func (c *_Command) addCommand(cmd *_Command) bool {
	if cmd.Valid() {
		for _, _cmd := range c.commands {
			_cmd.addExcludeKeys(cmd.getIncludeKeys())
		}
		c.commands = append(c.commands, cmd)
		return true
	} else if c.errFunc != nil {
		c.errFunc(errCommand, cmd)
	}
	return false
}

func (c *_Command) Command(usage string, args ...interface{}) Commander {
	if c.clone {
		usage = c.usage + " " + usage
	} else if c.Valid() {
		cmd := newCommand(false)
		cmd.Usage(usage, args...)
		cmd.addIncludeKeys(cmd.names)
		c.addCommand(cmd)
		return cmd
	}
	return c.Usage(usage, args...)
}

func (c *_Command) Option(usage string, args ...interface{}) Commander {
	if opt := newOption(usage, args...); opt.Valid() {
		c.options = append(c.options, opt)
	} else if c.errFunc != nil {
		c.errFunc(errOption, opt)
	}
	return c
}

func (c *_Command) Line(usage string, args ...interface{}) *_Command {
	cmd := newCommand(c.root)
	cmd.Usage(usage, args...)
	cmd.clone = true
	return cmd
}

func (c *_Command) LineArgument(usage string, args ...interface{}) Commander {
	usage = c.Name() + " " + usage
	cmd := c.Line(usage, args...)
	if cmd.arguments.IsEmpty() {
		return cmd
	}
	cmd.addIncludeKeys(cmd.arguments.Get())
	c.addCommand(cmd)
	return cmd
}

func (c *_Command) LineOption(usage string, args ...interface{}) Commander {
	cmd := c.Line(c.usage, args...)
	cmd.Option(usage, args...)
	if cmd.options.IsEmpty() {
		return cmd
	}
	c.addCommand(cmd)
	return cmd
}

func (c _Command) UsagesString() (r []string) {
	if !c.Valid() {
		return
	}
	str := c.usage
	if len(c.options) != 0 {
		uStrs := c.options.UsagesString(c.arguments.IsEmpty())
		for _, uStr := range uStrs {
			str += " " + uStr
		}
	}
	name := c.Name()
	if !(c.root || c.clone) || str != name {
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

func (c _Command) OptionsString() (r []string) {
	if !c.Valid() {
		return
	}
	r = append(r, c.options.OptionsString()...)
	r = append(r, c.commands.OptionsString()...)
	return
}

func (c _Command) HelpMessage() string {
	var hm _HelpMessage

	if len(c.desc) != 0 {
		hm.Description(c.desc)
	}

	if strs := c.UsagesString(); len(strs) != 0 {
		hm.Title("Usage")
		for _, str := range strs {
			hm.Subtitle(str)
		}
	}

	if strs := c.OptionsString(); len(strs) != 0 {
		hm.Line().Title("Options")
		strs = c.OptionsString()
		for _, str := range strs {
			hm.Subtitle(str)
		}
	}

	if c.annotation != nil {
		for title, contents := range c.annotation {
			hm.Line().Title(title)
			for _, content := range contents {
				hm.Subtitle(content)
			}
		}
	}

	return hm.String()
}

func (c _Command) ShowHelpMessage() string {
	s := c.HelpMessage()
	fmt.Println(s)
	return s
}

func (c _Command) Parse(args ...[]string) (Context, error) {
	return nil, nil
}

func (c _Command) run(context Context) _Result {
	if c.root || c.allow(context) {
		if r := c.commands.run(context); r != nil {
			return r
		} else if r := c.actor.run(context); r != nil {
			if r.Break() {
				return r
			}
			return c.options.run(context)
		}
	}
	return nil
}
