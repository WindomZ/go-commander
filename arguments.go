package commander

type Arguments []*Argument

func (a Arguments) IsEmpty() bool {
	return len(a) == 0
}

func (a *Arguments) Set(usage string) {
	*a = (*a)[:0]
	if strs := RegexpArgument(usage); len(strs) != 0 {
		for _, str := range strs {
			*a = append(*a, newArgument(str))
		}
	}
}

func (a Arguments) Get() (r []string) {
	for _, arg := range a {
		r = append(r, arg.name)
	}
	return
}

func (a Arguments) UsagesString() (r []string) {
	for _, arg := range a {
		r = append(r, arg.UsageString())
	}
	return
}

func (a Arguments) OptionsString() (r []string) {
	for _, arg := range a {
		if s := arg.OptionString(); len(s) != 0 {
			r = append(r, s)
		}
	}
	return
}
