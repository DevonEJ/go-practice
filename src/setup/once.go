package main

import (
	"fmt"
	"sync"
)

var i int

func setupEnv() {

	fmt.Println("Initialising i to 10...")

	i = 10

}

func addStuff(wt *sync.WaitGroup, mu *sync.Mutex, on *sync.Once) {

	// Initialise i, if not already done
	on.Do(setupEnv)

	fmt.Println("Adding 5 to i...")

	// Lock variable i whilst in use
	mu.Lock()

	i = i + 5

	mu.Unlock()

	wt.Done()
}

func addMoreStuff(wt *sync.WaitGroup, mu *sync.Mutex, on *sync.Once) {

	// Initialise i, if not already done
	on.Do(setupEnv)

	fmt.Println("Adding 10 to i...")

	// Lock variable whilst in use
	mu.Lock()

	i = i + 10

	mu.Unlock()

	wt.Done()

}

func main() {

	// Set up a wait group and mutex to allow control over go routines sharing variables
	var wg sync.WaitGroup
	var mx sync.Mutex

	// Setup 'once' to allow use of init function in go routines
	var on sync.Once

	// Add go routines to wait group
	wg.Add(2)

	// Execute routines
	go addStuff(&wg, &mx, &on)
	go addMoreStuff(&wg, &mx, &on)

	wg.Wait()

	fmt.Printf("This is the value of i now: %v", i)

}
