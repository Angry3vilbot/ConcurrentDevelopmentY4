package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Global variables shared between functions --A BAD IDEA

func addsAtomic(n int, total *atomic.Int64, wg *sync.WaitGroup) bool {
	for i := 0; i < n; i++ {
		total.Add(1)
	}
	wg.Done() //let waitgroup know we have finished
	return true
}

func main() {
	var total atomic.Int64
	var wg sync.WaitGroup

	//for loop using range option
	for i := range 10 {
		//the waitgroup is used as a barrier
		// init it to number of go routines
		wg.Add(1)
		fmt.Println("go Routine ", i)
		go addsAtomic(1000, &total, &wg)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done
	fmt.Println(total.Load())

}
