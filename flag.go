package commander

import "strings"

type Flags struct {
	name  string
	flag  string
	flags []string
}

func newFlag(flags ...string) *Flags {
	f := packFlags(flags...)
	return &Flags{
		flag:  f,
		flags: parseFlags(f),
	}
}

func (f *Flags) Name() string {
	if len(f.name) == 0 && len(f.flags) != 0 {
		for _, flag := range f.flags {
			if len(flag) > len(f.name) {
				f.name = flag
			}
		}
	}
	return f.name
}

func (f Flags) IsRequired() bool {
	return strings.Contains(f.flag, "<")
}

func (f Flags) IsOptional() bool {
	return strings.Contains(f.flag, "[")
}

func (f Flags) IsValid() bool {
	return f.flags != nil && len(f.flags) >= 1
}

func packFlags(flags ...string) string {
	if flags == nil || len(flags) == 0 {
		return ""
	} else if len(flags) == 1 {
		return flags[0]
	}
	return strings.Join(flags, ",")
}

func parseFlags(flag string) []string {
	flags := strings.Split(flag, ",")
	for i, f := range flags {
		flags[i] = strings.TrimSpace(f)
	}
	return flags
}
