package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.LoadOrStore("a", 1)
	m.Store("c", 3)
	// m.Delete("a")
	fmt.Println(m)

	m.Range(func(key, value any) bool {
		fmt.Println(key, " -> ", value)
		return true
	})

	var m2 = map[string]int{}
	m2["a"] = 1
	m2["b"] = 2
	for k, v := range m2 {
		fmt.Println(k, "-", v)
	}
}
