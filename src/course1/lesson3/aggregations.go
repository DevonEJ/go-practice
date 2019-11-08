package main

import "fmt"

func main() {
	// Define an array with 5 integers
	var x [5]int = [5]int{1,2,3,4,5}
	fmt.Printf("%v \n", x)

	// Iterate through the array
	for i, v := range x {
		fmt.Printf("This is the index: %d, and this is the value %d \n", i, v)
	}

	// Define array of strings
	var newArray [3]string = [3]string{"Devon", "Lizzie", "Beth"}

	// Iterate over array of strings
	for _, v := range newArray {
		fmt.Println("This is a person in the array: \n", v)
	}

	// Print slice of name array
	sliceOne := newArray[0:1]

	fmt.Printf("This is the slice: %v \n", sliceOne)

	// Example using a slice literal, instead of slicing an existing array
	sli := []string{"Devon", "Lizzie"}

	fmt.Printf("This is the slice literal: %v \n", sli)

	// Creating a slice with make()
	mk := make([]int, 2, 3)
	fmt.Println(mk)

	// Append an item to the slice and create a new one
	newSli := append(mk, 220)
	fmt.Println(newSli)
}

