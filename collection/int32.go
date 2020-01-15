// Code generated by "collection --type=int32"; DO NOT EDIT.

package collection

import "github.com/t3reezhou/golang_util/filter"

func Int32(length int, do func(i int) int32, funcs ...CollecitonFunc) []int32 {
	var o CollecitonOptions
	for _, f := range funcs {
		f(&o)
	}
	if o.judge == nil {
		o.judge = func(i int) bool { return true }
	}
	result := make([]int32, length)
	filter.Filter(length, o.judge, func(i int) { result[i] = do(i) })
	return result
}
