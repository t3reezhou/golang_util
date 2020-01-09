package reduce

func ReduceBool(f func(x, y bool) bool, ints []bool, initNum *bool) (result bool, err error) {
	if len(ints) == 0 {
		return false, errSliceEmpty
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
