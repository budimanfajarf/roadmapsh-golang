package main

func main() {
	// Normal loop
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	println(sum)

	// Looping over an array, slice, string, or map, or reading from a channel, a range clause can manage the loop.

	names := []string{"Budi", "Susi", "Joko"}

	// Use key & value
	for key, value := range names {
		println(key, value)
	}

	// Use key only
	for key := range names {
		println(key)
	}

	// Use value only
	for _, value := range names {
		println(value)
	}
}
