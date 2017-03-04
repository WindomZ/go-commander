package commander

import "fmt"

type Optional interface {
	Name() string
	UsageString() string
	OptionString() string
}

type Option struct {
	Flags
	desc     string
	exec     Exec
	required bool
}

func newOption(flags, desc string) *Option {
	return &Option{
		Flags: *newFlag(flags),
		desc:  desc,
	}
}

func (o Option) OptionString() string {
	if len(o.desc) == 0 {
		return o.Flags.OptionString()
	}
	sf := o.Flags.OptionString()
	if len(sf) >= 12 {
		return fmt.Sprintf("%s  %s", sf, o.desc)
	}
	return fmt.Sprintf("%-14s%s", sf, o.desc)
}

func (o *Option) Required() *Option {
	o.required = true
	return o
}

func (o Option) IsRequired() bool {
	return o.required
}

func (o Option) IsOptional() bool {
	return !o.IsRequired()
}
