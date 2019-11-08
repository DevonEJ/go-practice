package main

import "fmt"

func main() {

	type Person struct {
		name string
		age  int32
		pet  bool
	}

	p1 := new(Person)

	p1.name = "Devon"
	p1.age = 27
	p1.pet = false

	fmt.Println(p1.name)

	p2 := Person{name: "Lizzie", age: 29, pet: false}

	fmt.Println(p2)

}
