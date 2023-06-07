package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var exampleOfInt8 int8 = 127
	fmt.Printf("exampleOfInt8's type is %s\n", reflect.TypeOf(exampleOfInt8))
	fmt.Printf("exampleOfInt8 is %d bytes\n", unsafe.Sizeof(exampleOfInt8))
	fmt.Printf("exampleOfInt8 is %d bits\n", unsafe.Sizeof(exampleOfInt8)*8)
	fmt.Printf("exampleOfInt8 range is from %d to %d\n", -(1 << (8 - 1)), (1<<(8-1))-1)

	var exampleOfInt16 int16 = -32768
	fmt.Printf("exampleOfInt16's type is %s\n", reflect.TypeOf(exampleOfInt16))
	fmt.Printf("exampleOfInt16 is %d bytes\n", unsafe.Sizeof(exampleOfInt16))
	fmt.Printf("exampleOfInt16 is %d bits\n", unsafe.Sizeof(exampleOfInt16)*8)
	fmt.Printf("exampleOfInt16 range is from %d to %d\n", -(1 << (16 - 1)), (1<<(16-1))-1)

	var exampleOfInt32 int32 = 2147483647
	fmt.Printf("exampleOfInt32's type is %s\n", reflect.TypeOf(exampleOfInt32))
	fmt.Printf("exampleOfInt32 is %d bytes\n", unsafe.Sizeof(exampleOfInt32))
	fmt.Printf("exampleOfInt32 is %d bits\n", unsafe.Sizeof(exampleOfInt32)*8)
	fmt.Printf("exampleOfInt32 range is from %d to %d\n", -(1 << (32 - 1)), (1<<(32-1))-1)
}
