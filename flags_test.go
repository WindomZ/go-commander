package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestFlags_1(t *testing.T) {
	f := newFlag("-c")

	assert.Equal(t, f.Name(), "-c")
	assert.Equal(t, f.IsRequired(), false)
	assert.Equal(t, f.IsOptional(), false)
	assert.Equal(t, f.UsageString(), "-c")
	assert.Equal(t, f.OptionString(), "-c")
}

func TestFlags_2(t *testing.T) {
	f := newFlag("-c, --config")

	assert.Equal(t, f.Name(), "--config")
	assert.Equal(t, f.IsRequired(), false)
	assert.Equal(t, f.IsOptional(), false)
	assert.Equal(t, f.UsageString(), "(-c|--config)")
}

func TestFlags_3(t *testing.T) {
	f := newFlag("-c, --config <path>")

	assert.Equal(t, f.Name(), "--config")
	assert.Equal(t, f.IsRequired(), true)
	assert.Equal(t, f.IsOptional(), false)
	assert.Equal(t, f.UsageString(), "(-c|--config)=<path>")
}

func TestFlags_4(t *testing.T) {
	f := newFlag("--config, -c [type]")

	assert.Equal(t, f.Name(), "--config")
	assert.Equal(t, f.IsRequired(), false)
	assert.Equal(t, f.IsOptional(), true)
	assert.Equal(t, f.UsageString(), "(--config|-c)=[type]")
}

func TestFlags_5(t *testing.T) {
	f := newFlag("--config| -c [type]")

	assert.Equal(t, f.Name(), "--config")
	assert.Equal(t, f.IsRequired(), false)
	assert.Equal(t, f.IsOptional(), true)
	assert.Equal(t, f.UsageString(), "(--config|-c)=[type]")
}

func TestFlags_6(t *testing.T) {
	f := newFlag("--config  -c [type]")

	assert.Equal(t, f.Name(), "--config")
	assert.Equal(t, f.IsRequired(), false)
	assert.Equal(t, f.IsOptional(), true)
	assert.Equal(t, f.UsageString(), "(--config|-c)=[type]")
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
}
