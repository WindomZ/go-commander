package commander

type Context struct {
	Doc DocoptMap
}

func newContext(d DocoptMap) *Context {
	return &Context{
		Doc: d,
	}
}
