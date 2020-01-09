package reduce

func Int(f func(x, y int) int, ints []int, initNum *int) (result int) {
	if len(ints) == 0 {
		if initNum == nil {
			panic(errSliceEmpty)
		}
		return *initNum
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
