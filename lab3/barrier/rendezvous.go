// Author: Mykhailo Balaker
// License: MIT
// Provided help to: N/A
// Received help from: Oliwier Jakubiec
package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

//Global variables shared between functions --A BAD IDEA

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, chan1 chan int, chan2 chan int) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)
	if Num == 0 {
		chan1 <- 1
		<-chan2
	} else {
		chan2 <- 1
		<-chan1
	}

	fmt.Println("PartB", Num)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	threadCount := 2
	var chan1 = make(chan int, threadCount)
	var chan2 = make(chan int, threadCount)
	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, chan1, chan2)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done

}
