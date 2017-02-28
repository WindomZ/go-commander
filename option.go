package commander

type Optional interface {
	Name() string
	IsRequired() bool
	IsOptional() bool
	IsValid() bool
}

type Option struct {
	Flags
	desc string
	exec Exec
}

func newOption(flags, desc string) *Option {
	return &Option{
		desc:  desc,
		Flags: *newFlag(flags),
	}
}
