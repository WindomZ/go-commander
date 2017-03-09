package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestActor_Action(t *testing.T) {
	var result bool
	var a actor
	a.Action(func(c *Context) {
		result = true
		assert.Equal(t, c.Doc.GetMustBool("a"), true)
		assert.Equal(t, c.Doc.GetMustBool("b"), true)
		assert.Equal(t, c.Doc.GetMustBool("c"), true)
		assert.Equal(t, c.Doc.GetMustBool("d"), false)
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
