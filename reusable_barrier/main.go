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
// Created on 28/9/2026
// Modified by:
// Issues: N/A
// Help Received: N/A
// Help Provided: N/A
//--------------------------------------------

package main

import (
	"sync"
	"time"
)

func doWorkMultipleTimes(n int, wg *sync.WaitGroup, barrier *Barrier) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		println(n, " is working, loop number: ", i)
		time.Sleep(5 * time.Second)
		println(n, " finished part A of loop number ", i)
		// we wait here
		barrier.Wait()
		println(n, " finished part B of loop number ", i)
		// we wait here for the next iteration
		barrier.WaitReuse()
	}
}

type Barrier struct {
	count         int
	startingCount int
	mu            sync.Mutex
	notify        chan struct{}
	turnstile     chan struct{}
}

func NewBarrier(count int) *Barrier {
	return &Barrier{
		count:         count,
		startingCount: count,
		notify:        make(chan struct{}),
	}
}
func (b *Barrier) Wait() {
	b.mu.Lock()
	if b.count == b.startingCount {
		b.turnstile = make(chan struct{})
	}
	b.count--
	if b.count == 0 {
		close(b.notify)
		b.count = b.startingCount
	}
	b.mu.Unlock()

	// Wait for notification
	<-b.notify
}
func (b *Barrier) WaitReuse() {
	b.mu.Lock()
	if b.count == b.startingCount {
		b.notify = make(chan struct{})
	}
	b.count--
	if b.count == 0 {
		close(b.turnstile)
		b.count = b.startingCount
	}
	b.mu.Unlock()

	// Wait for notification
	<-b.turnstile
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	var barrier = NewBarrier(5)

	for i := 0; i < 5; i++ {
		go doWorkMultipleTimes(i, &wg, barrier)
	}

	wg.Wait() //wait for everyone to finish before exiting
}
