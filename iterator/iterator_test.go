package iterator

import (
	"testing"
	// . "github.com/smartystreets/goconvey/convey"
)

func TestIterator(t *testing.T) {
	data := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	emptyData := []int64{}
	// Convey("out of limit", t, func()
	{
		result := make([]int64, 0, 10)
		Iterator(len(data), 15, func(i, j int) error {
			result = append(result, data[i:j]...)
			return nil
		})
		if !SoInt64(len(result), func(i int) bool { return result[i] == data[i] }) {
			t.Errorf("result should be %+v, but it is %+v", data, result)
		}
	}
	{
		result := make([]int64, 0, 10)
		Iterator(len(data), 3, func(i, j int) error {
			result = append(result, data[i:j]...)
			return nil
		})

		if !SoInt64(len(result), func(i int) bool { return result[i] == data[i] }) {
			t.Errorf("result should be %+v, but it is %+v", data, result)
		}
	}
	// Convey("normal", t, func()
	{
		result := make([]int64, 0, 10)
		Iterator(len(data), 2, func(i, j int) error {
			result = append(result, data[i:j]...)
			return nil
		})
		if !SoInt64(len(result), func(i int) bool { return result[i] == data[i] }) {
			t.Errorf("result should be %+v, but it is %+v", data, result)
		}

	}
	// Convey("data is empty", t, func()
	{
		result := make([]int64, 0, 10)
		Iterator(len(emptyData), 2, func(i, j int) error {
			result = append(result, emptyData[i:j]...)
			return nil
		})
		if l := len(result); l > 0 {
			t.Errorf("result's length should be %d, but it is %d", 0, l)
		}
	}
}

func SoInt64(length int, shouldBe func(i int) bool) bool {
	for index := 0; index < length; index++ {
		if !shouldBe(index) {
			return false
		}
	}
	return true
}
