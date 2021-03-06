package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

/* To implement an interface in Go, we just need to implement all the methods in the interface. 
Here we implement geometry on rects. */
func (r rect) area() float64 {
	return r.height * r.width;
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

/* If a variable has an interface type, then we can call methods that are in the named interface. 
Here’s a generic measure function taking advantage of this to work on any geometry */
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area:", g.area())
	fmt.Println("perim:", g.perim())
}

func main() {
	c := rect{height: 3, width: 4}
	r := circle{radius: 5}

	/* The circle and rect struct types both implement the geometry 
	interface so we can use instances of these structs as arguments to measure. */
	measure(c)
	measure(r)
}