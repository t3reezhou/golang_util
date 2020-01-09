package reduce

func String(f func(x, y string) string, ints []string, initNum *string) (result string) {
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
