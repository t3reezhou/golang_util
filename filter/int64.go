package filter

func Int64(datas []int64, f func(i int) bool) []int64 {
	result := make([]int64, 0, len(datas))
	Filter(len(datas), f, func(i int) { result = append(result, datas[i]) })
	return result
}
