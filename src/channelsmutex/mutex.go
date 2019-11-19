package main

import (
	"fmt"
	"sync"
)

var i int64 = 0

func incrNum(wg *sync.WaitGroup, mut *sync.Mutex) {
	// Use Mutex to lock te i variable whilst in use in current go routine
	mut.Lock()

	i = i + 1

	// Unlock i variable after use
	mut.Unlock()

	wg.Done()
}

func main() {

	// Create wait group
	var wg sync.WaitGroup

	// Create mutex object - because all go routines sharing the i variable
	var mut sync.Mutex

	// call incrNum twice with two new go routines
	wg.Add(2)

	go incrNum(&wg, &mut)

	go incrNum(&wg, &mut)

	wg.Wait()

	// Print i
	fmt.Printf("This is i now: %v", i)

}
