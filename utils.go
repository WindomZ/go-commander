package commander

import "sort"

func sortStringMap(m map[string]string) []string {
	if m == nil || len(m) == 0 {
		return []string{}
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	r := make([]string, len(keys))
	sort.Strings(keys)
	for i, k := range keys {
		r[i] = m[k]
	}
	return r
}
