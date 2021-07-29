package main

import "fmt"

func main() {
	/*
		Numbers
		-------
		int8 int16 int32 int64
		uint8 uint16 uint32 uint64
		int uint
		rune -> synonym of int32
		byte -> synonym of int8
		uintptr -> unsigned integer type
		float32 float64
		complex64 complex128
	*/

	fmt.Println("1 + 1 =", 1 + 1)
	fmt.Println("7.0/3.0 =",7.0/3.0)

	var x uint8 = 225
	var y int16 = 32767

	fmt.Println(x, x-3)
	fmt.Println(y+2, y-2)

	a := 20.45
	b := 34.89

	c := b-a

	fmt.Printf("Result is: %f", c)
	fmt.Printf("\nThe type of c is : %T", c)

	var e complex128 = complex(6, 2)
	var f complex64 = complex(9, 2)

	fmt.Println()
	fmt.Println()
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("The type of e is %T and" + "the type of f is %T", e, f)
	fmt.Println()

	// Booleans
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Printf("\nThe type is %T", true)
	fmt.Println()

	// Strings
	fmt.Println("go" + "lang")
	str := "Go Tutorial"
	
	fmt.Printf("Length: %d", len(str))
	fmt.Printf("\nValue: %s", str)
	fmt.Printf("\nType: %T", str)
}

