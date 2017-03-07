package commander

import (
	"fmt"
	"regexp"
	"strings"
)

type Option struct {
	actor               // actor
	usage     string    // usage
	required  bool      // option required
	desc      string    // desc
	show      bool      // show on options
	arguments Arguments // arguments
}

func newOption(usage string, args ...interface{}) *Option {
	o := &Option{
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
		o.action, _ = args[1].(Action)
		if o.action == nil {
			o.action, _ = args[1].(func(args DocoptMap) Result)
		}
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

func (o *Option) regexpNames() {
	o.names = regexp.MustCompile(`-{1,2}[A-Za-z0-9_-]+\b`).
		FindAllString(regexp.MustCompile(`(<|\[)[A-Za-z0-9_\[\]<>-]+\b(>|])`).
			ReplaceAllString(o.usage, ""), -1)
}

func (o *Option) regexpArguments() {
	o.arguments.Set(o.usage)
}

func (o *Option) regexpRequired() {
	if strings.HasPrefix(o.usage, "(") {
		o.required = true
	}
}

func (o Option) Valid() bool {
	return len(o.names) != 0 && len(o.usage) != 0
}

func (o Option) Names() []string {
	return o.names
}

func (o Option) IsRequired() bool {
	return o.required
}

func (o Option) IsOptional() bool {
	return !o.IsRequired()
}

func (o Option) UsageString() (s string) {
	if ok, _ := regexp.MatchString(`^[\[(].+[)\]]$`, o.usage); ok {
		s = o.usage
	} else if o.IsRequired() {
		s = fmt.Sprintf("(%s)", o.usage)
	} else {
		s = fmt.Sprintf("[%s]", o.usage)
	}
	s = regexp.MustCompile(`(\s*,\s*-)|(\s-)`).ReplaceAllString(s, "|-")
	return
}

func (o Option) OptionString() (s string) {
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

func (o Option) run(d DocoptMap) Result {
	return o.actor.run(d)
}
