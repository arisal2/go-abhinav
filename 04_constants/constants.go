package main

import (
	"fmt"
	"math"
)

const s string = "constant"

// Integer Literals
// 212         /* Legal */
// 215u        /* Legal */
// 0xFeeL      /* Legal */
// 078         /* Illegal: 8 is not an octal digit */
// 032UU       /* Illegal: cannot repeat a suffix */

// 85         /* decimal */
// 0213       /* octal */
// 0x4b       /* hexadecimal */
// 30         /* int */
// 30u        /* unsigned int */
// 30l        /* long */
// 30ul       /* unsigned long */

// Floating-point Literals
// 3.14159       /* Legal */
// 314159E-5L    /* Legal */
// 510E          /* Illegal: incomplete exponent */
// 210f          /* Illegal: no decimal or exponent */
// .e55          /* Illegal: missing integer or fraction */

func main() {
	fmt.Println(s)
	const n = 50000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n)) // Sin expects float64

	const length int = 10
	const width int = 5

	var area int = length * width
	fmt.Printf("Area: %d", area)
}