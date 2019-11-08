package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Ask user for filename
	fmt.Println("Please enter the filename (Hint: it's file.txt): ")
	var fileName string
	fmt.Scan(&fileName)

	// Read info from file
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	// Define struct for person
	type Person struct {
		fname string
		lname string
	}

	// Initiatlise the slice to hold the Person structs
	personSlice := make([]Person, 0)

	// Convert byte array to string
	s := string(data)

	// Iterate over the string to get the individuals' names
	scanner := bufio.NewScanner(strings.NewReader(s))

	for scanner.Scan() {
		// Split string at the space to separate first and last names for each line
		test := strings.Split(scanner.Text(), " ")

		// Add fname and lname to a new struct and append it to the personSlice
		personSlice = append(personSlice, Person{fname: test[0], lname: test[1]})
	}

	for _, person := range personSlice {
		fmt.Printf("This is the current person's first name: %v, and last: %v \n", person.fname, person.lname)
	}

	fmt.Println("That's everyone. End.")

}
