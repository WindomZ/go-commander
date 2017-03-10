package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestActor_Action(t *testing.T) {
	var result bool
	var a actor
	a.Action(func(c Context) {
		result = true
		assert.Equal(t, c.GetBool("a"), true)
		assert.Equal(t, c.GetBool("b"), true)
		assert.Equal(t, c.GetBool("c"), true)
		assert.Equal(t, c.GetBool("d"), false)
	}, []string{"a", "b", "c"})

	a.run(newContext(nil, newDocoptMap(
		map[string]interface{}{
			"a": true,
			"b": true,
			"c": false,
		},
	),
	))
	assert.Equal(t, result, false)

	a.run(newContext(nil, newDocoptMap(
		map[string]interface{}{
			"a": true,
			"b": true,
			"c": true,
		},
	),
	))
	assert.Equal(t, result, true)
}
