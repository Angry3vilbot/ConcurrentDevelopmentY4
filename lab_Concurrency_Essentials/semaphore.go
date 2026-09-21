package main

import (
	"fmt"
	"sync"
	"time"
)

// make struct containing channel
// add init, acquire and release
type Semaphore struct {
	theCounter chan struct{}
}

func Init(numWorkers int) *Semaphore {
	if numWorkers < 0 {
		panic("negative weighted semaphore size")
	}
	return &Semaphore{make(chan struct{}, numWorkers)}
}

func (s *Semaphore) Acquire() {
	s.theCounter <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.theCounter
}

func main() {
	maxGoroutines := 5
	var semaphore = Init(maxGoroutines)
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			semaphore.Acquire()
			defer func() { semaphore.Release() }()

			// Simulate a task
			fmt.Printf("Running task %d\n", i)
			time.Sleep(2 * time.Second)
		}(i)
	}
	wg.Wait()
}
