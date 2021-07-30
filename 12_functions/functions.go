package main

import "fmt"

/* Go requires explicit returns, i.e. it won’t automatically return the value of the last expression. */

func add(a int, b int) int {
	return a + b
}

func addMore(a int, b int, c int) int {
	return a + b
}

func main() {
	sum := add(1, 2)
	fmt.Println("sum:", sum)
	sum = addMore(1, 2, 3)
	fmt.Println("sum:", sum)
}