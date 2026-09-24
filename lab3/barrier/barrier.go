//Copyright (c) 2026 Mykhailo Balaker

//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:

//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.

//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//SOFTWARE.

//--------------------------------------------
// Author: Mykhailo Balaker
// Created on 30/9/2024
// Modified by:
// Issues: N/A
// Help Received: N/A
// Help Provided: Oliwier Jakubiec
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup, barrier *Barrier) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	//we wait here until everyone has completed part A
	barrier.Wait()

	fmt.Println("Part B", goNum)
	wg.Done()
	return true
}

type Barrier struct {
	count  int
	mu     sync.Mutex
	notify chan struct{}
}

func NewBarrier(count int) *Barrier {
	return &Barrier{
		count:  count,
		notify: make(chan struct{}),
	}
}
func (b *Barrier) Wait() {
	b.mu.Lock()
	b.count--
	if b.count == 0 {
		close(b.notify)
	}
	b.mu.Unlock()

	// Wait for notification
	<-b.notify
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	var barrier = NewBarrier(totalRoutines)
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &wg, barrier)
	}

	wg.Wait() //wait for everyone to finish before exiting
}
