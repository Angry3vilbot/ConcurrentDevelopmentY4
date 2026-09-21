//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Mykhailo Balaker
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
