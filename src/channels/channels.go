package main

import "fmt"

//addNums adds passed in integers, sending the result to the channel
func addNums(num1 int64, num2 int64, ch chan int64) {

	res := num1 + num2

	// Send result to the channel
	ch <- res
}

func main() {

	// Create channel
	c := make(chan int64)

	// Calls go routines, passing in the channel they'll use to send data on
	go addNums(2, 4, c)
	go addNums(10, 20, c)

	// Grab results from channel
	res1 := <-c
	res2 := <-c

	fmt.Printf("The aggregated sum is: %v", (res1 + res2))
}
