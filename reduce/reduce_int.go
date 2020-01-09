package reduce

func ReduceInt(f func(x, y int) int, ints []int, initNum *int) (result int, err error) {
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
