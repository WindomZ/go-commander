package commander

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOption_1(t *testing.T) {
	o := newOption("-p", "add pepper")

	assert.Equal(t, o.Name(), "-p")
	assert.Equal(t, o.IsRequired(), false)
	assert.Equal(t, o.IsOptional(), false)
	assert.Equal(t, o.OptionString(), "-p            add pepper")
}

func TestOption_2(t *testing.T) {
	o := newOption("-p,--pepper", "add pepper")

	assert.Equal(t, o.Name(), "--pepper")
	assert.Equal(t, o.IsRequired(), false)
	assert.Equal(t, o.IsOptional(), false)
	assert.Equal(t, o.OptionString(), "-p, --pepper  add pepper")
}

func TestOption_3(t *testing.T) {
	o := newOption("-p,--pepper <path>", "add pepper directory")

	assert.Equal(t, o.Name(), "--pepper")
	assert.Equal(t, o.IsRequired(), true)
	assert.Equal(t, o.IsOptional(), false)
	assert.Equal(t, o.OptionString(), "-p, --pepper  add pepper directory")
}

func TestOption_4(t *testing.T) {
	o := newOption("-p,--pepper [path]", "add pepper directory")

	assert.Equal(t, o.Name(), "--pepper")
	assert.Equal(t, o.IsRequired(), false)
	assert.Equal(t, o.IsOptional(), true)
	assert.Equal(t, o.OptionString(), "-p, --pepper  add pepper directory")
}
