// Code generated by "filtersd --type=uint16"; DO NOT EDIT.

package filtersd

import util "github.com/t3reezhou/golang_util"

func Uint16(ss []uint16, fs []func(i int) bool, shouldReturn bool) [][]uint16 {
	result := make([][]uint16, len(fs)+1)
	util.FiltersD(len(ss), fs, func(i, j int) { result[i] = append(result[i], ss[j]) }, shouldReturn)
	return result
}