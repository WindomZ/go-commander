package commander

import "fmt"

type Option struct {
	Flags
	desc     string
	exec     ExecFunc
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
		return ""
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

func (o Option) IsOptionRequired() bool {
	return o.required
}

func (o Option) IsOptionOptional() bool {
	return !o.IsOptionRequired()
}
