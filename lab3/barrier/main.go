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
// Created on 21/9/2026
// Modified by:
// Issues: N/A
// Help Received: N/A
// Help Provided: Oliwier Jakubiec
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func fib(N int) int {
	if N < 2 {
		return 1
	} else {
		return fib(N-1) + fib(N-2)
	}
}

var cache = make(map[int]int)
var mutex sync.Mutex

func parFib(N int) int {
	var wg sync.WaitGroup
	var A, B int
	wg.Add(2)
	if N < 2 {
		return 1
	}
	mutex.Lock()
	if v, ok := cache[N]; ok {
		mutex.Unlock()
		return v
	}
	mutex.Unlock()

	go func(N int, Ans *int) {
		defer wg.Done()
		*Ans = parFib(N - 1)
	}(N, &A)
	go func(N int, Ans *int) {
		defer wg.Done()
		*Ans = parFib(N - 2)
	}(N, &B)
	wg.Wait()
	mutex.Lock()
	cache[N] = A + B
	mutex.Unlock()
	return A + B
}

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	for i := 0; i < 10; i++ {
		Seq := fib(i * 5)
		par := parFib(i * 5)
		fmt.Println(Seq, "---", par)
	}

}
