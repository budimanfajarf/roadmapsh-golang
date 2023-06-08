package main

import "fmt"

func main() {
	// Three-component loop
	sum := 0
	for i := 1; i < 5; i++ {
		sum++
	}
	fmt.Println(sum)

	// While loop
	sum = 1
	for sum < 1000 {
		sum++
	}
	fmt.Println(sum)

	// Infinite loop
	// sum = 0
	// for {
	// 	sum++ // repeated forever
	// }

	// For-each range loop
	words := []string{"hello", "world"}
	for index, word := range words {
		fmt.Println(index, word)
	}

	// Exit a loop
	sum = 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 { // skip odd numbers
			continue
		}
		fmt.Println(i)
	}
}
