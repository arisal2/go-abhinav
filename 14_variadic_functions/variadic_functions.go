package main

import "fmt"

/* Variadic functions can be called with any number of trailing arguments. 
For example, fmt.Println is a common variadic function. */
func sums(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sums(1, 2)
	sums(1, 2, 3)
	nums := []int{1, 2, 3}
	sums(nums...)
}