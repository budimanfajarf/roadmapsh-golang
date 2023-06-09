package main

import "fmt"

func main() {
	/*
		Basic for-each loop (slice or array)
	*/
	words := []string{"hello", "world"}

	for index, word := range words {
		println(index, word)
	}

	/*
		String iteration: runes or bytes
	*/
	const str = "日本語"
	// const str = "budi"

	// For strings, the range loop iterates over Unicode code points.
	for index, ch := range str {
		fmt.Printf("%#U starts at byte position %d\n", ch, index)
	}

	// To loop over individual bytes, simply use a normal for loop and string indexing:
	for index := 0; index < len(str); index++ {
		fmt.Printf("%x ", str[index])
	}

	/*
		Map iteration: keys and values
	*/
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	println()

	for key, value := range numbers {
		println(key, value)
	}

	/*
		Channel iteration
	*/
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for n := range ch {
		println(n)
	}
}
