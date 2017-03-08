package commander

import "fmt"

type Context struct {
	Args ContextArgs
	Doc  DocoptMap
}

func newContext(args []string, d DocoptMap) *Context {
	return &Context{
		Args: newContextArgs(args),
		Doc:  d,
	}
}

func (c Context) String() string {
	return fmt.Sprintf("%#v", c)
}

func (c Context) Contain(key string) bool {
	return c.Doc.Contain(key)
}
