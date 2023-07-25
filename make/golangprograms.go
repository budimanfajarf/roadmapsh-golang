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
}
