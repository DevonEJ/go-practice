package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var inputString string
	fmt.Println("Please enter your string for checking: ")

	//fmt.Scan(&inputString)
	reader := bufio.NewReader(os.Stdin)
	inputString, _ = reader.ReadString('\n')

	// Pre-process the string so that multiple words can be entered by the user
	testString := strings.TrimSuffix(strings.ReplaceAll(strings.ToLower(inputString), " ", ""), "\n")

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
