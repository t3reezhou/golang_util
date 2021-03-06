// Code generated by "collection --type=int"; DO NOT EDIT.

package collection

import "github.com/t3reezhou/golang_util/filter"

func Int(length int, do func(i int) int, funcs ...CollecitonFunc) []int {
	var o CollecitonOptions
	for _, f := range funcs {
		f(&o)
	}
	if o.judge == nil {
		o.judge = func(i int) bool { return true }
	}
	result := make([]int, length)
	filter.Filter(length, o.judge, func(i int) { result[i] = do(i) })
	return result
}
