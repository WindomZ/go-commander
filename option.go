package commander

import (
	"fmt"
	"regexp"
	"strings"
)

// _Option
type _Option struct {
	actor                // actor
	usage     string     // usage
	required  bool       // option required
	desc      string     // desc
	show      bool       // show on options
	arguments _Arguments // arguments
}

func newOption(usage string, args ...interface{}) *_Option {
	o := &_Option{
		usage: strings.TrimSpace(usage),
	}
	o.regexpNames()
	o.regexpArguments()
	o.regexpRequired()
	if len(args) >= 1 {
		o.desc, _ = args[0].(string)
		o.show = len(o.desc) != 0
	}
	if len(args) >= 2 {
		o.setAction(args[1])
	}
	if len(args) >= 3 {
		defs := make([]string, 0, len(args)-2)
		for _, arg := range args[2:] {
			defs = append(defs, fmt.Sprintf("%v", arg))
		}
		def := fmt.Sprintf("[default: %v]", strings.Join(defs, ","))
		if len(o.desc) != 0 {
			o.desc += " " + def
		} else {
			o.desc = def
		}
	}
	return o
}

func (o *_Option) regexpNames() {
	o.names = regexpOption(o.usage)
}

func (o *_Option) regexpArguments() {
	o.arguments.Set(o.usage)
}

func (o *_Option) regexpRequired() {
	if strings.HasPrefix(o.usage, "(") {
		o.required = true
	}
}

func (o _Option) Valid() bool {
	return len(o.names) != 0 && len(o.usage) != 0
}

func (o _Option) Names() []string {
	return o.names
}

func (o _Option) IsRequired() bool {
	return o.required
}

func (o _Option) IsOptional() bool {
	return !o.IsRequired()
}

func (o _Option) UsageString(ones ...bool) (s string) {
	if ok, _ := regexp.MatchString(`^[\[(].+[)\]]$`, o.usage); ok {
		s = o.usage
	} else if len(ones) != 0 && ones[0] {
		s = o.usage
	} else if o.IsRequired() {
		s = fmt.Sprintf("(%s)", o.usage)
	} else {
		s = fmt.Sprintf("[%s]", o.usage)
	}
	s = regexp.MustCompile(`(\s*,\s*-)|(\s-)`).ReplaceAllString(s, "|-")
	return
}

func (o _Option) OptionString() (s string) {
	if !o.show {
		return ""
	}
	s = regexp.MustCompile(`^[\[(].+[)\]]$`).
		ReplaceAllStringFunc(o.usage, func(str string) string {
			if len(str) > 2 {
				return str[1 : len(str)-1]
			}
			return str
		})
	s = regexp.MustCompile(`(\s*[,|]\s*-)`).ReplaceAllString(s, " -")
	if len(s) >= 12 {
		s = fmt.Sprintf("%s  %s", s, o.desc)
	} else {
		s = fmt.Sprintf("%-14s%s", s, o.desc)
	}
	return
}

func (o _Option) run(c Context) _Result {
	if r := o.actor.run(c); r != nil && r.Break() {
		return r
	}
	return nil
}
