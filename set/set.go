package set

import "sync"

type Set struct {
	setStruct map[any]struct{}
	size      int64
	mutex     sync.Mutex
}

func New() *Set {
	return &Set{
		setStruct: make(map[any]struct{}),
		size:      0,
		mutex:     sync.Mutex{},
	}
}
func (s *Set) Fill(items ...any) {
	for _, item := range items {
		s.mutex.Lock()
		s.setStruct[item] = struct{}{}
		s.size += 1
		s.mutex.Unlock()
	}
}

func (s *Set) Add(item any) {
	s.mutex.Lock()
	s.setStruct[item] = struct{}{}
	s.mutex.Unlock()
}

func (s *Set) GetItems() []any {
	var items []any
	for item := range s.setStruct {
		s.mutex.Lock()
		items = append(items, item)
		s.mutex.Unlock()
	}
	return items
}

func (s *Set) Intersection(secondSet *Set) {
	for k, _ := range s.setStruct {
		s.mutex.Lock()
		secondSet.setStruct[k] = struct{}{}
		s.mutex.Unlock()
	}
}
