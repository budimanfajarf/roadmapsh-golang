package main

import "fmt"

func main() {
	// Create a channel with make

	// Create channel of integer type
	number := make(chan int)

	// Access type and value of channel
	fmt.Printf("Channel Type: %T\n", number)
	fmt.Printf("Channel Value: %v\n", number)
}
