package commander

import (
	"fmt"
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestActor_IncludeKeys(t *testing.T) {
	var a actor
	var pass bool
	a.addIncludeKeys([]string{"a", "b", "c"})
	for _, key := range a.getIncludeKeys() {
		pass = false
		for _, k := range []string{"a", "b", "c"} {
			if key == k {
				pass = true
			}
		}
		if !pass {
			assert.FailNow(t, fmt.Sprintf("Error: %v", a.getIncludeKeys()))
		}
	}
}

func TestActor_ExcludeKeys(t *testing.T) {
	var a actor
	var pass bool
	a.addExcludeKeys([]string{"a", "b", "c"})
	for _, key := range a.getExcludeKeys() {
		pass = false
		for _, k := range []string{"a", "b", "c"} {
			if key == k {
				pass = true
			}
		}
		if !pass {
			assert.FailNow(t, fmt.Sprintf("Error: %v", a.getExcludeKeys()))
		}
	}
}

func TestActor_Action(t *testing.T) {
	var result bool
	var a actor
	a.Action(func(c Context) {
		result = true
		assert.Equal(t, c.MustBool("a"), true)
		assert.Equal(t, c.MustBool("b"), true)
		assert.Equal(t, c.MustBool("c"), true)
		assert.Equal(t, c.MustBool("d"), false)
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
