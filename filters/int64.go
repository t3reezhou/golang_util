package filters

func Int64(ss []int64, fs []func(i int) bool) [][]int64 {
	result := make([][]int64, len(fs)+1)
	Filters(len(ss), fs, func(i, j int) { result[i] = append(result[i], ss[j]) })
	return result
}
