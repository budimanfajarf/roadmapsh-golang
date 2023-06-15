package main

import "fmt"

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

func main() {
	// https://go.dev/tour/moretypes/6
	// An array's length is part of its type, so arrays cannot be resized.
	// This seems limiting, but don't worry; Go provides a convenient way of working with arrays.

	// https://go.dev/doc/effective_go#arrays
	// There are major differences between the ways arrays work in Go and C. In Go,
	// - Arrays are values. Assigning one array to another copies all the elements.
	// - In particular, if you pass an array to a function, it will receive a copy of the array, not a pointer to it.
	// - The size of an array is part of its type. The types [10]int and [20]int are distinct.

	// The value property can be useful but also expensive; if you want C-like behavior and efficiency, you can pass a pointer to the array.
	// numbers := [3]float64{7.0, 8.5, 9.1}
	// numbers := [...]float64{7.0, 8.5, 9.1}
	numbers := [3]float64{7.0, 8.5}
	sum := Sum(&numbers) // Note the explicit address-of operator
	fmt.Printf("Sum: %f\n", sum)
}
