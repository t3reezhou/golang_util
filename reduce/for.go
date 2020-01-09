package reduce

import "errors"

var errSliceEmpty = errors.New("slice is empty and init_num is nil")

func IsErrSliceEmpty(err error) bool { return err == errSliceEmpty }

func ForR(length int, do func(i int)) {
	if length == 0 {
		return
	}
	do(length - 1)
	ForR(length-1, do)
}

func For(length int, do func(i int)) {
	for index := 0; index < length; index++ {
		do(index)
	}
}