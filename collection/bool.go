package collection

import "github.com/t3reezhou/golang_util/filter"

func Bool(length int, f func(i int) bool) []bool {
	result := make([]bool, length)
	filter.Filter(length, func(i int) bool { return true }, func(i int) { result[i] = f(i) })
	return result
}
