package commander

type actor struct {
	names  []string
	action Action
}

func (a *actor) Action(action Action) {
	a.action = action
}

func (a actor) allow(c *Context) (b bool) {
	for _, key := range a.names {
		b, _ = c.Doc.GetBool(key)
	}
	return
}

func (a actor) run(c *Context) Result {
	if !a.allow(c) || a.action == nil {
	} else if r := a.action(c); r != nil {
		return r
	}
	return ResultPass
}
