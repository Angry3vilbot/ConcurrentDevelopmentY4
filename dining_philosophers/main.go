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
// Created on 24/9/2026
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

func eat(wg *sync.WaitGroup, fm chan struct{}, forks map[int]chan int, index int) {
	defer wg.Done()

	fm <- struct{}{}
	println(index+1, " wants to eat.")
	forks[index] <- 1
	println(index+1, " took fork ", index+1)
	forks[(index+1)%len(forks)] <- 1
	println(index+1, " took fork ", (index+1)%len(forks)+1)
	println(index+1, " eating...")

	time.Sleep(time.Second * 5)

	<-forks[index]
	<-forks[(index+1)%len(forks)]
	<-fm
	println(index+1, " done!")
}

func main() {
	waitGroup := sync.WaitGroup{}
	numberOfThreads := 5
	footMan := make(chan struct{}, numberOfThreads-1)
	forks := make(map[int]chan int)

	waitGroup.Add(numberOfThreads)
	for i := 0; i < numberOfThreads; i++ {
		forks[i] = make(chan int, 1)
	}
	for i := 0; i < numberOfThreads; i++ {
		go eat(&waitGroup, footMan, forks, i)
	}

	waitGroup.Wait()
}
