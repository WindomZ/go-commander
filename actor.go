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
// arg like: func(c Context) _Result
//           func(c Context) error
//           func(c Context)
//           func(m map[string]interface{}) error
func (a *actor) setAction(arg interface{}) {
	if action := parseAction(arg); action != nil {
		a.action = action
	}
}

// Action set executive function to actor.action and actor.musts
// action like: func(c Context) _Result
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
func (a actor) allow(c Context) bool {
	//println(fmt.Sprintf("allow:\n1 %#v\n2 %v\n3 %v",
	//	a, c.String(), a.action != nil))
	for _, key := range a.names {
		if c.Contain(key) {
			return true
		}
	}
	for _, key := range a.musts {
		if !c.Contain(key) {
			return false
		}
	}
	return len(a.musts) != 0
}

// run Common external function, if allow() than execute actor.action
func (a actor) run(c Context) _Result {
	//println(fmt.Sprintf("run:\n1 %#v\n2 %v\n3 %v",
	//	a, c.String(), a.action != nil))
	if !a.allow(c) || a.action == nil {
		//println(fmt.Sprintf("run:\n4 %v", false))
	} else if r := a.action(c); r != nil {
		return r
	}
	return resultPass
}
