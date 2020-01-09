package collection

import "github.com/t3reezhou/golang_util/filter"

func Int8(length int, f func(i int) int8) []int8 {
	result := make([]int8, length)
	filter.Filter(length, func(i int) bool { return true }, func(i int) { result[i] = f(i) })
	return result
}
