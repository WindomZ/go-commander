package commander

import "fmt"

type Optional interface {
	Name() string
	OptionString() string
}

type Option struct {
	Flags
	desc string
	exec Exec
}

func newOption(flags, desc string) *Option {
	return &Option{
		Flags: *newFlag(flags),
		desc:  desc,
	}
}

func (o Option) OptionString() string {
	if len(o.desc) == 0 {
		return o.Flags.String()
	}
	sf := o.Flags.String()
	if len(sf) >= 12 {
		return fmt.Sprintf("%s  %s", sf, o.desc)
	}
	return fmt.Sprintf("%-14s%s", sf, o.desc)
}
