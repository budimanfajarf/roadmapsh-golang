package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Create slice using make() function
	var intSlice = make([]int, 10)        // when length and capacity is same
	var strSlice = make([]string, 10, 20) // when length and capacity is different

	intSlice[0] = 1
	intSlice[1] = 2

	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Println("intSlice", intSlice)

	strSlice[0] = "Hello"
	strSlice[1] = "World"

	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
	fmt.Println(reflect.ValueOf(strSlice).Kind())
	fmt.Println("strSlice", strSlice)

	// Create map using make() function
	var employee = make(map[string]int)
	employee["Mark"] = 10
	employee["Sandy"] = 20
	fmt.Println("employee", employee)

	employeeList := make(map[string]int)
	employeeList["Mark"] = 10
	employeeList["Sandy"] = 20
	fmt.Println("employeeList", employeeList)

	// Experiment - Slices of map
	var people = make([]map[string]string, 10)

	people[0] = map[string]string{
		"name": "John Doe",
		"age":  "30",
	}

	people[1] = map[string]string{
		"name":       "Budi",
		"age":        "27",
		"occupation": "Software Engineer",
	}

	fmt.Println("people", people)

	var budi = people[1]
	fmt.Println("budi", budi)
}
