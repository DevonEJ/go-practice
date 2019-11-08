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

//BubbleSort applies bubble sort using swapping to a slice of integers, printing the sorted slice's values
func BubbleSort(numbers []int) {

	// Initiatlise swap slice - so we know if swaps were made on each pass through the slice or not
	noSwaps := false

	// Approximate a while loop
	for {

		// If noSwaps is false, continue sorting
		if noSwaps == false {

			// Iterate over the numbers
			for indx, val := range numbers {
				// If we are at the last number in the slice
				if (indx + 1) == len(numbers) {
					break
				}
				// If current value is greater than the value after it...
				if val > numbers[indx+1] {

					Swap(numbers, indx)
					// Set swaps tracker to true, to indicate
					noSwaps = false

				} else {
					noSwaps = true
				}

			}
		}
		// If no more swaps needed to sort, print sorted slice and exit loop
		fmt.Println(numbers)
		break

	}

}

func main() {

	unsortedNumbers := []int{4, 1, 3, 7, 2}

	BubbleSort(unsortedNumbers)

}
