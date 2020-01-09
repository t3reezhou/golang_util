package filters

func Int(ss []int, fs []func(i int) bool) [][]int {
	result := make([][]int, len(fs)+1)
	Filters(len(ss), fs, func(i, j int) { result[i] = append(result[i], ss[j]) })
	return result
}
