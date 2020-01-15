package filter

import util "github.com/t3reezhou/golang_util"

func Filter(length int, f func(i int) bool, do func(i int)) {
	util.Filter(length, f, do)
}
