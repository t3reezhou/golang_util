package collection

import "github.com/t3reezhou/golang_util/filter"

func Uint16(length int, f func(i int) uint16) []uint16 {
	result := make([]uint16, length)
	filter.Filter(length, func(i int) bool { return true }, func(i int) { result[i] = f(i) })
	return result
}
