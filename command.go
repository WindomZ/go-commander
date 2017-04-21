package commander

import (
	"fmt"
	"os"
	"path"
	"regexp"
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
	last       interface{}         // the last defined object
	doc        string              // define help message
}

func newCommand(root bool) *_Command {
	return &_Command{
		root: root,
	}
}

func (c *_Command) init() *_Command {
	if !c.Valid() {
		if dir, err := os.Getwd(); err == nil {
			c.Usage(path.Base(dir))
		}
	}
	return c
}

func (c _Command) isRoot() bool {
	return c.root && !c.clone
}

func (c *_Command) Usage(usage string, args ...interface{}) Commander {
	if len(usage) != 0 {
		c.usage = strings.TrimSpace(usage)
		c.regexpNames()
		c.regexpArguments()
	}
	if len(args) >= 1 {
		c.desc, _ = args[0].(string)
	}
	if len(args) >= 2 {
		c.setAction(args[1])
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

func (c *_Command) Doc(doc string) Commander {
	c.doc = doc
	return c.init()
}

func (c *_Command) Version(ver string) Commander {
	return c.init()
}

func (c _Command) ShowVersion() string {
	return ""
}

func (c *_Command) Description(desc string) Commander {
	if c.init().last != nil {
		switch obj := c.last.(type) {
		//case *_Command:
		case *_Option:
			obj.Description(desc)
			return c
		}
	}
	c.desc = desc
	return c
}

func (c *_Command) Annotation(title string, contents []string) Commander {
	if c.annotation == nil {
		c.annotation = make(map[string][]string)
	}
	c.annotation[title] = contents
	return c.init()
}

func (c *_Command) addCommand(cmd *_Command) bool {
	if c.init().Valid() && cmd.Valid() {
		for _, _cmd := range c.commands {
			_cmd.addExcludeKeys(cmd.getIncludeKeys())
		}
		c.commands = append(c.commands, cmd)
		return true
	} else {
		panicError("ERROR Command format:", cmd)
	}
	return false
}

func (c *_Command) Command(usage string, args ...interface{}) (commander Commander) {
	if param := firstParameter(usage); isArgument(param) {
		commander = c.LineArgument(usage, args...)
		goto SetLast
	} else if isOption(param) {
		commander = c.LineOption(usage, args...)
		return
	} else if c.clone {
		usage = c.usage + " " + usage
	} else if c.Valid() {
		cmd := newCommand(false)
		cmd.Usage(usage, args...)
		cmd.addIncludeKeys(cmd.names)
		c.addCommand(cmd)
		commander = cmd
		goto SetLast
	}
	commander = c.Usage(usage, args...)
SetLast:
	c.last = commander
	return
}

func (c *_Command) Aliases(aliases []string) Commander {
	if c.init().last != nil {
		switch obj := c.last.(type) {
		//case *_Command:
		case *_Option:
			obj.Aliases(aliases)
			return c
		}
	}
	name := c.Name()
	c.names = append(c.names, aliases...)
	c.usage = replaceCommand(c.usage, name, c.Name())
	return c
}

func (c *_Command) addOption(line bool, usage string, args ...interface{}) (opt *_Option) {
	opt = newOption(usage, args...)
	opt.line = line
	opt.break_off = line
	if opt.Valid() {
		c.init().options = append(c.options, opt)
	}
	c.last = opt
	return opt
}

func (c *_Command) Option(usage string, args ...interface{}) Commander {
	if opt := c.addOption(false, usage, args...); !opt.Valid() {
		panicError("ERROR Option format:", opt)
	}
	return c
}

func (c *_Command) Line(usage string, args ...interface{}) *_Command {
	cmd := newCommand(c.root)
	cmd.Usage(usage, args...)
	cmd.clone = true
	cmd.ignore = true
	return cmd
}

func (c *_Command) LineArgument(usage string, args ...interface{}) Commander {
	usage = c.Name() + " " + usage
	cmd := c.Line(usage, args...)
	if cmd.arguments.IsEmpty() {
		return cmd
	}
	cmd.ignore = false
	cmd.addIncludeKeys(cmd.arguments.Get())
	c.addCommand(cmd)
	return cmd
}

func (c *_Command) LineOption(usage string, args ...interface{}) Commander {
	cmd := c.Line(c.usage, args...)
	opt := cmd.addOption(true, usage, args...)
	if cmd.options.IsEmpty() {
		return cmd
	}
	cmd.addIncludeKeys(opt.Names())
	c.addCommand(cmd)
	return cmd
}

func (c *_Command) Action(action interface{}, keys ...[]string) Commander {
	if c.init().last != nil {
		switch obj := c.last.(type) {
		//case *_Command:
		case *_Option:
			obj.actor.Action(action, keys...)
			if c.hasAction() {
				return c
			}
		}
	}
	c.actor.Action(action, keys...)
	return c
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
	return
}

func (c _Command) OptionsString() (r map[string]string) {
	if !c.Valid() {
		return
	}
	r = c.options.OptionsString()
	opts := c.commands.OptionsString()
	for k, v := range opts {
		r[k] = v
	}
	return
}

func (c _Command) CommandsString(prefix string) (r []string) {
	if !c.Valid() {
		return
	}
	name := regexp.MustCompile(`[()]`).ReplaceAllString(c.Name(), "")
	if c.root {
		name = prefix
	} else {
		if len(prefix) != 0 {
			name = prefix + " " + name
		}
		if len(c.desc) != 0 {
			r = append(r, FormatDescription(name, c.desc))
		}
	}
	r = append(r, c.commands.CommandsString(name)...)
	return
}

// HelpMessage get string of help message that generated according to the docopt format
func (c _Command) HelpMessage() string {
	if len(c.doc) != 0 {
		return c.doc
	}

	var hm _HelpMessage

	// Description
	if len(c.desc) != 0 {
		hm.Description(c.desc)
	}

	// Usages
	if strs := c.UsagesString(); len(strs) != 0 {
		hm.Title("Usage")
		for _, str := range strs {
			hm.Subtitle(str)
		}
	}

	// Options
	if opts := c.OptionsString(); len(opts) != 0 {
		strs := sortStringMap(opts)
		hm.Line().Title("Options")
		for _, str := range strs {
			hm.Subtitle(str)
		}
	}

	// Commands
	if strs := c.CommandsString(""); len(strs) != 0 {
		hm.Line().Title("Commands")
		for _, str := range strs {
			hm.Subtitle(str)
		}
	}

	// Annotation
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
		}
		if r := c.options.run(context); r != nil && r.Break() {
			return r
		}
		return c.actor.run(context, c.isRoot())
	}
	return nil
}

func (c *_Command) ErrorHandling(f func(error)) Commander {
	return c
}
