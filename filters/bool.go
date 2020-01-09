package filters

func Bool(ss []bool, fs []func(i int) bool) [][]bool {
	result := make([][]bool, len(fs)+1)
	Filters(len(ss), fs, func(i, j int) { result[i] = append(result[i], ss[j]) })
	return result
}
