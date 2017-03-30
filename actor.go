package commander

import "fmt"

// actor Assigned to execute the command
type actor struct {
	names     []string        // the keys contain one of names than execute action
	triggers  map[string]bool // the keys contain all true values and none false value in triggers than execute action
	action    Action          // executed command
	ignore    bool            // ignore this action
	break_off bool            // break off both actions
}

// addIncludeKeys append include keys to actor.triggers
func (a *actor) addIncludeKeys(keys []string) {
	if keys != nil && len(keys) != 0 {
		if a.triggers == nil {
			a.triggers = make(map[string]bool)
		}
		for _, key := range keys {
			a.triggers[key] = true
		}
	}
}

// getIncludeKeys get list of include keys
func (a actor) getIncludeKeys() (keys []string) {
	for key, ok := range a.triggers {
		if ok {
			keys = append(keys, key)
		}
	}
	return keys
}

// addExcludeKeys append exclude keys to actor.triggers
func (a *actor) addExcludeKeys(keys []string) {
	if keys != nil && len(keys) != 0 {
		if a.triggers == nil {
			a.triggers = make(map[string]bool)
		}
		for _, key := range keys {
			if _, ok := a.triggers[key]; !ok {
				a.triggers[key] = false
			}
		}
	}
}

// getExcludeKeys get list of exclude keys
func (a actor) getExcludeKeys() (keys []string) {
	for key, ok := range a.triggers {
		if !ok {
			keys = append(keys, key)
		}
	}
	return keys
}

// setAction set executive function to actor.action
// arg is ACTION function, see ./action.go
func (a *actor) setAction(arg interface{}) {
	if action := parseAction(arg); action != nil {
		a.action = action
	}
}

// hasAction there is a legitimate actor.action
func (a actor) hasAction() bool {
	return a.action != nil
}

// Action set executive function to actor.action and include keys to actor.triggers
// action is ACTION function, see ./action.go
func (a *actor) Action(action interface{}, keys ...[]string) {
	a.setAction(action)
	if len(keys) != 0 {
		a.addIncludeKeys(keys[0])
	}
}

// allow Determine whether meet the requirements(actor.names or actor.triggers) for the execution
func (a actor) allow(c Context) (pass bool) {
	if DEBUG {
		defer func() {
			fmt.Printf("----------allow----------"+
				"\n  1.actor  %#v\n  2.argv   %v\n  3.action %v\n  4.pass   %v\n",
				a, c.Map(), a.action != nil, pass)
		}()
	}
	for key, ok := range a.triggers {
		if !ok && c.Contain(key) {
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

	for key, ok := range a.triggers {
		if ok && !c.Contain(key) {
			pass = false
			return
		}
	}
	pass = len(a.triggers) != 0
	return
}

// run Common external function, if allow() than execute actor.action
func (a actor) run(c Context, force ...bool) (result _Result) {
	if DEBUG {
		defer func() {
			fmt.Printf("----------run----------"+
				"\n  1.actor  %#v\n  2.argv   %v\n  3.action %v\n  4.result %#v\n",
				a, c.Map(), a.action != nil, result)
		}()
	}
	if a.action == nil {
		return
	} else if len(force) != 0 && force[0] {
	} else if !a.allow(c) {
		return
	} else if a.ignore {
		return resultPass()
	}
	result = a.action(c)
	if a.break_off && result != nil && !result.Break() {
		result.setBreak()
	}
	return
}
