package main

import (
	"fmt"
)

/*
"GenDisplaceFn computes displacement as a function of time (t),
assuming the given values acceleration (a), initial velocity (vo), and initial displacement (so)."
*/
func GenDisplaceFn(a, vo, so float64) func(t float64) float64 {

	timeFunc := func(t float64) float64 {

		res := (a/2)*(t*t) + (vo * t) + so

		return res

	}
	return timeFunc

}

func main() {

	// Define input vars
	var acc float64
	var vel float64
	var dis float64
	var secs float64

	// Get user to enter acceleration, initial velocity, and initial displacement
	fmt.Println("Please enter the acceleration value in format 00.00: ")
	fmt.Scanf("%f", &acc)

	fmt.Println("Please enter the initial velocity value in format 00.00: ")
	fmt.Scanf("%f", &vel)

	fmt.Println("Please enter the initial displacement value in format 00.00: ")
	fmt.Scanf("%f", &dis)

	// Call GenFunc, passing in the user's input
	GFunc := GenDisplaceFn(acc, vel, dis)

	// Get the time displacement value from the user
	fmt.Println("Please enter the value of time, in seconds, in format 00.00 (E.g. 00.03 = 3 seconds): ")
	fmt.Scanf("%f", &secs)

	// Use user's time input value to displace the calculation
	fmt.Printf("The value after the time displacement you specified is %v", GFunc(secs))

}
