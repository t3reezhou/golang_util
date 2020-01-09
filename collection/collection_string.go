package collection

import "github.com/t3reezhou/golang_util/filter"

func CollectionString(length int, f func(i int) string) []string {
	result := make([]string, length)
	filter.Filter(length, func(i int) bool { return true }, func(i int) { result[i] = f(i) })
	return result
}
