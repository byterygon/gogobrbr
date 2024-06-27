package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]int)
	//Insert / update map
	m["a"] = 1
	m["a"] = 2
	m["b"] = 3

	delete(m, "a")

	_, ok := m["a"] // ok = true
	_, ok = m["A"]  // ok = false

	fmt.Println(ok)
}
