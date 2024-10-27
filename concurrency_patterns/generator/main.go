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

func main() {
	joe := boring("Joe!")
	marie := boring("Marie!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say %q\n", <-joe)
		fmt.Printf("You say %q\n", <-marie)
	}
	fmt.Printf("I am leaving.")
}
