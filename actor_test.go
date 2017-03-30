package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestActor_IncludeKeys(t *testing.T) {
	var a actor
	a.addIncludeKeys([]string{"a", "b", "c"})
	assert.Equal(t, a.getIncludeKeys(), []string{"a", "b", "c"})
}

func TestActor_ExcludeKeys(t *testing.T) {
	var a actor
	a.addExcludeKeys([]string{"a", "b", "c"})
	assert.Equal(t, a.getExcludeKeys(), []string{"a", "b", "c"})
}

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
