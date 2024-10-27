package main

import (
	"fmt"
	"runtime"
	"time"
)

func demo1() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 3; i++ {
		i := i
		go func() {
			for j := 0; j < 10; j++ {
				fmt.Printf("i: %d, j: %d\n", i, j)
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)
}

func main() {
	demo1()
}
