package commander

import "errors"

var (
	ErrNil    error = errors.New("error nil")
	ErrOption       = errors.New("error option")
)
