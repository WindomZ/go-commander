package commander

import "errors"

type Result interface {
	Code() int
	Error() error
	Break() bool
}

var (
	ResultPass  Result = &ResultCode{}
	ResultBreak        = &ResultCode{_break: true}
)

type ResultCode struct {
	code   int
	error  error
	_break bool
}

func (e ResultCode) Code() int {
	return e.code
}

func (e ResultCode) Error() error {
	return e.error
}

func (e ResultCode) Break() bool {
	return e._break || e.error != nil
}

func NewResult(text string) Result {
	return &ResultCode{error: errors.New(text)}
}

func NewResultCode(code int, text ...string) Result {
	var err error
	if len(text) != 0 {
		err = errors.New(text[0])
	}
	return &ResultCode{
		code:  code,
		error: err,
	}
}

func NewResultError(err error, codes ...int) Result {
	var code int = 0
	if len(codes) != 0 {
		code = codes[0]
	}
	return &ResultCode{
		code:  code,
		error: err,
	}
}
