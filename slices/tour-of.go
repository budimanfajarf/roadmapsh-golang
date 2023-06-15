package main

import "fmt"

func main() {
	// An array has a fixed size.
	// A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.
	// In practice, slices are much more common than arrays.

	primes := [6]int{2, 3, 5, 7, 11, 13}

	// A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	// a[low : high]

	// This selects a half-open range which includes the first element, but excludes the last one.
	// The following expression creates a slice which includes elements 1 through 3 of a:
	// a[1:4]
	var s []int = primes[1:4]
	println(s)
	fmt.Println(s)

	// Note: println and fmt.Println are different, check the terminal output
}
