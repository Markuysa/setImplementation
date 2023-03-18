package main

type Set struct {
	setStruct map[any]struct{}
	size      int64
}

func (s *Set) Init(items ...struct{}) {
	for item := range items {
		s.setStruct[item] = struct{}{}
	}
}
