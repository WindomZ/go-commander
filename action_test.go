package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestAction_Action(t *testing.T) {
	assert.Equal(t,
		emptyAction(parseAction(func(c Context) _Result { return nil })), false)
	assert.Equal(t,
		emptyAction(parseAction(func() _Result { return nil })), false)
	assert.Equal(t,
		emptyAction(parseAction(func(c Context) error { return nil })), false)
	assert.Equal(t,
		emptyAction(parseAction(func(c Context) {})), false)
	assert.Equal(t,
		emptyAction(parseAction(func() {})), false)
	assert.Equal(t,
		emptyAction(parseAction(func() error { return nil })), false)
	assert.Equal(t,
		emptyAction(parseAction(func(m map[string]interface{}) error { return nil })), false)

	assert.Equal(t, emptyAction(parseAction(func() int { return 0 })), true)
}
