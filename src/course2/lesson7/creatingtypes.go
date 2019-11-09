package main

import "fmt"

// Define a new type called SpecialInt
type SpecialInt int

// Create a method for SpecialInt
func (si SpecialInt) Increment() int {
	return int(si + 1)
}

func main() {

	// Create instance of SpecialInt
	val := SpecialInt(5)

	// Apply the Increment method to it - using dot notation
	fmt.Println(val.Increment())

}
