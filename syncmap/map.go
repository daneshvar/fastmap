package syncmap

import (
	"sync"
)

type bucket struct {
	values map[uint]int64
	lock   sync.Mutex
}

type Map struct {
	values []bucket
}

func New() *Map {
	m := Map{
		values: make([]bucket, 10),
	}
	for i := 0; i < len(m.values); i++ {
		m.values[i].values = make(map[uint]int64)
	}
	return &m
}

func (m *Map) getBucket(key uint) *bucket {
	i := key % uint(len(m.values))
	return &m.values[i]
}

func (m *Map) Store(key uint, value int64) {
	s := m.getBucket(key)
	s.lock.Lock()
	defer s.lock.Unlock()
	s.values[key] = value
}

func (m *Map) Load(key uint) (value int64, ok bool) {
	s := m.getBucket(key)
	s.lock.Lock()
	defer s.lock.Unlock()
	value, ok = s.values[key]
	return
}
