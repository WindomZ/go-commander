package commander

import (
	"github.com/WindomZ/testify/assert"
	"testing"
)

func TestOptions_UsagesString(t *testing.T) {
	o := Options{
		newOption("-a, --about", "about description"),
		newOption("-b, --bold=<kn>", "bold description"),
		newOption("-c, --config", "config description"),
		newOption("-d, --drop", "drop description"),
	}
	assert.Equal(t, o.UsagesString(),
		[]string{
			"[-a|--about]",
			"[(-b|--bold)=<kn>]",
			"[-c|--config]",
			"[-d|--drop]",
		})
}

func TestOptions_OptionsString(t *testing.T) {
	o := Options{
		newOption("-a, --about", "about description"),
		newOption("-b, --bold=<kn>", "bold description"),
		newOption("-c, --config", "config description"),
		newOption("-d, --drop", "drop description"),
	}
	assert.Equal(t, o.OptionsString(),
		[]string{
			"-a, --about   about description",
			"-b <kn>, --bold <kn>  bold description",
			"-c, --config  config description",
			"-d, --drop    drop description",
		})
}
