// Code generated by "filtersf --type=int32"; DO NOT EDIT.

package filtersf

import util "github.com/t3reezhou/golang_util"

func Int32(ss []int32, fs []func(i int) bool, shouldReturn bool) [][]int32 {
	result := make([][]int32, len(fs)+1)
	util.FiltersF(len(ss), fs, func(i, j int) { result[i] = append(result[i], ss[j]) }, shouldReturn)
	return result
}
