package util

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIterator(t *testing.T) {
	data := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	emptyData := []int64{}
	Convey("out of limit", t, func() {
		result := make([]int64, 0, 10)
		Iterator(len(data), 15, func(i, j int) error {
			result = append(result, data[i:j]...)
			return nil
		})
		So(result, ShouldResemble, data)
	})
	Convey("last one out of limit", t, func() {
		result := make([]int64, 0, 10)
		Iterator(len(data), 3, func(i, j int) error {
			result = append(result, data[i:j]...)
			return nil
		})
		So(result, ShouldResemble, data)
	})
	Convey("normal", t, func() {
		result := make([]int64, 0, 10)
		Iterator(len(data), 2, func(i, j int) error {
			result = append(result, data[i:j]...)
			return nil
		})
		So(result, ShouldResemble, data)
	})
	Convey("data is empty", t, func() {
		result := make([]int64, 0, 10)
		Iterator(len(emptyData), 2, func(i, j int) error {
			result = append(result, emptyData[i:j]...)
			return nil
		})
		So(result, ShouldResemble, emptyData)
	})
}
