package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestFlags_1(t *testing.T) {
	f := newFlag("-c")

	assert.Equal(t, f.Name(), "-c")
	assert.Equal(t, f.IsArgumentsRequired(), false)
	assert.Equal(t, f.IsArgumentsOptional(), false)
	assert.Equal(t, f.UsageString(), "-c")
	assert.Equal(t, f.OptionString(), "-c")
}

func TestFlags_2(t *testing.T) {
	f := newFlag("-c, --config")

	assert.Equal(t, f.Name(), "--config")
	assert.Equal(t, f.IsArgumentsRequired(), false)
	assert.Equal(t, f.IsArgumentsOptional(), false)
	assert.Equal(t, f.UsageString(), "-c|--config")
	assert.Equal(t, f.OptionString(), "-c, --config")
}

func TestFlags_3(t *testing.T) {
	f := newFlag("-c, --config <path>")

	assert.Equal(t, f.Name(), "--config")
	assert.Equal(t, f.IsArgumentsRequired(), true)
	assert.Equal(t, f.IsArgumentsOptional(), false)
	assert.Equal(t, f.UsageString(), "(-c|--config)=<path>")
	assert.Equal(t, f.OptionString(), "-c <path> --config=<path>")
}

func TestFlags_4(t *testing.T) {
	f := newFlag("--config, -c [type]")

	assert.Equal(t, f.Name(), "--config")
	assert.Equal(t, f.IsArgumentsRequired(), false)
	assert.Equal(t, f.IsArgumentsOptional(), true)
	assert.Equal(t, f.UsageString(), "(--config|-c)=[type]")
	assert.Equal(t, f.OptionString(), "--config=[type] -c [type]")
}

func TestFlags_5(t *testing.T) {
	f := newFlag("--config| -c [type]")

	assert.Equal(t, f.Name(), "--config")
	assert.Equal(t, f.IsArgumentsRequired(), false)
	assert.Equal(t, f.IsArgumentsOptional(), true)
	assert.Equal(t, f.UsageString(), "(--config|-c)=[type]")
	assert.Equal(t, f.OptionString(), "--config=[type] -c [type]")
}

func TestFlags_6(t *testing.T) {
	f := newFlag("--config  -c [type]")

	assert.Equal(t, f.Name(), "--config")
	assert.Equal(t, f.IsArgumentsRequired(), false)
	assert.Equal(t, f.IsArgumentsOptional(), true)
	assert.Equal(t, f.UsageString(), "(--config|-c)=[type]")
	assert.Equal(t, f.OptionString(), "--config=[type] -c [type]")
}

func TestFlags_7(t *testing.T) {
	f := newFlag("--config, -c = [type] a")
	assert.Equal(t, f.regexpFlags(), []string{"--config", "-c"})
	assert.Equal(t, f.regexpArguments(), []string{"[type]"})

	f = newFlag("--config| -c [type] b")
	assert.Equal(t, f.regexpFlags(), []string{"--config", "-c"})
	assert.Equal(t, f.regexpArguments(), []string{"[type]"})

	f = newFlag("--config  -c=<type> c")
	assert.Equal(t, f.regexpFlags(), []string{"--config", "-c"})
	assert.Equal(t, f.regexpArguments(), []string{"<type>"})

	f = newFlag("--config  -c=123")
	assert.Equal(t, f.regexpFlags(), []string{"--config", "-c"})
	assert.Equal(t, f.regexpArguments(), []string{"123"})
}
