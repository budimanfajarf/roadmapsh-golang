package main

import "fmt"

func main() {
	// https://go.dev/tour/basics/14
	// When declaring a variable without specifying an explicit type
	// (either by using the := syntax or var = expression syntax),
	// the variable's type is inferred from the value on the right hand side.

	// https://www.callicoder.com/golang-variables-zero-values-type-inference/#type-inference
	// Although Go is a Statically typed language,
	// It doesn’t require you to explicitly specify the type of every variable you declare.

	// Type inference

	var name = "Budi"
	fmt.Printf("Variable 'name' is of type %T\n", name)

	// name = 1234 // Compiler Error

	// Multiple variable declarations with inferred types
	var firstName, lastName, age, salary = "Budiman Fajar", "Firdaus", 28, 50000.0
	fmt.Printf("firstName: %T, lastName: %T, age: %T, salary: %T\n", firstName, lastName, age, salary)

	// Short Declaration

	// Short variable declaration syntax
	name2 := "Budiman Fajar"
	age, salary, isProgrammer := 35, 50000.0, true

	fmt.Printf("name2: %T, age: %T, salary: %T, isProgrammer: %T\n", name2, age, salary, isProgrammer)
}
