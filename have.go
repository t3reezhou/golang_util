package util

func Have(length int, f func(i int) bool) bool {
	var result bool
	Filter(length, f, func(i int) { result = true }, filterOptionsOnce())
	return result
}
