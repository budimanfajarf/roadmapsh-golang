package main

import "fmt"

func main() {
	// Slices it looks like arrays
	// but it doesn't have a fixed size, it's dynamic, so we don't have to specify the length
	a := []int{1, 2, 3}
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	b := a
	b[1] = 5

	// Note: unlike arrays, when we assign a slice to another variable
	// it will create a REFERENCE to the slice
	fmt.Println(a)
	fmt.Println(b)
}
