package commander

type Action func(c *Context) Result

type Action1 func(c *Context) error
type Action2 func(c *Context)
type Action3 func(m map[string]interface{}) error
