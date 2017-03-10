package commander

// _Options
type _Options []*_Option

func (o _Options) IsEmpty() bool {
	return len(o) == 0
}

func (o _Options) UsagesString(ones ...bool) (r []string) {
	var one bool
	if len(o) == 1 && len(ones) != 0 {
		one = ones[0]
	}
	for _, opt := range o {
		r = append(r, opt.UsageString(one))
	}
	return
}

func (o _Options) OptionsString() (r []string) {
	for _, opt := range o {
		if s := opt.OptionString(); len(s) != 0 {
			r = append(r, s)
		}
	}
	return
}

func (o _Options) run(c Context) Result {
	for _, opt := range o {
		if r := opt.run(c); r != nil && r.Break() {
			return r
		}
	}
	return ResultPass
}
