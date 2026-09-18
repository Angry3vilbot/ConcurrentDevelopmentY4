package main

import (
	"fmt"
	"time"
)

// Displays a message first
func taskOne(delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Print("I ")
	fmt.Print("must ")
	fmt.Print("print ")
	fmt.Println("first")
	// Tell taskTwo to start now
	semaphore <- 1
}

// Displays a message second
func taskTwo() {
	<-semaphore
	fmt.Print("This ")
	fmt.Print("will ")
	time.Sleep(5 * time.Second)
	fmt.Print("appear ")
	fmt.Println("second")
}

var semaphore = make(chan int)

func main() {
	taskOneDelay := 5
	go taskOne(taskOneDelay)
	go taskTwo()
	fmt.Print("Launched from the main\n")
	time.Sleep(15 * time.Second)
}
