package utils

import "slices"

func Mapkey(m map[string][]string, value []string) (key string, ok bool) {
	for k, v := range m {
		if slices.Equal(v, value) {
			key = k
			ok = true
			return
		}
	}
	return
}
