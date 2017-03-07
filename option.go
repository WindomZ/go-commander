package commander

import (
	"fmt"
	"regexp"
	"strings"
)

type Option struct {
	usage     string    // usage
	names     []string  // names
	required  bool      // option required
	desc      string    // desc
	arguments Arguments // arguments
	execFunc  ExecFunc  // exec function
	show      bool      // show on options
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
	}
	if len(args) >= 2 {
		o.execFunc, _ = args[1].(ExecFunc)
	}
	if len(args) >= 3 {
		def := fmt.Sprintf("[default: %v]", args[2])
		if len(o.desc) != 0 {
			o.desc += fmt.Sprintf(" %s", def)
		} else {
			o.desc = def
		}
	}
	if len(args) >= 4 {
		o.show, _ = args[3].(bool)
	} else {
		o.show = len(o.desc) != 0
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
	if strings.Contains(o.usage, "(") {
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
