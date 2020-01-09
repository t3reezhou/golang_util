package loop

import util "github.com/t3reezhou/golang_util"

func Loop(offset, limit int, do func(int, int) (int, error)) error {
	return util.Loop(offset, limit, do)
}
