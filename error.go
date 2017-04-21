package commander

import (
	"errors"
	"fmt"
)

func newError(a ...interface{}) error {
	return errors.New("go-commander: " + fmt.Sprint(a...))
}

func panicError(a ...interface{}) {
	panic(errors.New("\ngo-commander:\n  ERROR >>> " + fmt.Sprint(a...)))
}
