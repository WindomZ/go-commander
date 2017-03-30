package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestOption_1(t *testing.T) {
	o := newOption("-p", "add pepper")

	assert.Equal(t, o.Names(), []string{"-p"})
	assert.Equal(t, o.IsRequired(), false)
	assert.Equal(t, o.IsOptional(), true)
	assert.Equal(t, o.UsageString(), "-p")
	assert.Equal(t, o.OptionString(), "-p            add pepper")
}

func TestOption_2(t *testing.T) {
	o := newOption("-p,--pepper", "add pepper")

	assert.Equal(t, o.Names(), []string{"-p", "--pepper"})
	assert.Equal(t, o.IsRequired(), false)
	assert.Equal(t, o.IsOptional(), true)
	assert.Equal(t, o.UsageString(), "[-p|--pepper]")
	assert.Equal(t, o.OptionString(), "-p --pepper   add pepper")
}

func TestOption_3(t *testing.T) {
	o := newOption("-p=<path>|--pepper=<path>", "add pepper directory")

	assert.Equal(t, o.Names(), []string{"-p", "--pepper"})
	assert.Equal(t, o.IsRequired(), false)
	assert.Equal(t, o.IsOptional(), true)
	assert.Equal(t, o.UsageString(), "[-p=<path>|--pepper=<path>]")
	assert.Equal(t, o.OptionString(), "-p=<path> --pepper=<path>\n              add pepper directory")
}

func TestOption_4(t *testing.T) {
	o := newOption("[-p=[path],--pepper=[path]]", "add pepper directory", func() {}, "xxx")

	assert.Equal(t, o.Names(), []string{"-p", "--pepper"})
	assert.Equal(t, o.IsRequired(), false)
	assert.Equal(t, o.IsOptional(), true)
	assert.Equal(t, o.UsageString(), "[-p=[path]|--pepper=[path]]")
	assert.Equal(t, o.OptionString(), "-p=[path] --pepper=[path]\n              add pepper directory [default: xxx]")
}

func TestOption_5(t *testing.T) {
	o := newOption("-p=[path]", "add pepper directory", func() {}, "xxx")

	assert.Equal(t, o.Names(), []string{"-p"})
	assert.Equal(t, o.IsRequired(), false)
	assert.Equal(t, o.IsOptional(), true)
	assert.Equal(t, o.UsageString(), "-p=[path]")
	assert.Equal(t, o.OptionString(), "-p=[path]     add pepper directory [default: xxx]")
}

func TestOption_6(t *testing.T) {
	o := newOption("(-p=[path])", "add pepper directory", func() {}, "xxx")

	assert.Equal(t, o.Names(), []string{"-p"})
	assert.Equal(t, o.IsRequired(), true)
	assert.Equal(t, o.IsOptional(), false)
	assert.Equal(t, o.UsageString(), "(-p=[path])")
	assert.Equal(t, o.OptionString(), "-p=[path]     add pepper directory [default: xxx]")
}
