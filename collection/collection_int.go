package collection

import "github.com/t3reezhou/golang_util/filter"

func CollectionInt(length int, f func(i int) int) []int {
	result := make([]int, length)
	filter.Filter(length, func(i int) bool { return true }, func(i int) { result[i] = f(i) })
	return result
}
