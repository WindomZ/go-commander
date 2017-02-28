package commander

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlags_1(t *testing.T) {
	f := newFlag("-c")

	assert.Equal(t, f.Name(), "-c")
	assert.Equal(t, f.IsRequired(), false)
	assert.Equal(t, f.IsOptional(), false)
	assert.Equal(t, f.String(), "-c")
}

func TestFlags_2(t *testing.T) {
	f := newFlag("-c, --config")

	assert.Equal(t, f.Name(), "--config")
	assert.Equal(t, f.IsRequired(), false)
	assert.Equal(t, f.IsOptional(), false)
	assert.Equal(t, f.String(), "-c, --config")
}

func TestFlags_3(t *testing.T) {
	f := newFlag("-c, --config <path>")

	assert.Equal(t, f.Name(), "--config")
	assert.Equal(t, f.IsRequired(), true)
	assert.Equal(t, f.IsOptional(), false)
	assert.Equal(t, f.String(), "-c, --config <path>")
}

func TestFlags_4(t *testing.T) {
	f := newFlag("--config, -c [type]")

	assert.Equal(t, f.Name(), "--config")
	assert.Equal(t, f.IsRequired(), false)
	assert.Equal(t, f.IsOptional(), true)
	assert.Equal(t, f.String(), "--config, -c [type]")
}

func TestFlags_5(t *testing.T) {
	f := newFlag("--config| -c [type]")

	assert.Equal(t, f.Name(), "--config")
	assert.Equal(t, f.IsRequired(), false)
	assert.Equal(t, f.IsOptional(), true)
	assert.Equal(t, f.String(), "--config | -c [type]")
}

func TestFlags_6(t *testing.T) {
	f := newFlag("--config  -c [type]")

	assert.Equal(t, f.Name(), "--config")
	assert.Equal(t, f.IsRequired(), false)
	assert.Equal(t, f.IsOptional(), true)
	assert.Equal(t, f.String(), "--config -c [type]")
}
