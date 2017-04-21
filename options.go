package commander

import "strings"

// _Options
type _Options []*_Option

func (o _Options) IsEmpty() bool {
	return len(o) == 0
}

func (o _Options) UsagesString(ones ...bool) (r []string) {
	if len(o) == 0 {
		return
	}
	var one bool
	if len(o) == 1 && len(ones) != 0 {
		one = ones[0]
	}
	rs := make([]string, 0, len(o))
	for _, opt := range o {
		if opt.line {
			r = append(r, opt.UsageString(one))
		} else {
			rs = append(rs, opt.UsageString(one))
		}
	}
	if len(rs) != 0 {
		r = append(r, strings.Join(rs, " "))
	}
	return
}

func (o _Options) OptionsString() (r map[string]string) {
	r = make(map[string]string, len(o))
	for _, opt := range o {
		if s := opt.OptionString(); len(s) != 0 {
			r[opt.Name()] = s
		}
	}
	return
}

func (o _Options) run(c Context) (result _Result) {
	for _, opt := range o {
		if r := opt.run(c); r != nil && r.Break() {
			if result = r; r.Error() != nil {
				break
			}
		}
	}
	return
}
