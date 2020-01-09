package reduce

func ReduceString(f func(x, y string) string, ints []string, initNum *string) (result string, err error) {
	if len(ints) == 0 {
		return "", errSliceEmpty
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
