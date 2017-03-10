package commander

// actor Assigned to execute the command
type actor struct {
	names  []string // if true keys contain one of names than execute action
	musts  []string // if true keys contain all of musts than execute action
	action Action   // execute the command
}

// addMustKeys append keys to actor.musts
func (a *actor) addMustKeys(keys []string) {
	if keys != nil && len(keys) != 0 {
		a.musts = append(a.musts, keys...)
	}
}

// setAction set executive function to actor.action
// arg like: func(c *Context) Result
//           func(c *Context) error
//           func(c *Context)
//           func(m map[string]interface{}) error
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

// Action set executive function to actor.action and actor.musts
// action like: func(c *Context) Result
//              func(c *Context) error
//              func(c *Context)
//              func(m map[string]interface{}) error
func (a *actor) Action(action interface{}, keys ...[]string) {
	a.setAction(action)
	if len(keys) != 0 {
		a.addMustKeys(keys[0])
	}
}

// allow Determine whether meet the requirements(actor.names or actor.musts) for the execution
func (a actor) allow(c *Context) (b bool) {
	for _, key := range a.names {
		if b, _ = c.Doc.GetBool(key); b {
			return
		}
	}
	for _, key := range a.musts {
		if b = c.Doc.GetMustBool(key); b {
		} else if b = c.Doc.Contain(key); b &&
			containArgument(key) {
		} else {
			b = false
			break
		}
	}
	return
}

// run Common external function, if allow() than execute actor.action
func (a actor) run(c *Context) Result {
	if !a.allow(c) || a.action == nil {
	} else if r := a.action(c); r != nil {
		return r
	}
	return ResultPass
}
