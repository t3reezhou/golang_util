// Code generated by "collection --type=uint"; DO NOT EDIT.

package collection

import "github.com/t3reezhou/golang_util/filter"

func Uint(length int, do func(i int) uint, funcs ...CollecitonFunc) []uint {
	var o CollecitonOptions
	for _, f := range funcs {
		f(&o)
	}
	if o.judge == nil {
		o.judge = func(i int) bool { return true }
	}
	result := make([]uint, length)
	filter.Filter(length, o.judge, func(i int) { result[i] = do(i) })
	return result
}
