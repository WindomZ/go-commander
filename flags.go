package commander

import "strings"

type Flags struct {
	name string
	flag string
}

func newFlag(flag string) *Flags {
	return &Flags{
		flag: packFlags(flag),
	}
}

func (f *Flags) Name() string {
	if len(f.name) == 0 && len(f.flag) != 0 {
		flags := parseFlags(f.flag)
		for _, flag := range flags {
			if len(flag) > len(f.name) {
				f.name = flag
			}
		}
	}
	return f.name
}

func (f Flags) String() string {
	return f.flag
}

func (f Flags) IsRequired() bool {
	return strings.Contains(f.flag, "<")
}

func (f Flags) IsOptional() bool {
	return strings.Contains(f.flag, "[")
}

func (f Flags) Valid() bool {
	return len(f.Name()) != 0 && len(f.flag) != 0
}

func packFlags(flag string) string {
	if flag = strings.TrimSpace(flag); len(flag) != 0 {
		flag = strings.Replace(flag, ",-", ", -", -1)
		flag = strings.Replace(flag, "| -", " | -", -1)
		flag = strings.Replace(flag, "|-", " | -", -1)
		for strings.Contains(flag, "  ") {
			flag = strings.Replace(flag, "  ", " ", -1)
		}
	}
	return flag
}

func parseFlags(flag string) (flags []string) {
	if strings.Contains(flag, ",") {
		flags = strings.Split(flag, ",")
	} else if strings.Contains(flag, "|") {
		flags = strings.Split(flag, "|")
	} else if strings.Contains(flag, " ") {
		flags = strings.Split(flag, " ")
	} else {
		flags = []string{flag}
	}
	for i, f := range flags {
		if i := strings.Index(f, "<"); i >= 0 {
			f = f[:i]
		} else if i := strings.Index(f, "["); i >= 0 {
			f = f[:i]
		}
		flags[i] = strings.TrimSpace(f)
	}
	return flags
}
