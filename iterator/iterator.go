package iterator

import util "github.com/t3reezhou/golang_util"

type Int64Iterator = util.Int64Iterator

func NewInt64Iterator(src []int64, limit int) *Int64Iterator {
	return util.NewInt64Iterator(src, limit)
}

func Iterator(lenght, limit int, do func(int, int) error) error {
	return util.Iterator(lenght, limit, do)
}
