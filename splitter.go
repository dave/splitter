package splitter

import (
	"github.com/gopherjs/gopherjs/js"
)

type Split struct {
	*js.Object
	name        string
	initialised bool
}

func New(name string) *Split {
	return &Split{name: name}
}

func (s *Split) Init(args ...interface{}) {
	s.Object = js.Global.Call("Split", args...)
	s.initialised = true
}

func (s *Split) Destroy(args ...interface{}) {
	s.Object.Call("destroy")
	s.initialised = false
}

func (s *Split) Initialised() bool {
	if s == nil {
		return false
	}
	return s.initialised
}

func (s *Split) SetSizes(sizes []float64) {
	out := make([]interface{}, len(sizes))
	for i, v := range sizes {
		out[i] = v
	}
	s.Object.Call("setSizes", out)
}

func (s *Split) Changed(sizes []float64) bool {
	current := s.GetSizes()
	if len(current) != len(sizes) {
		return true
	}
	for i, v := range sizes {
		if current[i] != v {
			return true
		}
	}
	return false
}

func (s *Split) SetSizesIfChanged(sizes []float64) {
	if s.Changed(sizes) {
		s.SetSizes(sizes)
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
