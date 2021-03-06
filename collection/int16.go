// Code generated by "collection --type=int16"; DO NOT EDIT.

package collection

import "github.com/t3reezhou/golang_util/filter"

func Int16(length int, do func(i int) int16, funcs ...CollecitonFunc) []int16 {
	var o CollecitonOptions
	for _, f := range funcs {
		f(&o)
	}
	if o.judge == nil {
		o.judge = func(i int) bool { return true }
	}
	result := make([]int16, length)
	filter.Filter(length, o.judge, func(i int) { result[i] = do(i) })
	return result
}
