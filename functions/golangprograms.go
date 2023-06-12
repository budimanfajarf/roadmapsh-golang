package main

// Simple function to print a message
func SimpleFunction() {
	println("Simple Function")
}

// Function accepting parameters
func add(x int, y int) {
	total := x + y
	println(total)
}

// Function with int as return type
func multiply(x int, y int) int {
	return x * y
}

// The return value of a function can be named in Golang
func rectangle(l int, b int) (area int) {
	var parameter int
	parameter = 2 * (l + b)
	println("Parameter: ", parameter)

	area = l * b
	return // Return statement without specify variable name
}

// Golang Functions Returning Multiple Values
func rentangleV2(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}

func main() {
	SimpleFunction()

	add(10, 20)

	println(multiply(10, 20))

	println("Area: ", rectangle(10, 20))

	area, parameter := rentangleV2(10, 20)
	println("Area: ", area)
	println("Parameter: ", parameter)
}
