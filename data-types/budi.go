package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// Integers
	var exampleOfInt8 int8 = 127
	fmt.Printf("\nexampleOfInt8's type is %s\n", reflect.TypeOf(exampleOfInt8))
	fmt.Printf("exampleOfInt8 is %d bytes\n", unsafe.Sizeof(exampleOfInt8))
	fmt.Printf("exampleOfInt8 is %d bits\n", unsafe.Sizeof(exampleOfInt8)*8)
	fmt.Printf("exampleOfInt8 range is (-2^7 to 2^7 -1) which is (%d to %d)\n\n", -(1 << (8 - 1)), (1<<(8-1))-1)

	var exampleOfInt16 int16 = -32768
	fmt.Printf("exampleOfInt16's type is %s\n", reflect.TypeOf(exampleOfInt16))
	fmt.Printf("exampleOfInt16 is %d bytes\n", unsafe.Sizeof(exampleOfInt16))
	fmt.Printf("exampleOfInt16 is %d bits\n", unsafe.Sizeof(exampleOfInt16)*8)
	fmt.Printf("exampleOfInt16 range is (-2^15 to 2^15 -1) which is (%d to %d)\n\n", -(1 << (16 - 1)), (1<<(16-1))-1)

	var exampleOfInt32 int32 = 2147483647
	fmt.Printf("exampleOfInt32's type is %s\n", reflect.TypeOf(exampleOfInt32))
	fmt.Printf("exampleOfInt32 is %d bytes\n", unsafe.Sizeof(exampleOfInt32))
	fmt.Printf("exampleOfInt32 is %d bits\n", unsafe.Sizeof(exampleOfInt32)*8)
	fmt.Printf("exampleOfInt32 range is (-2^31 to 2^31 -1) which is (%d to %d)\n\n", -(1 << (32 - 1)), (1<<(32-1))-1)

	var exampleOfInt64 int64 = -9223372036854775808
	fmt.Printf("exampleOfInt64's type is %s\n", reflect.TypeOf(exampleOfInt64))
	fmt.Printf("exampleOfInt64 is %d bytes\n", unsafe.Sizeof(exampleOfInt64))
	fmt.Printf("exampleOfInt64 is %d bits\n", unsafe.Sizeof(exampleOfInt64)*8)
	fmt.Printf("exampleOfInt64 range is (-2^63 to 2^63 -1) which is (%d to %d)\n\n", -(1 << (64 - 1)), (1<<(64-1))-1)

	// Unsigned Integers
	var exampleOfUint8 uint8 = 255
	fmt.Printf("exampleOfUint8's type is %s\n", reflect.TypeOf(exampleOfUint8))
	fmt.Printf("exampleOfUint8 is %d bytes\n", unsafe.Sizeof(exampleOfUint8))
	fmt.Printf("exampleOfUint8 is %d bits\n", unsafe.Sizeof(exampleOfUint8)*8)
	fmt.Printf("exampleOfUint8 range is (0 to 2^8 -1) which is (%d to %d)\n\n", 0, (1<<8)-1)

	var exampleOfUint16 uint16 = 65535
	fmt.Printf("exampleOfUint16's type is %s\n", reflect.TypeOf(exampleOfUint16))
	fmt.Printf("exampleOfUint16 is %d bytes\n", unsafe.Sizeof(exampleOfUint16))
	fmt.Printf("exampleOfUint16 is %d bits\n", unsafe.Sizeof(exampleOfUint16)*8)
	fmt.Printf("exampleOfUint16 range is (0 to 2^16 -1) which is (%d to %d)\n\n", 0, (1<<16)-1)

	var exampleOfUint32 uint32 = 4294967295
	fmt.Printf("exampleOfUint32's type is %s\n", reflect.TypeOf(exampleOfUint32))
	fmt.Printf("exampleOfUint32 is %d bytes\n", unsafe.Sizeof(exampleOfUint32))
	fmt.Printf("exampleOfUint32 is %d bits\n", unsafe.Sizeof(exampleOfUint32)*8)
	fmt.Printf("exampleOfUint32 range is (0 to 2^32 -1) which is (%d to %d)\n\n", 0, (1<<32)-1)

	var exampleOfUint64 uint64 = 18446744073709551615
	fmt.Printf("exampleOfUint64's type is %s\n", reflect.TypeOf(exampleOfUint64))
	fmt.Printf("exampleOfUint64 is %d bytes\n", unsafe.Sizeof(exampleOfUint64))
	fmt.Printf("exampleOfUint64 is %d bits\n", unsafe.Sizeof(exampleOfUint64)*8)
	fmt.Printf("exampleOfUint64 range is (0 to 2^64 -1) which is (%d to %d)\n\n", uint64(0), uint64((1<<64)-1))

	// Floats
	var exampleOfFloat32 float32 = -3.4e+38
	var exampleMinValueOfFloat32 float32 = -3.4e+38
	var exampleMaxValueOfFloat32 float32 = 3.4e+38
	fmt.Printf("exampleOfFloat32's type is %s\n", reflect.TypeOf(exampleOfFloat32))
	fmt.Printf("exampleOfFloat32 is %d bytes\n", unsafe.Sizeof(exampleOfFloat32))
	fmt.Printf("exampleOfFloat32 is %d bits\n", unsafe.Sizeof(exampleOfFloat32)*8)
	fmt.Printf("exampleOfFloat32 range is (-3.4e+38 to 3.4e+38) which is (%f to %f) \n\n", exampleMinValueOfFloat32, exampleMaxValueOfFloat32)

	var exampleOfFloat64 float64 = -1.7e+308
	var exampleMinValueOfFloat64 float64 = -1.7e+308
	var exampleMaxValueOfFloat64 float64 = 1.7e+308
	fmt.Printf("exampleOfFloat64's type is %s\n", reflect.TypeOf(exampleOfFloat64))
	fmt.Printf("exampleOfFloat64 is %d bytes\n", unsafe.Sizeof(exampleOfFloat64))
	fmt.Printf("exampleOfFloat64 is %d bits\n", unsafe.Sizeof(exampleOfFloat64)*8)
	fmt.Printf("exampleOfFloat64 range is (-1.7e+308 to 1.7e+308) which is (%f to %f) \n\n", exampleMinValueOfFloat64, exampleMaxValueOfFloat64)

	var exampleByte byte = 'a'
	fmt.Printf("exampleByte's type is %s\n", reflect.TypeOf(exampleByte))
	fmt.Printf("exampleByte is %d bytes\n", unsafe.Sizeof(exampleByte))
	fmt.Printf("exampleByte is %d bits\n", unsafe.Sizeof(exampleByte)*8)
	fmt.Printf("exampleByte character is %c\n\n", exampleByte)
}
