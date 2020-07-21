package syncmap

import (
	"fmt"
	"sync"
)

type Nodes map[rune]*leaf

type leaf struct {
	value int
	nodes Nodes
}

type Map struct {
	lock  sync.RWMutex
	nodes Nodes
}

func New() *Map {
	return &Map{
		nodes: make(map[rune]*leaf),
	}
}

func printNodes(prefix string, nodes Nodes) {
	for k, v := range nodes {
		fmt.Printf("%s%s:%d\n", prefix, string(k), v.value)
		if v.nodes != nil {
			printNodes(prefix+"\t", v.nodes)
		}
	}
}

func (m *Map) PrintNodes() {
	printNodes("", m.nodes)
}

func (m *Map) Store(key string, value int) {
	m.lock.Lock()
	defer m.lock.Unlock()

	l := len(key)
	n := m.nodes
	for _, c := range key {
		v, ok := n[c]
		if !ok {
			v = &leaf{
				nodes: make(map[rune]*leaf),
				value: -1,
			}
			n[c] = v
		}

		l--
		if l > 0 {
			n = v.nodes
		} else {
			v.value = value
		}
	}
}

func (m *Map) Load(key string) (value int, ok bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	var v *leaf
	n := m.nodes
	for _, c := range key {
		v, ok = n[c]
		if !ok {
			return
		}
		n = v.nodes
	}

	return v.value, true
}
