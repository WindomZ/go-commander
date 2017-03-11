package commander

import "errors"

var (
	errNil     error = errors.New("error nil")
	errCommand       = errors.New("error command")
	errOption        = errors.New("error option")
)

type errFunc func(err error, obj interface{})
