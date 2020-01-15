package util

func FiltersF(length int, fs []func(i int) bool, do func(i, j int), shouldReturn bool) {
	for i, f := range append(fs, func(int) bool { return true }) {
		if shouldReturn {
			Filter(length, func(j int) bool { return f(j) }, func(j int) { do(i, j) }, filterOptionsOnce())
		} else {
			Filter(length, func(j int) bool { return f(j) }, func(j int) { do(i, j) })
		}
	}
}

func FiltersD(length int, fs []func(i int) bool, do func(i, j int), shouldReturn bool) {
	fs = append(fs, func(int) bool { return true })
	for i := 0; i < length; i++ {
		if shouldReturn {
			Filter(len(fs), func(j int) bool { return fs[j](i) }, func(j int) { do(j, i) }, filterOptionsOnce())
		} else {
			Filter(len(fs), func(j int) bool { return fs[j](i) }, func(j int) { do(j, i) })
		}
	}
}
