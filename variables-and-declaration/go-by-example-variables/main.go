package main

import "fmt"

func main() {
	// declare one variable
	var a = "initial"
	fmt.Println(a)

	var d = true
	fmt.Println(d)

	// declare multiple variables
	var b, c int = 1, 2
	fmt.Println(b, c)

	// declare variable without initial value will have "zero-valued", e.g 0 for int
	var e int
	fmt.Println(e)

	// shorthand for declaring and initializing a variable
	f := "shorthand"
	fmt.Println(f)
}
