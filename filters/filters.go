package filters

import "github.com/t3reezhou/golang_util/filter"

func Filters(length int, fs []func(i int) bool, do func(i, j int)) {
	for i, f := range append(fs, func(int) bool { return true }) {
		filter.Filter(length, func(j int) bool { return f(j) }, func(j int) { do(i, j) })
	}
}
