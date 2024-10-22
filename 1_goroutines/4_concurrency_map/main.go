package main

import (
	"fmt"
	"time"
)

func main() {
	m := make(map[int]int)
	go func() {
		for i := 0; i < 1e6; i++ {
			m[0] = i
		}
	}()
	go func() {
		for i := 0; i < 1e6; i++ {
			if m[0] < 0 {
				fmt.Println("Do smth")
			}
		}
	}()

	time.Sleep(time.Second / 2)
}
