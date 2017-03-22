package commander

// _Commands Command line commands implementation
type _Commands []*_Command

// OptionsString
func (c _Commands) OptionsString() (r []string) {
	for _, cmd := range c {
		r = append(r, cmd.OptionsString()...)
	}
	return
}

func (c _Commands) run(context Context) _Result {
	for _, cmd := range c {
		if r := cmd.run(context); r != nil {
			return r
		}
	}
	return nil
}
