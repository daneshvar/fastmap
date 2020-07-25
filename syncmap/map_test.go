package syncmap

import (
	"sync"
	"testing"
)

func TestStoreLoad(t *testing.T) {
	m := New()

	m.Store(568, 10)
	m.Store(210, 11)

	if v, ok := m.Load(568); !ok || v != 10 {
		t.Fail()
	}

	if v, ok := m.Load(210); !ok || v != 11 {
		t.Fail()
	}

	if _, ok := m.Load(12); ok {
		t.Fail()
	}

	m.Store(1234, 124)
	if v, ok := m.Load(1234); !ok || v != 124 {
		t.Fail()
	}
}

func BenchmarkMap(b *testing.B) {
	m := New()
	wg := &sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(uint(i), int64(i)*2)
		}(i)
	}
	wg.Wait()

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if v, ok := m.Load(uint(i)); !ok || v != int64(i)*2 {
				b.Fail()
			}
		}(i)
	}
	wg.Wait()
}

func BenchmarkSyncMap(b *testing.B) {
	m := &sync.Map{}

	wg := &sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(uint(i), int64(i)*2)
		}(i)
	}
	wg.Wait()

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if v, ok := m.Load(uint(i)); !ok || v != int64(i)*2 {
				b.Fail()
			}
		}(i)
	}
	wg.Wait()
}
