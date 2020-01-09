package reduce

func ReduceInt64(f func(x, y int64) int64, ints []int64, initNum *int64) (result int64, err error) {
	if len(ints) == 0 {
		return -1, errSliceEmpty
	}
	if initNum != nil {
		result = *initNum
	} else {
		result = ints[0]
		ints = ints[1:]
	}
	For(len(ints), func(i int) { result = f(result, ints[i]) })
	return
}
