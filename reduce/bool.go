package reduce

func Bool(f func(x, y bool) bool, ints []bool, initNum *bool) (result bool) {
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
