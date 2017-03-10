package commander

// The following are ACTION functions, chose one if you like it.
type (
	Action             func(c Context) Result // default internal function
	ActionNormal       func(c Context) error
	ActionSimple       func(c Context)
	ActionNative       func()
	ActionNativeSimple func() error
	ActionNativeDocopt func(m map[string]interface{}) error
)

func parseAction(arg interface{}) (a Action) {
	switch action := arg.(type) {
	case Action:
		a = action
	case ActionNormal:
		a = func(c Context) Result {
			if err := action(c); err != nil {
				return NewResultError(err)
			}
			return ResultPass
		}
	case ActionSimple:
		a = func(c Context) Result {
			action(c)
			return ResultPass
		}
	case ActionNative:
		a = func(c Context) Result {
			action()
			return ResultPass
		}
	case ActionNativeSimple:
		a = func(c Context) Result {
			action()
			return ResultPass
		}
	case ActionNativeDocopt:
		a = func(c Context) Result {
			if err := action(c.Map()); err != nil {
				return NewResultError(err)
			}
			return ResultPass
		}
	default:
		a = nil
	}
	return
}
