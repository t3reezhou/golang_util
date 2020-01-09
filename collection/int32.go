package collection

import "github.com/t3reezhou/golang_util/filter"

func Int32(length int, f func(i int) int32) []int32 {
	result := make([]int32, length)
	filter.Filter(length, func(i int) bool { return true }, func(i int) { result[i] = f(i) })
	return result
}
