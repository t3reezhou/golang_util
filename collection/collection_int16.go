package collection

import "github.com/t3reezhou/golang_util/filter"

func CollectionInt16(length int, f func(i int) int16) []int16 {
	result := make([]int16, length)
	filter.Filter(length, func(i int) bool { return true }, func(i int) { result[i] = f(i) })
	return result
}
