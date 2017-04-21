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
	line      bool       // show on options in new line
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
		desc, _ := args[0].(string)
		o.Description(desc)
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

func (o _Option) Name() string {
	return strings.Join(o.names, "|")
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

func (o *_Option) Description(desc string) *_Option {
	o.desc = strings.TrimSpace(desc)
	o.show = len(o.desc) != 0
	return o
}

func (o *_Option) Aliases(aliases []string) *_Option {
	s := strings.TrimSpace(strings.Join(aliases, " "))
	if len(s) != 0 {
		o.usage += " " + s
		o.regexpNames()
		o.regexpArguments()
		o.regexpRequired()
	}
	return o
}

func (o _Option) UsageString(ones ...bool) string {
	if ok, _ := regexp.MatchString(`^[\[(].+[)\]]$`, o.usage); !ok || len(o.names) > 1 {
		strs := make([]string, len(o.names))
		arg := strings.Join(o.arguments.Get(), " ")
		for i, name := range o.names {
			if len(arg) != 0 {
				strs[i] = name + "=" + arg
			} else {
				strs[i] = name
			}
		}
		s := strings.Join(strs, "|")

		if o.line && len(ones) != 0 && ones[0] {
		} else if o.line || o.IsRequired() {
			s = fmt.Sprintf("(%s)", s)
		} else {
			s = fmt.Sprintf("[%s]", s)
		}
		return s
	}
	return regexp.MustCompile(`(\s*,\s*-)|(\s-)`).ReplaceAllString(o.usage, "|-")
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
	s = FormatDescription(s, o.desc)
	return
}

func (o _Option) run(c Context) _Result {
	if r := o.actor.run(c); r != nil && r.Break() {
		return r
	}
	return nil
}
