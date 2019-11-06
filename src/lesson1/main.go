package main

import "fmt"

func main() {
	fmt.Printf("Hello world. \n")

	for i := 1; i < 10; i++ {
		fmt.Printf("i is currently %v \n", i)
	}

	b := 1
	for  b < 20 {
		b++
		if b == 5 { break }
		fmt.Printf("Not yet 5. \n")
	}

}

