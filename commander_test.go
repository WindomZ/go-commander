package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestNewCommander(t *testing.T) {
	c := NewCommander("hello")
	assert.NotEmpty(t, c)
}
