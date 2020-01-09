package util

func Have(length int, f func(i int) bool) bool {
	for index := 0; index < length; index++ {
		if f(index) {
			return true
		}
	}
	return false
}
