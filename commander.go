package commander

type Commander interface {
	Version(ver string, flags ...string) Commander
	Description(desc string) Commander
	Usage(usage string) Commander
	Command(name string) Commander
	Alias(alias string) Commander
	Option(flags, desc string) Commander
}

func NewCommander(name string) Commander {
	return newCommand(name)
}
