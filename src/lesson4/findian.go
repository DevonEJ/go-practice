package main

import (
	"fmt"
	"strings"
)

func main() {

	var inputString string
	fmt.Println("Please enter your string for checking: ")

	fmt.Scan(&inputString)

	testString := strings.ToLower(inputString)

	// Check for the required letters
	containsA := strings.Contains(testString, "a")

	endsN := strings.HasSuffix(testString, "n")

	startsI := strings.HasPrefix(testString, "i")

	if startsI && containsA && endsN {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}

}
