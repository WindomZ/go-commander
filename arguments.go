package commander

// _Arguments Implementation of command line parameters
type _Arguments []*_Argument

func (a _Arguments) IsEmpty() bool {
	return len(a) == 0
}

func (a *_Arguments) Set(usage string) {
	*a = (*a)[:0]
	if strs := regexpArgument(usage); len(strs) != 0 {
		m := make(map[string]bool, len(strs))
		for _, str := range strs {
			if _, ok := m[str]; !ok {
				*a = append(*a, newArgument(str))
			}
			m[str] = true
		}
	}
}

func (a _Arguments) Get() (r []string) {
	for _, arg := range a {
		r = append(r, arg.name)
	}
	return
}

//func (a _Arguments) UsagesString() (r []string) {
//	for _, arg := range a {
//		r = append(r, arg.UsageString())
//	}
//	return
//}
//
//func (a _Arguments) OptionsString() (r []string) {
//	for _, arg := range a {
//		if s := arg.OptionString(); len(s) != 0 {
//			r = append(r, s)
//		}
//	}
//	return
//}
