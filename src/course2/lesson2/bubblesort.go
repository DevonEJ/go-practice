package main

import "fmt"

/*
Swap takes an index and a slice of integers, and swaps the order of the value at the given index, with the value
at the next index - modifying the original slice.
*/
func Swap(numSlice []int, originalIndx int) {

	// Define the index to move the value at originaIndx to
	swapIndx := originalIndx + 1

	// Swap the order of the desired values in the original slice
	numSlice[originalIndx], numSlice[swapIndx] = numSlice[swapIndx], numSlice[originalIndx]
}

/*
BubbleSort takes a slice of integers, and applies the bubble sort algorithm to print the integers in ascending order.
*/
func BubbleSort(numbers []int) {

	// Define swaps tracking variable for later use
	var swaps bool

	// Approximate a while loop
	for {

		// Initiatlise swap var to false as default - so we know if swaps were later made on each pass through the slice or not
		swaps = false

		// Iterate over the numbers in the slice
		for indx, val := range numbers {

			// If we are at the last number in the slice - break - nothing to compare current value to
			if (indx + 1) == len(numbers) {
				break
			} else {
				// If current value is greater than the value after it then swap their order
				if val > numbers[indx+1] {

					Swap(numbers, indx)

					// Indicate that a swap took place
					swaps = true

				}
			}
		}
		// If no swaps were made in this pass of the slice, slice must be sorted, so exit loop
		if swaps == false {

			fmt.Println("Here are your sorted numbers: ", numbers)

			break
		}
	}

}

func main() {

	// Define and retrieve user input - slice of integers
	var unsortedNumbers []int

	fmt.Println("Please enter up to 10 integers separate by a space: ")

	fmt.Scan(&unsortedNumbers)

	// Apply bubble sort
	BubbleSort(unsortedNumbers)

}
