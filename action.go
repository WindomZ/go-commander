package commander

// The following are ACTION functions, chose one if you like it.
type (
	Action       func(c *Context) Result // default internal function
	ActionNormal func(c *Context) error
	ActionSimple func(c *Context)
	ActionNative func(m map[string]interface{}) error
)
