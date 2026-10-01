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
// Created on 1/10/2026
// Modified by:
// Issues: N/A
// Help Received: N/A
// Help Provided: N/A
//--------------------------------------------

package main

import (
	"math/rand/v2"
	"sync"
	"time"
)

// Function ran by the producer threads
func produce(index int, buffer chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < index; i++ {
		artifact := rand.IntN(100)
		time.Sleep(time.Duration(artifact) * time.Millisecond)
		println("Producer ", index, " produced an artifact: ", artifact)
		buffer <- artifact
		// Slightly misleading, it will do one of two things:
		//
		// 1. If there are free consumers waiting on an artifact, Go will directly pass the artifact to one,
		// bypassing the channel entirely.
		// 2. If all consumers are busy, it will try storing it in the buffer (does so if there is space,
		// locks if there isn't).
		//
		// For example, say you have 5 consumers and 5 slots in the buffer.
		// You can send off 10 artifacts before the producers start getting locked
		// if all consumers stay busy.
		println("Producer ", index, " stored ", artifact, " inside the buffer")
	}
}

// Function ran by the consumer threads
func consume(index int, buffer chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Retrieves data from the channel until it is closed
	for artifact := range buffer {
		time.Sleep(time.Duration(artifact) * time.Millisecond)
		println("Consumer ", index, " consumed an artifact: ", artifact)
	}
}

func main() {
	producerWG := sync.WaitGroup{}
	consumerWG := sync.WaitGroup{}
	buffer := make(chan int, 5)

	for i := 0; i < 5; i++ {
		producerWG.Add(2)
		consumerWG.Add(1)
		go produce(i, buffer, &producerWG)
		go produce(i, buffer, &producerWG)
		go consume(i, buffer, &consumerWG)
	}
	producerWG.Wait()
	// Close the channel once the producers are done
	// Running the consumers permanently with an infinite for loop works
	// But if the app continued executing after this producer-consumer group was done
	// the consumers would stay there permanently, wasting resources
	close(buffer)
	consumerWG.Wait()
	println("The buffer has: ", len(buffer), " elements")
}
