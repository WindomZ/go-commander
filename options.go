package commander

type Options []*Option

func (o Options) UsagesString() (r []string) {
	for _, opt := range o {
		r = append(r, opt.UsageString())
	}
	return
}

func (o Options) OptionsString() (r []string) {
	for _, opt := range o {
		if s := opt.OptionString(); len(s) != 0 {
			r = append(r, s)
		}
	}
	return
}

func (o Options) run(d DocoptMap) Result {
	for _, opt := range o {
		if r := opt.run(d); r != nil && r.Break() {
			return r
		}
	}
	return ResultPass
}
