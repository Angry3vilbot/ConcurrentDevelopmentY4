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
