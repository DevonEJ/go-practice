package bunnies

import "fmt"

// Create type Bunny containing private data, not to be accessed by importing main package
type Bunny struct {
	name        string
	babyBunnies int
}

// The use of pointer receivers for methods of the Bunny type allow for modification of the underlying Bunny when calling
func (bn *Bunny) CreateBunny() {

}

// Define public method for Bunny receiver type
func (bn *Bunny) NameBunny(name string) {
	bn.name = name
}

func (bn *Bunny) AddBabyBunnies(babies int) {
	bn.babyBunnies = babies
}

func (bn *Bunny) GreetBunny() {
	fmt.Printf("Hello, bunny called %v!", bn.name)
}

func (bn *Bunny) MatingSeason() {
	bn.babyBunnies = bn.babyBunnies * 2
	fmt.Printf("%v now has %v baby bunnies!", bn.name, bn.babyBunnies)
}
