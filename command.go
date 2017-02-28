package commander

type Command struct {
	name     string
	alias    string
	desc     string
	version  Version
	usage    Usage
	Commands []Command
	Options  []Option
	exec     Exec
}

func newCommand(name string) *Command {
	return &Command{
		name: name,
		desc: "",
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
	c.Commands = append(c.Commands, *cmd)
	return cmd
}

func (c *Command) Alias(alias string) Commander {
	c.alias = alias
	return c
}

func (c *Command) Option(flags, desc string) Commander {
	opt := newOption(flags, desc)
	c.Options = append(c.Options, *opt)
	return c
}
