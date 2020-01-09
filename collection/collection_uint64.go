package collection

import "github.com/t3reezhou/golang_util/filter"

func CollectionUInt64(length int, f func(i int) uint64) []uint64 {
	result := make([]uint64, length)
	filter.Filter(length, func(i int) bool { return true }, func(i int) { result[i] = f(i) })
	return result
}
