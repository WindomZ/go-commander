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
// arg like: func(c Context) Result
//           func(c Context) error
//           func(c Context)
//           func(m map[string]interface{}) error
func (a *actor) setAction(arg interface{}) {
	if action := parseAction(arg); action != nil {
		a.action = action
	}
}

// Action set executive function to actor.action and actor.musts
// action like: func(c Context) Result
//              func(c Context) error
//              func(c Context)
//              func(m map[string]interface{}) error
func (a *actor) Action(action interface{}, keys ...[]string) {
	a.setAction(action)
	if len(keys) != 0 {
		a.addMustKeys(keys[0])
	}
}

// allow Determine whether meet the requirements(actor.names or actor.musts) for the execution
func (a actor) allow(c Context) (b bool) {
	for _, key := range a.names {
		if b = c.GetBool(key); b {
			return
		}
	}
	for _, key := range a.musts {
		if b = c.GetBool(key); b {
		} else if b = c.Contain(key); b &&
			containArgument(key) {
		} else {
			b = false
			break
		}
	}
	return
}

// run Common external function, if allow() than execute actor.action
func (a actor) run(c Context) Result {
	if !a.allow(c) || a.action == nil {
	} else if r := a.action(c); r != nil {
		return r
	}
	return ResultPass
}
