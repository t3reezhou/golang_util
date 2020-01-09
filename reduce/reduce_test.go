package reduce

import "testing"

func TestReduce(t *testing.T) {
	result, err := ReduceInt(func(x, y int) int { return x + y }, []int{1, 2, 3, 4}, nil)
	if err != nil {
		t.Error(err.Error())
	}
	if result != 10 {
		t.Errorf("result should be 10,but it is %d", result)
	}

	result, err = ReduceInt(func(x, y int) int { return x * y }, []int{1, 2, 3, 4}, nil)
	if err != nil {
		t.Error(err.Error())
	}
	if result != 24 {
		t.Errorf("result should be 24,but it is %d", result)
	}
}

// 无法用递归实现reduce，因为无法在遇到边界问题的做任何处理
func reduceFunc(f func(x, y int) int, xy []int) int {
	if len(xy) == 0 {
		return 1 //  无法在遇到边界问题的做任何处理, 当f为x+y，那么应该return 0，当f为x * y,那么应该return 1
	}
	return f(xy[0], reduceFunc(f, xy[1:]))
}
