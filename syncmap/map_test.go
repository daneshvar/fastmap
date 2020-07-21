package syncmap

import (
	"testing"
)

func TestStoreLoad(t *testing.T) {
	m := New()

	m.Store("ali", 10)
	m.Store("reza", 11)

	if v, ok := m.Load("ali"); !ok || v != 10 {
		t.Fail()
	}

	if v, ok := m.Load("reza"); !ok || v != 11 {
		t.Fail()
	}

	if v, ok := m.Load("al"); !ok || v != -1 {
		t.Fail()
	}

	if _, ok := m.Load("alireza"); ok {
		t.Fail()
	}

	m.Store("ali", 124)
	if v, ok := m.Load("ali"); !ok || v != 124 {
		t.Fail()
	}
}
