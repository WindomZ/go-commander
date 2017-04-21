package commander

// _Commands Command line commands implementation
type _Commands []*_Command

// OptionsString
func (c _Commands) OptionsString() (r map[string]string) {
	r = make(map[string]string, len(c))
	for _, cmd := range c {
		opts := cmd.OptionsString()
		for k, v := range opts {
			r[k] = v
		}
	}
	return
}

// CommandsString
func (c _Commands) CommandsString(prefix string) (r []string) {
	for _, cmd := range c {
		r = append(r, cmd.CommandsString(prefix)...)
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
