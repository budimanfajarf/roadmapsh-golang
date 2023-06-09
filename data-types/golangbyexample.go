package main

import (
	"fmt"
	"math/bits"
	"reflect"
	"unsafe"
)

func main() {
	//This is computed as const uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	sizeOfIntInBits := bits.UintSize
	fmt.Printf("%d bits\n", sizeOfIntInBits)

	var a int
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))

	b := 2
	fmt.Printf("b's typs is %s\n", reflect.TypeOf(b))

	//This is computed as const uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	sizeOfUintInBits := bits.UintSize
	fmt.Printf("%d bits\n", sizeOfUintInBits)

	/*
		Byte
		- byte in Go is an alias for uint8 meaning it is an integer value
		- This integer value is of 8 bits and it represents one byte i.e number between 0-255
		- byte is used to represent ASCII characters
	*/
	var r byte = 'a'

	//Print Size
	fmt.Printf("Size: %d\n", unsafe.Sizeof(r))

	//Print Type
	fmt.Printf("Type: %s\n", reflect.TypeOf(r))

	//Print Character
	fmt.Printf("Character: %c\n", r)
	s := "abc"

	//This will the decimal value of byte
	fmt.Println([]byte(s))

	/*
		Rune
		- rune in Go is an alias for int32 meaning it is an integer value
		- This integer value is meant to represent a Unicode Code Point
		- Article about Unicode joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/
		- List of Unicode characters https://en.wikipedia.org/wiki/List_of_Unicode_characters
	*/

	fmt.Printf("%U\n", []rune("0b£"))

	rPound := '£'
	fmt.Printf("Type: %s\n", reflect.TypeOf(rPound))
	fmt.Printf("Unicode CodePoint: %U\n", rPound)
	fmt.Printf("Character: %c\n", rPound)

	theRune := 'a'

	//Print Size
	fmt.Printf("Size: %d\n", unsafe.Sizeof(theRune))

	//Print Type
	fmt.Printf("Type: %s\n", reflect.TypeOf(theRune))

	//Print Code Point
	fmt.Printf("Unicode CodePoint: %U\n", theRune)

	//Print Character
	fmt.Printf("Character: %c\n", theRune)

	ss := "0b£"

	//This will print the Unicode Points
	fmt.Printf("%U\n", []rune(ss))

	//This will the decimal value of Unicode Code Point
	fmt.Println([]rune(ss))
}
