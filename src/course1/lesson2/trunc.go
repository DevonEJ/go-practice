package main

import (
	"fmt";
	"log"
)


func main() {

	// Takes user input and returns it as an integer - or returns error

	fmt.Println("Please enter a floating point number:")

	var userInputAddr float64

	_, err := fmt.Scan(&userInputAddr)

	if err == nil {
		fmt.Println(int64(userInputAddr))
	} else {
		log.Fatal(err)
	}

}