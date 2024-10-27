package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var mux sync.Mutex
	var wg sync.WaitGroup // Create a WaitGroup

	for i := 0; i < 1e5; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go func() {
			defer wg.Done() // Decrement the counter when the goroutine completes
			mux.Lock()
			count++
			mux.Unlock()
		}()
	}

	wg.Wait()                          // Wait for all goroutines to finish
	fmt.Println("Final count:", count) // Print the final count
}
