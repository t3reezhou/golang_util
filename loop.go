package util

func Loop(offset, limit int, do func(int, int) (int, error)) error {
	if offset < 0 || limit < 0 {
		return nil
	}
	for {
		l, err := do(offset, limit)
		if err != nil {
			return err
		}
		if l <= 0 || l < limit {
			return nil
		}
		offset += l
	}
}
