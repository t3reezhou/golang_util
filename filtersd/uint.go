// Code generated by "filtersd --type=uint"; DO NOT EDIT.

package filtersd

import util "github.com/t3reezhou/golang_util"

func Uint(ss []uint, fs []func(i int) bool, shouldReturn bool) [][]uint {
	result := make([][]uint, len(fs)+1)
	util.FiltersD(len(ss), fs, func(i, j int) { result[i] = append(result[i], ss[j]) }, shouldReturn)
	return result
}
