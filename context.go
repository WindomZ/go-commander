package commander

// Context
type Context struct {
	Args ContextArgs `json:"arguments"`
	Doc  DocoptMap   `json:"docopt"`
}

func newContext(args []string, d DocoptMap) *Context {
	return &Context{
		Args: newContextArgs(args),
		Doc:  d,
	}
}

func (c Context) String() string {
	return c.Doc.String()
}

func (c Context) Contain(key string) bool {
	return c.Doc.Contain(key)
}
