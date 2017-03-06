package commander

type Commands []*Command

func (c Commands) OptionsString() (r []string) {
	for _, cmd := range c {
		r = append(r, cmd.OptionsString()...)
	}
	return
}
