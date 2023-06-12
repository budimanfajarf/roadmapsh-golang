package main

func main() {
	x := 20
	y := 12
	max := 15
	min := 10

	// Basic Syntax, if
	if x > max {
		x = max
	}

	println(x)

	// Basic Syntax, if else
	if x <= y {
		min = x
	} else {
		min = y
	}

	println(min)

	// With init statement
	// if x := f(); x <= y {
	// 	return x
	// }
	if i := 2; i > 2 {
		println("if")
	} else {
		println("else")
	}

	// Nested if statements
	// if x := f(); x < y {
	// 	return x
	// } else if x > z {
	// 	return z
	// } else {
	// 	return y
	// }
	y = 2
	if x := 5; x < y {
		println("if")
	} else if x > 2 {
		println("else if")
	} else {
		println("else")
	}
}
