package main

import (
	"fmt"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return c
}
func fanIn(inputs ...<-chan string) <-chan string {
	out := make(chan string)
	for _, inp := range inputs {
		go func(input <-chan string) {
			for {
				out <- <-input
			}
		}(inp)
	}
	return out
}

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)

	}
}
