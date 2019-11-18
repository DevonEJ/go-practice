package main

import (
	"fmt"
	"sync"
)

func printName(wg *sync.WaitGroup, name string) {

	fmt.Println(name)

	// Call Done to signal exit of go routine
	defer wg.Done()
}

func printLocation(wg *sync.WaitGroup, loc string) {

	fmt.Println(loc)

	defer wg.Done()
}

func main() {

	// Create a wait group
	var wg sync.WaitGroup

	// Call WG Add to block main go routine until subsequent routine exits
	wg.Add(1)

	// Call go routines - pass in wait group as first argument
	go printName(&wg, "Devon")
	// Wait until printName go routine is done
	wg.Wait()

	wg.Add(1)
	go printLocation(&wg, "London")
	wg.Wait()

	fmt.Println("Go routines complete. Exiting main go routine...")

}
