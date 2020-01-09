package util

func Filter(length int, f func(i int) bool, do func(i int)) {
	for index := 0; index < length; index++ {
		if f(index) {
			do(index)
		}
	}
}
