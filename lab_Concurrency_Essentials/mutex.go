// Author: Mykhailo Balaker
// License: MIT
// Provided help to: N/A
// Received help from: N/A
package main

import (
	"fmt"
	"sync"
)

func adds(n int, theLock *sync.Mutex, total *int64, wg *sync.WaitGroup) bool {
	for i := 0; i < n; i++ {
		theLock.Lock()
		*total++
		theLock.Unlock()
	}
	wg.Done() //let waitgroup know we have finished
	return true
}

func main() {
	var wg sync.WaitGroup
	var total int64
	//theLock will be passed by reference between go routines
	//better than using a global variable
	var theLock sync.Mutex

	total = 0
	//the waitgroup is used as a barrier
	// init it to number of go routines
	wg.Add(10)

	//for loop using range option
	for i := range 10 {
		//starting
		fmt.Println(i)
		go adds(1000, &theLock, &total, &wg)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done
	fmt.Println(total)
}
