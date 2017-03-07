package commander

type Commands []*Command

func (c Commands) OptionsString() (r []string) {
	for _, cmd := range c {
		r = append(r, cmd.OptionsString()...)
	}
	return
}

func (c Commands) run(d DocoptMap) Result {
	for _, cmd := range c {
		if r := cmd.run(d); r != nil && r.Break() {
			return r
		}
	}
	return ResultPass
}
