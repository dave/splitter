package splitter

import (
	"github.com/gopherjs/gopherjs/js"
)

func New(name string) *Split {
	return &Split{name: name}
}

func (s *Split) Init(args ...interface{}) {
	s.Object = js.Global.Call("Split", args...)
}

func (s *Split) SetSizes(sizes []float64) {
	out := make([]interface{}, len(sizes))
	for i, v := range sizes {
		out[i] = v
	}
	s.Object.Call("setSizes", out)
}

func (s *Split) SetSizesIfChanged(sizes []float64) {
	current := s.GetSizes()
	if len(current) != len(sizes) {
		s.SetSizes(sizes)
		return
	}
	for i, v := range sizes {
		if current[i] != v {
			s.SetSizes(sizes)
			return
		}
	}
}

func (s *Split) GetSizes() []float64 {
	raw := s.Call("getSizes").Interface().([]interface{})
	var out []float64
	for _, v := range raw {
		if f, ok := v.(float64); ok {
			out = append(out, f)
		}
	}
	return out
}

type Split struct {
	*js.Object
	name string
}
