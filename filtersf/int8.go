// Code generated by "filtersf --type=int8"; DO NOT EDIT.

package filtersf

import util "github.com/t3reezhou/golang_util"

func Int8(ss []int8, fs []func(i int) bool, shouldReturn bool) [][]int8 {
	result := make([][]int8, len(fs)+1)
	util.FiltersF(len(ss), fs, func(i, j int) { result[i] = append(result[i], ss[j]) }, shouldReturn)
	return result
}
