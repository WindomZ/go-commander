package commander

import "errors"

var (
	ErrNil     error = errors.New("error nil")
	ErrCommand       = errors.New("error command")
	ErrOption        = errors.New("error option")
)

type ErrFunc func(err error, obj interface{})
