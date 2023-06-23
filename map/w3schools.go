package main

import "fmt"

func main() {
	// https://www.w3schools.com/go/go_maps.php

	// Creating Maps Using var and :=
	fmt.Println("Creating Maps Using var and :=")

	var car = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	cityMap := map[string]int{"New York": 8336697, "Los Angeles": 3857799, "Chicago": 2714856}

	fmt.Printf("car\t%v\n", car)
	fmt.Printf("cityMap\t%v\n", cityMap)

	// Creating Maps Using Using make()Function:
	fmt.Println("\nCreating Maps Using Using make()Function:")

	var car2 = make(map[string]string)
	car2["brand"] = "Ford"
	car2["model"] = "Mustang"
	car2["year"] = "1964"

	cityMap2 := make(map[string]int)
	cityMap2["New York"] = 8336697
	cityMap2["Los Angeles"] = 3857799
	cityMap2["Chicago"] = 2714856

	fmt.Printf("car2\t%v\n", car2)
	fmt.Printf("cityMap2\t%v\n", cityMap2)

	// Creating an Empty Map
	fmt.Println("\nCreating an Empty Map")

	var car3 = make(map[string]string)
	var cityMap3 map[string]int

	fmt.Printf("car3\t%v\n", car3)
	fmt.Println("car3 == nil:", car3 == nil) // car3 == nil: false
	fmt.Printf("cityMap3\t%v\n", cityMap3)
	fmt.Println("cityMap3 == nil:", cityMap3 == nil) // cityMap3 == nil: true

	// Accessing Map Elements
	fmt.Println("\nAccessing Map Elements")

	fmt.Println("car[\"brand\"]:", car["brand"]) // car["brand"]: Ford

	// Updating Map Elements
	fmt.Println("\nUpdating Map Elements")

	car["year"] = "2020"
	fmt.Println("car[\"year\"]:", car["year"]) // car["year"]: 2020

	// Adding Items to a Map
	fmt.Println("\nAdding Items to a Map")

	car["color"] = "red"
	fmt.Println("car[\"color\"]:", car["color"]) // car["color"]: red
	fmt.Printf("car\t%v\n", car)                 // car	map[brand:Ford color:red model:Mustang year:2020]

	// Deleting Items Using delete() Function from a Map
	fmt.Println("\nDeleting Items Using delete() Function from a Map")

	delete(car, "color")
	fmt.Printf("car\t%v\n", car) // car	map[brand:Ford model:Mustang year:2020]

	// Check For Specific Elements in a Map
	fmt.Println("\nCheck For Specific Elements in a Map")

	var car4 = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}

	val1, ok1 := car4["brand"] // Checking for existing key and its value
	val2, ok2 := car4["color"] // Checking for non-existing key and its value
	val3, ok3 := car4["day"]   // Checking for existing key and its value
	_, ok4 := car4["model"]    // Only checking for existing key and not its value
	_, ok5 := car4["color"]    // Only checking for non-existing key and not its value

	fmt.Println("val1:", val1, "ok1:", ok1) // val1: Ford ok1: true
	fmt.Println("val2:", val2, "ok2:", ok2) // val2:  ok2: false
	fmt.Println("val3:", val3, "ok3:", ok3) // val3:  ok3: true
	fmt.Println("ok4:", ok4)                // ok4: true
	fmt.Println("ok5:", ok5)                // ok5: false

	// Maps Are References
	fmt.Println("\nMaps Are References")

	var car5 = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	car6 := car5

	fmt.Printf("car5\t%v\n", car5) // car5	map[brand:Ford model:Mustang year:1964]
	fmt.Printf("car6\t%v\n", car6) // car6	map[brand:Ford model:Mustang year:1964]

	car6["year"] = "2020"
	fmt.Printf("car5\t%v\n", car5) // car5	map[brand:Ford model:Mustang year:2020]
	fmt.Printf("car6\t%v\n", car6) // car6	map[brand:Ford model:Mustang year:2020]

	// Iterating Over a Map Using for range
	fmt.Println("\nIterating Over a Map Using for range")

	for key, value := range car {
		fmt.Printf("key: %v, value: %v\n", key, value)
	}
}
