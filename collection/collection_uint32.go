package collection

import "github.com/t3reezhou/golang_util/filter"

func CollectionUInt32(length int, f func(i int) uint32) []uint32 {
	result := make([]uint32, length)
	filter.Filter(length, func(i int) bool { return true }, func(i int) { result[i] = f(i) })
	return result
}
