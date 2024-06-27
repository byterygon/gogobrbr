package main

func main() {
	//Array
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	primes := [6]int{2, 3, 5, 7, 11, 13}

	//Slice.
	// Slice thực chất là section của underlying array
	var s []int = primes[1:4] // 3, 5, 7
	s = append(s, 1)          // s: {3, 5, 7, 1} ; primes: {2, 3, 5, 7, 1, 13}

	s[0] = 1123412 // s: {1123412, 5, 7, 1} ; primes: {1123412, 3, 5, 7, 1, 13}
	//
}
