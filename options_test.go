package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestOptions_UsagesString(t *testing.T) {
	var o _Options

	o = _Options{}
	assert.Equal(t, o.UsagesString(), []string(nil))

	o = _Options{
		newOption("-a, --about", "about description"),
		newOption("-b=<kn>, --bold=<kn>", "bold description"),
		newOption("-c, --config", "config description"),
		newOption("-d, --drop", "drop description"),
	}
	assert.Equal(t, o.UsagesString(),
		[]string{"[-a|--about] [-b=<kn>|--bold=<kn>] [-c|--config] [-d|--drop]"})
}

func TestOptions_OptionsString(t *testing.T) {
	o := _Options{
		newOption("-a, --about", "about description"),
		newOption("-b=<kn>, --bold=<kn>", "bold description"),
		newOption("-c, --config", "config description"),
		newOption("-d, --drop", "drop description"),
	}
	opts := o.OptionsString()
	assert.Equal(t, opts["-a|--about"], "-a --about    about description")
	assert.Equal(t, opts["-b|--bold"], "-b=<kn> --bold=<kn>\n              bold description")
	assert.Equal(t, opts["-c|--config"], "-c --config   config description")
	assert.Equal(t, opts["-d|--drop"], "-d --drop     drop description")
}
