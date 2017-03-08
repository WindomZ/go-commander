package commander

type actor struct {
	names  []string
	action Action
}

func (a *actor) Action(action Action) {
	a.action = action
}

func (a actor) allow(d DocoptMap) (b bool) {
	for _, key := range a.names {
		b, _ = d.GetBool(key)
	}
	return
}

func (a actor) run(d DocoptMap) Result {
	if !a.allow(d) || a.action == nil {
	} else if r := a.action(d); r != nil {
		return r
	}
	return ResultPass
}
