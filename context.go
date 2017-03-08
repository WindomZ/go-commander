package commander

import "fmt"

type Context struct {
	Args ContextArgs
	Doc  DocoptMap
}

func newContext(args []string, d DocoptMap) *Context {
	return &Context{
		Args: ContextArgs(args),
		Doc:  d,
	}
}

func (c Context) String() string {
	return fmt.Sprintf("%#v", c)
}
