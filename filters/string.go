package filters

func String(ss []string, fs []func(i int) bool) [][]string {
	result := make([][]string, len(fs)+1)
	Filters(len(ss), fs, func(i, j int) { result[i] = append(result[i], ss[j]) })
	return result
}
