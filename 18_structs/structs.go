package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	/* You can safely return a pointer to local variable as a local variable will survive the scope of the function. */
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"John", 29}) //This syntax creates a new struct.
	fmt.Println(person{name: "Alice", age: 24}) //You can name the fields when initializing a struct.
	fmt.Println(person{name: "Dennis"}) //Omitted fields will be zero-valued.
	fmt.Println(&person{name: "Ann", age: 40}) //An & prefix yields a pointer to the struct.

	// It’s idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Bob"))

	// Access struct fields with a dot.
	s := person{name: "John", age: 42}
	fmt.Println(s.name)

	// You can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable
	sp.age = 43
	fmt.Println(sp.age)
}