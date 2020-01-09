package example

import (
	"github.com/t3reezhou/golang_util/collection"
	"github.com/t3reezhou/golang_util/filter"
	"github.com/t3reezhou/golang_util/pointer"
	"github.com/t3reezhou/golang_util/reduce"
)

type File struct {
	ID   int64
	Name int64
	Size int64
}

type Files []*File

type Volume struct {
	ID        int64
	Name      string
	CreatorID int64
	CTime     int64
	MTime     int64
	DTime     int64
	Files     Files
}

type Volumes []*Volume

func (vs Volumes) Filter(f func(*Volume) bool) Volumes {
	result := make(Volumes, 0, len(vs))
	for _, v := range vs {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func (vs Volumes) Filter2(f func(*Volume) bool) Volumes {
	result := make(Volumes, 0, len(vs))
	filter.Filter(len(vs), func(i int) bool { return f(vs[i]) }, func(i int) { result = append(result, vs[i]) })
	return result
}

func (vs Volumes) IDs() []int64 {
	result := make([]int64, 0, len(vs))
	for _, v := range vs {
		result = append(result, v.ID)
	}
	return result
}

func (vs Volumes) IDs2() []int64 {
	result := make([]int64, 0, len(vs))
	filter.Filter(len(vs), func(i int) bool { return true }, func(i int) { result = append(result, vs[i].ID) })
	return result
}

func (vs Volumes) IDs3() []int64 {
	return collection.Int64(len(vs), func(i int) int64 { return vs[i].ID })
}

func (vs Volumes) HasOne(key func(*Volume) string) map[string]*Volume {
	result := make(map[string]*Volume)
	for _, v := range vs {
		result[key(v)] = v
	}
	return result
}

func (vs Volumes) HasOne2(key func(*Volume) string) map[string]*Volume {
	result := make(map[string]*Volume)
	filter.Filter(len(vs), func(i int) bool { return true }, func(i int) { result[key(vs[i])] = vs[i] })
	return result
}

func (vs Volumes) HasMany(key func(*Volume) string) map[string]Volumes {
	result := make(map[string]Volumes)
	for _, v := range vs {
		if result[key(v)] == nil {
			result[key(v)] = Volumes{v}
		} else {
			result[key(v)] = append(result[key(v)], v)
		}
	}
	return result
}

func (vs Volumes) HasMany2(key func(*Volume) string) map[string]Volumes {
	result := make(map[string]Volumes)
	filter.Filter(len(vs), func(i int) bool { return true }, func(i int) {
		if result[key(vs[i])] == nil {
			result[key(vs[i])] = Volumes{vs[i]}
		} else {
			result[key(vs[i])] = append(result[key(vs[i])], vs[i])
		}
	})
	return result
}

func (vs Volumes) SumFiles() int {
	var result int
	for _, v := range vs {
		result += len(v.Files)
	}
	return result
}

func (vs Volumes) SumFiles2() int {
	return reduce.Int(func(x, y int) int { return x + y },
		collection.Int(len(vs), func(i int) int { return len(vs[i].Files) }),
		pointer.Int(0))
}

func (fs Files) SumSize() int64 {
	var result int64
	for _, f := range fs {
		result += f.Size
	}
	return result
}

func (fs Files) SumSize2() int64 {
	return reduce.Int64(func(x, y int64) int64 { return x + y },
		collection.Int64(len(fs), func(i int) int64 { return fs[i].Size }),
		pointer.Int64(0))
}

func (vs Volumes) SumFilesSize() int64 {
	var result int64
	for _, v := range vs {
		result += v.Files.SumSize()
	}
	return result
}

func (vs Volumes) SumFilesSize2() int64 {
	return reduce.Int64(func(x, y int64) int64 { return x + y },
		collection.Int64(len(vs), func(i int) int64 { return vs[i].Files.SumSize2() }),
		pointer.Int64(0))
}
