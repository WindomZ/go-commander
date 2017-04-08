package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestError_newError(t *testing.T) {
	assert.Equal(t, newError("test").Error(), "go-commander: test")
}

func TestError_panicError(t *testing.T) {
	defer func() {
		assert.NotEmpty(t, recover())
	}()
	panicError("test")
	assert.Fail(t, "fail to panic error!")
}
