package filtersd

import util "github.com/t3reezhou/golang_util"

func FiltersD(length int, fs []func(i int) bool, do func(i, j int), shouldReturn bool) {
	util.FiltersD(length, fs, do, shouldReturn)
}
