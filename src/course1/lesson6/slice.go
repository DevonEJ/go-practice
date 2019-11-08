package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

/*
Programme requests numbers from the user, adding these to a slice, until the user enters 'x' or 'X' to quit
the programme.
Then the slice is returned, sorted.
*/
func main() {

	// Define slice variable
	numSlice := make([]int, 0)

	for {

		fmt.Println("Please enter a number - or enter X to quit: ")

		// Get initial user integer input
		var userInput string
		_, err := fmt.Scan(&userInput)

		// Check for errors
		if err != nil {
			log.Fatal(err)
		}

		// Check if user wants to quit
		if userInput == "x" || userInput == "X" {
			if len(numSlice) > 0 {
				sort.Ints(numSlice)
				fmt.Println("You have chosen to quit. The numbers you entered were: ", numSlice)
			} else {
				fmt.Println("You have chosen to quit. You entered no numbers.")
			}
			break
		} else {
			// Convert input to integer
			newInput, err2 := strconv.Atoi(userInput)

			if err2 == nil {
				// Add to the slice
				numSlice = append(numSlice, newInput)

			} else {
				fmt.Println("There was an error. Please make sure you enter a number or x to quit.")
				log.Fatal(err2)
			}

		}
	}
}
