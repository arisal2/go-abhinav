package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	num := 0

	/* range on arrays and slices provides both the index and value for each entry. 
	Above we didn’t need the index, so we ignored it with the blank identifier _. 
	Sometimes we actually want the indexes though.
	*/
	for _, num = range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	/* range on strings iterates over Unicode code points. 
	The first value is the starting byte index of the rune and the second the rune itself. */
	for i, c := range "go" {
        fmt.Println(i, c)
    }
}