package main

import (
	"fmt"
	"time"
)

const (
	num_threads = 100
)

var sharedVariable = 0

/* 	updateTask
   	An Implementation of Mutual Exclusion using Semaphores
*/
// displays a message that is split into 2 sections to show how a rendezvous works
func updateTask(numUpdates int) {
	for i := 0; i < numUpdates; i++ {
		//UPDATE SHARED VARIABLE HERE!
		sharedVariable++
	}
}

func main() {
	var aSemaphore = make(chan int, 1)
	aSemaphore <- 1
	// Launch the threads
	for i := 0; i < num_threads; i++ {
		go func() {
			<-aSemaphore
			updateTask(1000)
			aSemaphore <- 1
		}()
	}
	fmt.Println("Launched from the main")
	time.Sleep(2 * time.Second)
	fmt.Println(sharedVariable)
}
