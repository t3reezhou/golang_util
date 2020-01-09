package pointer

func Int(i int) *int          { return &i }
func Bool(b bool) *bool       { return &b }
func Int64(i int64) *int64    { return &i }
func String(s string) *string { return &s }
