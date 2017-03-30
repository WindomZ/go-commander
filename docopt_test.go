package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestDocoptMap_Get(t *testing.T) {
	var m map[string]interface{} = map[string]interface{}{
		"string":  "abc",
		"strings": []string{"abc", "efg"},
		"bool":    true,
		"int":     1,
		"int8":    8,
		"int16":   16,
		"int32":   32,
		"int64":   64,
		"float32": 32,
		"float64": 64,
	}
	var d DocoptMap = newDocoptMap(m)

	assert.Equal(t, d.Contain("string"), true)
	assert.Equal(t, d.MustString("string"), "abc")

	assert.Equal(t, d.Contain("strings"), true)
	assert.Equal(t, d.MustStrings("strings"), []string{"abc", "efg"})
	assert.Equal(t, d.MustStrings("strings1"), []string{})

	assert.Equal(t, d.Contain("bool"), true)
	assert.Equal(t, d.MustBool("bool"), true)

	assert.Equal(t, d.Contain("int"), true)
	assert.Equal(t, d.MustInt("int"), 1)

	assert.Equal(t, d.Contain("int8"), true)
	assert.Equal(t, d.MustInt("int8"), 8)

	assert.Equal(t, d.Contain("int16"), true)
	assert.Equal(t, d.MustInt("int16"), 16)

	assert.Equal(t, d.Contain("int32"), true)
	assert.Equal(t, d.MustInt("int32"), 32)

	assert.Equal(t, d.Contain("int64"), true)
	assert.Equal(t, d.MustInt64("int64"), int64(64))

	assert.Equal(t, d.Contain("float32"), true)
	assert.Equal(t, d.MustFloat("float32"), float32(32))

	assert.Equal(t, d.Contain("float64"), true)
	assert.Equal(t, d.MustFloat64("float64"), float64(64))
}
