package main

import "fmt"

func main() {
	c := make(chan int) // or make buffered chan with buffer=5
	go generate(c)

	for x := range c {
		fmt.Println(x)
	}
}

func generate(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}
