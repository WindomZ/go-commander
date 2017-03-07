package commander

import (
	"fmt"
	"regexp"
	"strings"
)

type Flags struct {
	name string
	flag string
}

func newFlag(flag string) *Flags {
	return &Flags{
		flag: strings.TrimSpace(flag),
	}
}

func (f Flags) regexpFlags() []string {
	return regexp.MustCompile(`-{1,2}[A-Za-z0-9_-]+`).FindAllString(f.flag, -1)
}

func (f Flags) regexpArguments() []string {
	return append(
		regexp.MustCompile(`(?i:<|\[)[A-Za-z0-9_\[\]<>-]+(?i:>|])`).FindAllString(f.flag, -1),
		regexp.MustCompile(`[A-Z0-9]+`).FindAllString(f.flag, -1)...,
	)
}

func (f *Flags) longestFlag() (s string) {
	if flags := f.regexpFlags(); len(flags) != 0 {
		for _, flag := range flags {
			if len(flag) > len(s) {
				s = flag
			}
		}
	}
	return
}

func (f *Flags) shortestFlag() (s string) {
	if flags := f.regexpFlags(); len(flags) != 0 {
		s = flags[0]
		for _, flag := range flags {
			if len(flag) < len(s) {
				s = flag
			}
		}
	}
	return
}

func (f *Flags) Name() string {
	if len(f.name) == 0 && len(f.flag) != 0 {
		f.name = f.longestFlag()
	}
	return f.name
}

func (f Flags) UsageString() (s string) {
	flags := f.regexpFlags()
	if len(flags) != 0 {
		if len(flags) == 1 {
			s = flags[0]
		} else {
			s = strings.Join(flags, "|")
		}
	} else {
		return
	}
	if args := f.regexpArguments(); len(args) != 0 {
		if len(flags) != 1 {
			s = fmt.Sprintf("(%s)", s)
		}
		if len(args) == 1 {
			s += fmt.Sprintf("=%s", args[0])
		} else if f.IsArgumentsRequired() {
			s += fmt.Sprintf("=(%s)", strings.Join(args, "|"))
		} else if f.IsArgumentsOptional() {
			s += fmt.Sprintf("=[%s]", strings.Join(args, "|"))
		} else {
			s += fmt.Sprintf("=%s", strings.Join(args, "|"))
		}
	}
	return
}

func (f Flags) OptionString() (s string) {
	flags := f.regexpFlags()
	if len(flags) == 0 {
		return
	}
	if args := f.regexpArguments(); len(args) != 0 {
		if len(args) == 1 {
			s = fmt.Sprintf("%s", args[0])
		} else if f.IsArgumentsRequired() {
			s = fmt.Sprintf("(%s)", strings.Join(args, "|"))
		} else if f.IsArgumentsOptional() {
			s = fmt.Sprintf("[%s]", strings.Join(args, "|"))
		} else {
			s = fmt.Sprintf("%s", strings.Join(args, "|"))
		}
		var strs []string
		for _, flag := range flags {
			if strings.Count(flag, "-") == 1 {
				strs = append(strs,
					fmt.Sprintf("%s %s", flag, s))
			} else {
				strs = append(strs,
					fmt.Sprintf("%s=%s", flag, s))
			}
		}
		s = strings.Join(strs, " ")
	} else {
		s = strings.Join(flags, ", ")
	}
	return
}

func (f Flags) IsArgumentsRequired() bool {
	return strings.Contains(f.flag, "<")
}

func (f Flags) IsArgumentsOptional() bool {
	return strings.Contains(f.flag, "[")
}

func (f Flags) Valid() bool {
	return len(f.Name()) != 0
}
