package util

type filterOptions struct {
	once bool
}

type filterOptionsFunc func(*filterOptions)

func filterOptionsOnce() filterOptionsFunc {
	return func(o *filterOptions) { o.once = true }
}

func Filter(length int, f func(i int) bool, do func(i int), fs ...filterOptionsFunc) {
	var o filterOptions
	for _, f := range fs {
		f(&o)
	}
	for index := 0; index < length; index++ {
		if f(index) {
			do(index)
			if o.once {
				return
			}
		}
	}
}
