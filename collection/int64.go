package collection

import "github.com/t3reezhou/golang_util/filter"

func Int64(length int, f func(i int) int64) []int64 {
	result := make([]int64, length)
	filter.Filter(length, func(i int) bool { return true }, func(i int) { result[i] = f(i) })
	return result
}
