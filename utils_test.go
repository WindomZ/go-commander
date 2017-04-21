package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestUtils_sortStringMap(t *testing.T) {
	assert.Empty(t, sortStringMap(nil))
	assert.Empty(t, sortStringMap(make(map[string]string)))

	assert.Equal(t, sortStringMap(map[string]string{
		"1": "1",
		"3": "3",
		"2": "2",
	}), []string{"1", "2", "3"})
}
