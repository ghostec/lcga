package ds

type Set struct {
	values map[interface{}]bool
}

func NewSet() *Set {
	return &Set{values: map[interface{}]bool{}}
}

func (s Set) Clone() *Set {
	ss := NewSet()
	for k, v := range s.values {
		ss.values[k] = v
	}
	return ss
}

func (s *Set) Add(i interface{}) {
	s.values[i] = true
}

func (s *Set) Remove(i interface{}) {
	delete(s.values, i)
}

func (s Set) Size() int {
	return len(s.values)
}

func (s Set) Slice() []interface{} {
	sl := make([]interface{}, 0, len(s.values))
	for k := range s.values {
		sl = append(sl, k)
	}
	return sl
}
