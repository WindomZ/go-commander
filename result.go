package commander

import "errors"

type _Result interface {
	Code() int
	Error() error
	Break() bool
	setBreak()
}

type _ResultCode struct {
	code   int
	error  error
	_break bool
}

func (e _ResultCode) Code() int {
	return e.code
}

func (e _ResultCode) Error() error {
	return e.error
}

func (e _ResultCode) Break() bool {
	return e._break || (e.error != nil)
}

func (e *_ResultCode) setBreak() {
	e._break = true
}

func newResult(text string) _Result {
	return &_ResultCode{error: errors.New(text)}
}

func newResultCode(code int, text ...string) _Result {
	var err error
	if len(text) != 0 {
		err = errors.New(text[0])
	}
	return &_ResultCode{
		code:  code,
		error: err,
	}
}

func newResultError(err error, codes ...int) _Result {
	var code int = 0
	if len(codes) != 0 {
		code = codes[0]
	}
	return &_ResultCode{
		code:  code,
		error: err,
	}
}

func resultPass() _Result {
	return &_ResultCode{}
}

func resultBreak() _Result {
	return &_ResultCode{_break: true}
}
