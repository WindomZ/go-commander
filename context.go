package commander

type Context interface {
	// _ContextArguments
	GetArg(index int) string
	ArgsString() string
	ArgsStringSeparator(sep string, offsets ...int) string
	// DocoptMap
	Map() map[string]interface{}
	Get(key string) interface{}
	Contain(key string) bool
	GetString(key string) string
	GetStrings(key string) []string
	GetBool(key string) bool
	GetInt64(key string) (int64, bool)
	GetInt(key string) (int, bool)
	GetFloat64(key string) (float64, bool)
	GetFloat(key string) (float32, bool)
	String() string
}

// _Context
type _Context struct {
	_ContextArguments `json:"arguments"`
	DocoptMap         `json:"docopt"`
}

func newContext(args []string, d DocoptMap) *_Context {
	return &_Context{
		_ContextArguments: newContextArguments(args),
		DocoptMap:         d,
	}
}

func (c _Context) Contain(key string) bool {
	ok := c.DocoptMap.Contain(key)
	return ok
}

func (c _Context) String() string {
	return c.DocoptMap.String()
}
