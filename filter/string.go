package filter

func Stirng(datas []string, f func(i int) bool) []string {
	result := make([]string, 0, len(datas))
	Filter(len(datas), f, func(i int) { result = append(result, datas[i]) })
	return result
}
