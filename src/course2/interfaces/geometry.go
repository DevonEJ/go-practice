package main

import (
	"fmt"
	"math"
)

// Geometry interface requires an area() and perim() method for each type implementing it
type geometry interface {
	area() float64
	perim() float64
}

// Define 2 types - circle and rectangle
type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

// Define the methods for each type - each has a different implementation of the area() method
// Area methods
func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Perimeter methods
func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Helper functions take the interface as argument - making use of the method that is defined for the correct type
func measureArea(g geometry) {
	// Call area on whichever type is being operated on
	fmt.Println(g.area())
}

func measurePerim(g geometry) {
	fmt.Println(g.perim())
}

// Main
func main() {

	// Define instances of the shape types
	rec := rectangle{width: 3, height: 4}
	circ := circle{radius: 10}

	// Apply measure helper function to the rec instance
	measureArea(rec)

	// Apply measure helper function to the rec instance
	measurePerim(circ)

}
