// Code generated by "filtersf --type=uint32"; DO NOT EDIT.

package filtersf

import util "github.com/t3reezhou/golang_util"

func Uint32(ss []uint32, fs []func(i int) bool, shouldReturn bool) [][]uint32 {
	result := make([][]uint32, len(fs)+1)
	util.FiltersF(len(ss), fs, func(i, j int) { result[i] = append(result[i], ss[j]) }, shouldReturn)
	return result
}
