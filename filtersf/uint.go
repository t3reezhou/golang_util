// Code generated by "filtersf --type=uint"; DO NOT EDIT.

package filtersf

import util "github.com/t3reezhou/golang_util"

func Uint(ss []uint, fs []func(i int) bool, shouldReturn bool) [][]uint {
	result := make([][]uint, len(fs)+1)
	util.FiltersF(len(ss), fs, func(i, j int) { result[i] = append(result[i], ss[j]) }, shouldReturn)
	return result
}
