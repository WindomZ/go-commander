package commander

import "strconv"

type actor struct {
	names  []string
	action Action
}

func (a *actor) Action(action Action) {
	a.action = action
}

func (a actor) allow(d DocoptMap) (b bool) {
	for _, key := range a.names {
		if v, ok := d[key]; !ok {
		} else if b, ok = v.(bool); ok {
		} else if str, ok := v.(string); ok && len(str) != 0 {
			b, _ = strconv.ParseBool(str)
		}
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
