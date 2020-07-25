package main

import (
	"fmt"

	"github.com/daneshvar/prj1/syncmap"
)

func main() {
	m := syncmap.New()

	m.Store(1000, 10)
	m.Store(1100, 11)
	m.Store(7, 10)

	v, ok := m.Load(1000)
	fmt.Printf("v:%d, ok:%t\n", v, ok)
}
