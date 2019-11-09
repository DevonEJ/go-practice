package main

import "fmt"

func printEndStart() {

	// Printed third
	defer fmt.Println("End.")

	// Printed second
	defer fmt.Println("Do some stuff.")

	// Printed first
	fmt.Println("Start.")
}

func printNum(num int) {

	defer fmt.Printf("This is the number printed in a deferred statement: %v", num)

}

func main() {
	printEndStart()

	number := 7

	/*
	 This will print the number 9 (not 10) - because the argument number is evaluated (as 7) before it is incremented
	 on line 31, even though the function printNum contains a deferred Println call.

	 Execution of print is deffered, but eval of the argument is not.
	*/
	printNum(number + 2)

	number++

	fmt.Println(number)

}
