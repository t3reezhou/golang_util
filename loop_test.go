package util

import (
	"testing"
)

func TestLoop(t *testing.T) {
	loopTime := 0
	o, l := 0, 1000
	Loop(o, l, func(offset, limit int) (int, error) {
		if o != offset {
			t.Errorf("offset should be %d, but it is %d", o, offset)
		}
		if offset == 3000 {
			return 100, nil
		}
		loopTime++
		o += l
		return limit, nil
	})
	if loopTime != 3 {
		t.Errorf("loopTime should be 3, but it is %d", loopTime)
	}
}
