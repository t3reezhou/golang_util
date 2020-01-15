package collection

//go:generate echo "start"
//go:generate go run cmd/collection/main.go --type=int
//go:generate go run cmd/collection/main.go --type=int8
//go:generate go run cmd/collection/main.go --type=int16
//go:generate go run cmd/collection/main.go --type=int32
//go:generate go run cmd/collection/main.go --type=int64
//go:generate go run cmd/collection/main.go --type=uint
//go:generate go run cmd/collection/main.go --type=uint8
//go:generate go run cmd/collection/main.go --type=uint16
//go:generate go run cmd/collection/main.go --type=uint32
//go:generate go run cmd/collection/main.go --type=uint64
//go:generate go run cmd/collection/main.go --type=bool
//go:generate go run cmd/collection/main.go --type=string
//go:generate echo "end"

type CollecitonOptions struct {
	judge func(i int) bool
}
type CollecitonFunc func(*CollecitonOptions)

func Judge(judge func(i int) bool) CollecitonFunc {
	return func(o *CollecitonOptions) { o.judge = judge }
}
