// Code generated by "filtersf --type=int"; DO NOT EDIT.

package filtersf

import util "github.com/t3reezhou/golang_util"

func Int(ss []int, fs []func(i int) bool, shouldReturn bool) [][]int {
	result := make([][]int, len(fs)+1)
	util.FiltersF(len(ss), fs, func(i, j int) { result[i] = append(result[i], ss[j]) }, shouldReturn)
	return result
}
