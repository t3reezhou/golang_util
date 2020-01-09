package collection

import "github.com/t3reezhou/golang_util/filter"

func Uint(length int, f func(i int) uint) []uint {
	result := make([]uint, length)
	filter.Filter(length, func(i int) bool { return true }, func(i int) { result[i] = f(i) })
	return result
}
