package main

import "fmt"

// Note: There is no ternary if in Go, so you'll neeed to use a full if statement even for basic conditions
func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else { 
		fmt.Println("7 is odd")
	}
}