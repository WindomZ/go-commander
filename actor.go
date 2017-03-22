package commander

// actor Assigned to execute the command
type actor struct {
	names    []string // the keys contain one of names than execute action
	musts    []string // the keys contain all of musts than execute action
	excludes []string // the keys contain one of excludes than execute action no longer
	action   Action   // executed command
}

// addMustKeys append keys to actor.musts
func (a *actor) addMustKeys(keys []string) {
	if keys != nil && len(keys) != 0 {
		a.musts = append(a.musts, keys...)
	}
}

// addExcludeKeys append keys to actor.excludes
func (a *actor) addExcludeKeys(keys []string) {
	if keys != nil && len(keys) != 0 {
		for _, key := range keys {
			for _, must := range a.musts {
				if key != must {
					a.excludes = append(a.excludes, key)
				}
			}
		}
	}
}

// setAction set executive function to actor.action
// arg is ACTION function, see ./action.go
func (a *actor) setAction(arg interface{}) {
	if action := parseAction(arg); action != nil {
		a.action = action
	}
}

// Action set executive function to actor.action and actor.musts
// action is ACTION function, see ./action.go
func (a *actor) Action(action interface{}, keys ...[]string) {
	a.setAction(action)
	if len(keys) != 0 {
		a.addMustKeys(keys[0])
	}
}

// allow Determine whether meet the requirements(actor.names or actor.musts) for the execution
func (a actor) allow(c Context) (pass bool) {
	//defer func() {
	//	println(fmt.Sprintf("allow:\n 1 %#v\n 2 %v\n 3 %v\n 4 %v",
	//		a, c.String(), a.action != nil, pass))
	//}()
	for _, key := range a.excludes {
		if c.Contain(key) {
			pass = false
			return
		}
	}

	for _, key := range a.names {
		if c.Contain(key) {
			pass = true
			return
		}
	}

	for _, key := range a.musts {
		if !c.Contain(key) {
			pass = false
			return
		}
	}
	pass = len(a.musts) != 0
	return
}

// run Common external function, if allow() than execute actor.action
func (a actor) run(c Context) (reuslt _Result) {
	//defer func() {
	//	println(fmt.Sprintf("run:\n 1 %#v\n 2 %v\n 3 %v\n 4 %#v",
	//		a, c.String(), a.action != nil, reuslt))
	//}()
	if !a.allow(c) || a.action == nil {
		reuslt = nil
	} else {
		reuslt = a.action(c)
	}
	return
}
