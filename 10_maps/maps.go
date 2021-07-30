package main

import "fmt"

/* To create an empty map, use the builtin make: make(map[key-type]val-type).
Set key/value pairs using typical name[key] = val syntax. */

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m))

	/* The builtin delete removes key/value pairs from a map. */
	delete(m, "k2")
	fmt.Println("map:", m)

	/* In this statement, the first value (i) is assigned the value stored under the key "route". 
	If that key doesn't exist, i is the value type's zero value (0). 
	The second value (k1) is a bool that is true if the key exists in the map, and false if not. */
	i, k1 := m["k1"]
	fmt.Println("i:", i)
	fmt.Println("k1:", k1)

	/*  The optional second return value when getting a value from a map indicates if the key was present in the map. 
	This can be used to disambiguate between missing keys and keys with zero values like 0 or "". 
	Here we didn’t need the value itself, so we ignored it with the blank identifier _. */
	_, k2 := m["k2"]
	fmt.Println("k2:", k2)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:",n)

	/* 
		To iterate over the contents of a map, use the range keyword.
		Range will be discussed more in the next section.
	*/
	for key, value := range n {
		fmt.Println("Key:", key, "Value:", value)
	}
}