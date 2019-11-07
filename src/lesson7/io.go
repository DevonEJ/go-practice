package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	// Get user input for name and address
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your name: ")
	inputName, err1 := reader.ReadString('\n')

	fmt.Println("Please enter your address: ")
	inputAddr, err2 := reader.ReadString('\n')

	if err1 != nil && err2 != nil {
		fmt.Println("There was an error - please try again.")
	}

	// Create map for storing user's details
	personMap := make(map[string]string)

	// Add details
	personMap["name"] = inputName
	personMap["address"] = inputAddr

	// Convert map to JSON object
	bObject, err := json.Marshal(personMap)

	if err != nil {
		fmt.Println("There was an error converting map to JSON object.")
	}

	fmt.Printf("Here is the JSON object: \n %v", bObject)

}
