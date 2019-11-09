package main

import "fmt"

// Using functions as variables

var funcVar func(int)

//printr prints a integer passed in
func printr(x int) {
	fmt.Println(x)
}

//runFuncs executes a function passed in as parameter, with the argument val
func runFuncs(funcs func(int), val int) {
	funcs(val)
}

func main() {
	// Assign a function to the variable funcVar
	funcVar = printr
	// Execute the function that funcVar now points to
	funcVar(7)

	/*
	 Pass runFuncs a function to execute - in this case, using the name of the variable that points to it,
	 instead of the function's defined name
	*/
	runFuncs(funcVar, 8)

	// runFunc using an anonymous (lambda) function
	runFuncs(func(x int) { fmt.Println(x * 2) }, 2)

}
