package main

import "time"

func main() {
	switch today := time.Now(); {
	case today.Day() < 5:
		println("Clean your house.")
	case today.Day() < 10:
		println("Buy some coffee.")
	case today.Day() > 15:
		println("Visit a doctor.")
	case today.Day() == 25:
		println("Buy some food.")
	default:
		println("No information available for that day.")
	}
}
