package main

import "fmt"

func addInts(x, y int) int {
	res := x + y
	return res
}

func isEven(x int) (int, bool) {
	if x%2 == 0 {
		return x, true
	}
	return x, false
}

func main() {

	fmt.Println(addInts(2, 3))

	fmt.Println(isEven(5))

	fmt.Println(isEven(10))

}
