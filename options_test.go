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
	assert.Equal(t, o.OptionsString(),
		[]string{
			"-a --about    about description",
			"-b=<kn> --bold=<kn>\n              bold description",
			"-c --config   config description",
			"-d --drop     drop description",
		})
}
