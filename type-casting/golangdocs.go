package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// Basic Type-casting
	var a int = 42
	b := float64(a)
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))

	// Conversions between string and int
	var s string = "42"
	v, _ := strconv.Atoi(s)
	fmt.Println(v)
	fmt.Println(reflect.TypeOf(v))

	var i int = 42
	str := strconv.Itoa(i)
	fmt.Println(str)
	fmt.Println(reflect.TypeOf(str))

	// Conversions between int and float
	ff := 3.1415926
	ii := int(ff) // loses precision
	fmt.Println(ii)
	fmt.Println(reflect.TypeOf(ii))

	iii := 3
	fff := float64(iii)
	fmt.Println(fff)
	fmt.Println(reflect.TypeOf(fff))

	// Strings and bytes conversion
	var string1 string = "Hello World"
	var bytes1 []byte = []byte(string1) // convert ty bytes
	fmt.Println(bytes1)
	fmt.Println(reflect.TypeOf(bytes1))

	string2 := string(bytes1) // convert to string
	fmt.Println(string2)
	fmt.Println(reflect.TypeOf(string2))

	char0 := string(string1[0]) // get the first character
	fmt.Println(char0)
	fmt.Println(reflect.TypeOf(char0))

	// Type conversion during division
	divResult1 := 6 / 3 // both are int, divResult1 is int
	fmt.Println(divResult1)
	fmt.Println(reflect.TypeOf(divResult1))

	divResult2 := 6 / 3.0 // float and int, divResult2 is float
	fmt.Println(divResult2)
	fmt.Println(reflect.TypeOf(divResult2))
}
