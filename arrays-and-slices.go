package main

import "fmt"

func main() {
	exampleArray := [...]int{1, 2, 3, 4, 5}

	copyOfExampleArray := exampleArray
	copyOfExampleArray[0] = 100
	fmt.Println("exampleArray", exampleArray)
	fmt.Println("copyOfExampleArray", copyOfExampleArray)
	fmt.Println("exampleArray[0]", exampleArray[0])

	fmt.Println()

	exampleSlice := []int{1, 2, 3, 4, 5}

	copyOfExampleSlice := exampleSlice
	copyOfExampleSlice[0] = 100
	fmt.Println("exampleSlice", exampleSlice)
	fmt.Println("copyOfExampleSlice", copyOfExampleSlice)
	fmt.Println("exampleSlice[0]", exampleSlice[0])

	fmt.Println()

	pointedExampleArray := &exampleArray
	pointedExampleArray[0] = 100
	fmt.Println("exampleArray", exampleArray)
	fmt.Println("pointedExampleArray", pointedExampleArray)
	fmt.Println("exampleArray[0]", exampleArray[0])
}
