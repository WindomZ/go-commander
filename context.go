package commander

type Context interface {
	// _Argv
	GetArg(index int) string
	GetArgs(offsets ...int) []string
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
	_Argv     `json:"argv"`
	DocoptMap `json:"docopt"`
}

func newContext(args []string, d DocoptMap) *_Context {
	return &_Context{
		_Argv:     newArgv(args),
		DocoptMap: d,
	}
}

func (c _Context) Contain(key string) bool {
	return c.DocoptMap.Contain(key)
}

func (c _Context) String() string {
	return c.DocoptMap.String()
}
