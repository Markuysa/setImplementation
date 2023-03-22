package _map

import "sync"

type Map struct {
	_map  map[any]any
	mutex sync.Mutex
}

func New() *Map {
	return &Map{
		_map:  make(map[any]any),
		mutex: sync.Mutex{},
	}
}

func (m *Map) Values() []any {
	var values []any
	for _, v := range m._map {
		values = append(values, v)
	}
	return values
}
func (m *Map) Keys() []any {
	var keys []any
	for k, _ := range m._map {
		keys = append(keys, k)
	}
	return keys
}
