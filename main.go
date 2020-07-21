package main

import (
	"fmt"

	"github.com/daneshvar/prj1/syncmap"
)

func main() {
	m := syncmap.New()

	m.Store("ali", 10)
	m.Store("reza", 11)
	m.Store("alireza", 10)

	v, ok := m.Load("ali")
	fmt.Printf("v:%d, ok:%t\n", v, ok)

	m.PrintNodes()
}
