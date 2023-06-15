package main

import "fmt"

func main() {
	// If we initialize the values of an array, we don't have to specify the length
	// grades := [3]int{97, 85, 93}
	// fmt.Printf("Grades: %v\n", grades)

	grades := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Budi"
	students[1] = "Lisa"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[1])

	// Note: unlike other programming languages
	// built-in len function will return the length we defined
	// not the length of the values
	fmt.Printf("Number of Students: %v\n", len(students))

	// 2D array

	// var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	// fmt.Printf("Identity Matrix: %v\n", identityMatrix)

	// // Remove redundant type of array
	var identityMatrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	fmt.Printf("Identity Matrix: %v\n", identityMatrix)

	// More readable
	var identityMatrixV2 [3][3]int
	identityMatrixV2[0] = [3]int{1, 0, 0}
	identityMatrixV2[1] = [3]int{0, 1, 0}
	identityMatrixV2[2] = [3]int{0, 0, 1}
	fmt.Printf("Identity Matrix V2: %v\n", identityMatrixV2)
}
