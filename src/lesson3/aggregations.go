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
}

