package commander

import "fmt"

type Options []*Option

func (o Options) UsagesString() (r []string) {
	for _, opt := range o {
		if opt.IsRequired() {
			r = append(r, fmt.Sprintf("(%s)",
				opt.UsageString()))
		} else {
			r = append(r, fmt.Sprintf("[%s]",
				opt.UsageString()))
		}
	}
	return
}

func (o Options) OptionsString() (r []string) {
	for _, opt := range o {
		r = append(r, opt.OptionString())
	}
	return
}
