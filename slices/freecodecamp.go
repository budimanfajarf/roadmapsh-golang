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
	// When we change `b` it will also change `a`
	fmt.Println(a)
	fmt.Println(b)

	// data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// sliceOfAllElements := data[:]             // slice of all elements
	// sliceFrom4thToEnd := data[3:]             // slice from 4th element to end
	// sliceFirst6Elements := data[:6]           // slice first 6 elements
	// sliceThe4th5thAnd6thElements := data[3:6] // slice the 4th, 5th and 6th elements
	// fmt.Println("data", data)
	// fmt.Println("sliceOfAllElements", sliceOfAllElements)
	// fmt.Println("sliceFrom4thToEnd", sliceFrom4thToEnd)
	// fmt.Println("sliceFirst6Elements", sliceFirst6Elements)
	// fmt.Println("sliceThe4th5thAnd6thElements", sliceThe4th5thAnd6thElements)

	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Note: index start from 0
	sliceOfAllElements := data[:]                   // slice of all elements
	sliceFromIndex3ToEnd := data[3:]                // slice from index 3 to end
	sliceFirst6Elements := data[:6]                 // slice first 6 elements
	sliceFromIndex3ToIndex5 := data[3:6]            // slice from index 3 to index 5

	data[4] = 44                    // It will change all the values that reference to the data
	sliceFromIndex3ToIndex5[0] = 33 // It also will change all the values that reference to the data

	fmt.Println("data", data)
	fmt.Println("sliceOfAllElements", sliceOfAllElements)
	fmt.Println("sliceFromIndex3ToEnd", sliceFromIndex3ToEnd)
	fmt.Println("sliceFirst6Elements", sliceFirst6Elements)
	fmt.Println("sliceFromIndex3ToIndex5", sliceFromIndex3ToIndex5)
}
