package main

import "fmt"

/* The (int, int) in this function signature shows that the function returns 2 ints. */
func vals() (int, int) {
	return 3,7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	d, _ := vals()
	fmt.Println(d)
}