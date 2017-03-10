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

// parseAction handle function to Action type
func parseAction(arg interface{}) (a Action) {
	switch action := arg.(type) {
	case func(c Context) Result: // Action
		a = action
	case func(c Context) error: // ActionNormal
		a = func(c Context) Result {
			if err := action(c); err != nil {
				return NewResultError(err)
			}
			return ResultPass
		}
	case func(c Context): // ActionSimple
		a = func(c Context) Result {
			action(c)
			return ResultPass
		}
	case func(): // ActionNative
		a = func(c Context) Result {
			action()
			return ResultPass
		}
	case func() error: // ActionNativeSimple
		a = func(c Context) Result {
			action()
			return ResultPass
		}
	case func(m map[string]interface{}) error: // ActionNativeDocopt
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
