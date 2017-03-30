package commander

import (
	"errors"
	"github.com/WindomZ/testify/assert"
	"testing"
)

func testAction(a Action) bool {
	if emptyAction(a) {
		return false
	}
	a(newContext(nil, nil))
	return true
}

func TestAction_Action(t *testing.T) {
	var err error = errors.New("test error")

	assert.Equal(t,
		testAction(parseAction(func(c Context) _Result { return nil })), true)

	assert.Equal(t,
		testAction(parseAction(func() _Result { return nil })), true)

	assert.Equal(t,
		testAction(parseAction(func(c Context) error { return nil })), true)
	assert.Equal(t,
		testAction(parseAction(func(c Context) error { return err })), true)

	assert.Equal(t,
		testAction(parseAction(func(c Context) {})), true)

	assert.Equal(t,
		testAction(parseAction(func() {})), true)

	assert.Equal(t,
		testAction(parseAction(func() error { return nil })), true)
	assert.Equal(t,
		testAction(parseAction(func() error { return err })), true)

	assert.Equal(t,
		testAction(parseAction(func(m map[string]interface{}) error { return nil })), true)
	assert.Equal(t,
		testAction(parseAction(func(m map[string]interface{}) error { return err })), true)

	assert.Equal(t, testAction(parseAction(func() int { return 0 })), false)
}
