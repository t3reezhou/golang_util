package have

import util "github.com/t3reezhou/golang_util"

func Have(length int, f func(i int) bool) bool {
	return util.Have(length, f)
}
