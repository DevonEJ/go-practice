package main

import "fmt"

//getMaxNum takes a variable length slice of integers, and returns the max integer in that slice
func getMaxNum(vals ...int) int {

	maxV := -1

	for _, num := range vals {
		if num > maxV {
			maxV = num
		}

	}
	return maxV
}

func main() {

	// Can call with a comma sep list of integers, like this...
	fmt.Println(getMaxNum(3, 200, 3, 6, 55))

	valSlice := []int{3, 5, 2, 99, 7, 2, 8}

	// Or - the ... denotes that a slice has been passed instead of a comma sep list of integers
	fmt.Println(getMaxNum(valSlice...))

}
