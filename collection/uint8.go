package collection

import "github.com/t3reezhou/golang_util/filter"

func Uint8(length int, f func(i int) uint8) []uint8 {
	result := make([]uint8, length)
	filter.Filter(length, func(i int) bool { return true }, func(i int) { result[i] = f(i) })
	return result
}
