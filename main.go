package main

import "fmt"

type Nodes map[rune]*leaf

type leaf struct {
	value int
	nodes Nodes
}

type Map struct {
	nodes Nodes
}

func main() {
	m := Map{
		nodes: make(map[rune]*leaf),
	}

	m.Store("ali", 10)
	m.Store("reza", 11)
	m.Store("alireza", 10)
	PrintNodes("", m.nodes)
}

func PrintNodes(prefix string, nodes Nodes) {
	for k, v := range nodes {
		fmt.Printf("%s%s:%d\n", prefix, string(k), v.value)
		if v.nodes != nil {
			PrintNodes(prefix+"\t", v.nodes)
		}
	}
}

func (m *Map) Store(key string, value int) {
	l := len(key)
	n := m.nodes
	for _, c := range key {
		v, ok := n[c]
		if !ok {
			v = &leaf{
				nodes: make(map[rune]*leaf),
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

func (m *Map) Load(key string) int {
	return 0
}
