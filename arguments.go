package commander

// _Arguments the collection of _Argument implementation.
type _Arguments []*_Argument

// IsEmpty returns true if this _Arguments is empty.
func (a _Arguments) IsEmpty() bool {
	return len(a) == 0
}

// Set append usage into this _Arguments.
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

// Get returns a slice contains all names of this _Arguments.
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
