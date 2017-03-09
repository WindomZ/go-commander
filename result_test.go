package commander

import (
	"errors"
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestNewResult(t *testing.T) {
	r := NewResult("test")
	assert.Equal(t, r.ErrorString(), "test")
}

func TestNewResultCode(t *testing.T) {
	r := NewResultCode(1, "test")
	assert.Equal(t, r.Code(), 1)
	assert.Equal(t, r.ErrorString(), "test")
}

func TestNewResultError(t *testing.T) {
	r := NewResultError(errors.New("Test"), 2)
	assert.Equal(t, r.Code(), 2)
	assert.Equal(t, r.ErrorString(), "Test")
}
