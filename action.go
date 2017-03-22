package commander

// The following are ACTION functions, chose one if you like it.
type (
	Action             func(c Context) _Result // default internal function
	ActionNormal       func(c Context) error
	ActionSimple       func(c Context)
	ActionNative       func()
	ActionNativeSimple func() error
	ActionNativeDocopt func(m map[string]interface{}) error
)

// parseAction handle function to Action type
func parseAction(arg interface{}) (a Action) {
	switch action := arg.(type) {
	case func(c Context) _Result: // Action
		a = action
	case func(c Context) error: // ActionNormal
		a = func(c Context) _Result {
			if err := action(c); err != nil {
				return newResultError(err)
			}
			return resultPass
		}
	case func(c Context): // ActionSimple
		a = func(c Context) _Result {
			action(c)
			return resultPass
		}
	case func(): // ActionNative
		a = func(c Context) _Result {
			action()
			return resultPass
		}
	case func() error: // ActionNativeSimple
		a = func(c Context) _Result {
			if err := action(); err != nil {
				return newResultError(err)
			}
			return resultPass
		}
	case func(m map[string]interface{}) error: // ActionNativeDocopt
		a = func(c Context) _Result {
			if err := action(c.Map()); err != nil {
				return newResultError(err)
			}
			return resultPass
		}
	default:
		a = nil
	}
	return
}
