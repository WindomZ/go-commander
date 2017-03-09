package commander

type actor struct {
	names  []string
	action Action
	keys   []string
}

func (a *actor) setAction(arg interface{}) {
	if action, ok := arg.(Action); ok {
		a.action = action
	} else if action, ok := arg.(func(c *Context) Result); ok {
		a.action = action
	} else if action, ok := arg.(func(c *Context) error); ok {
		a.action = func(c *Context) Result {
			if err := action(c); err != nil {
				return NewResultError(err)
			}
			return ResultPass
		}
	} else if action, ok := arg.(func(c *Context)); ok {
		a.action = func(c *Context) Result {
			action(c)
			return ResultPass
		}
	} else if action, ok := arg.(func(m map[string]interface{}) error); ok {
		a.action = func(c *Context) Result {
			if err := action(c.Doc.Map()); err != nil {
				return NewResultError(err)
			}
			return ResultPass
		}
	}
}

func (a *actor) addKeys(keys []string) {
	if keys != nil && len(keys) != 0 {
		a.keys = append(a.keys, keys...)
	}
}

func (a *actor) Action(action interface{}, keys ...[]string) {
	a.setAction(action)
	if len(keys) != 0 {
		a.addKeys(keys[0])
	}
}

func (a actor) allow(c *Context) (b bool) {
	for _, key := range a.names {
		if b, _ = c.Doc.GetBool(key); b {
			return
		}
	}
	for _, key := range a.keys {
		if b, _ = c.Doc.GetBool(key); b {
		} else if b = c.Doc.Contain(key); b {
		} else {
			break
		}
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
