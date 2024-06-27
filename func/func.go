package main

import "fmt"

func add(a, b float64) float64 {
	return a + b
}
func calc(a, b float64) (float64, float64) {
	return a + b, a * b
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	result := compute(add)
	fmt.Println(result)
}
