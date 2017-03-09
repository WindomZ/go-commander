package commander

// The following are ACTION functions, chose one if you like it.
type Action func(c *Context) Result // default internal function
type Action1 func(c *Context) error
type Action2 func(c *Context)
type Action3 func(m map[string]interface{}) error
