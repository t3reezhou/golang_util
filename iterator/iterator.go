package iterator

type Int64Iterator struct {
	self              []int64
	start, end, limit int
}

func (i *Int64Iterator) HasNext() bool {
	return i.end < len(i.self)
}

func (i *Int64Iterator) Next() []int64 {
	if i.end+i.limit > len(i.self) {
		i.start, i.end = i.end, len(i.self)
	} else {
		i.start, i.end = i.end, i.end+i.limit
	}
	return i.self[i.start:i.end]
}

func NewInt64Iterator(src []int64, limit int) *Int64Iterator {
	if limit <= 0 || limit > len(src) {
		limit = len(src)
	}
	return &Int64Iterator{self: src, limit: limit}
}

func Iterator(lenght, limit int, do func(int, int) error) error {
	if limit > lenght {
		limit = lenght
	}
	for start, end := 0, 0; end < lenght; {
		if end+limit > lenght {
			start, end = end, lenght
		} else {
			start, end = end, end+limit
		}
		if err := do(start, end); err != nil {
			return err
		}
	}
	return nil
}
