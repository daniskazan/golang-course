package main

import (
	"fmt"
	"math"
)

func main() {
	squares := make(chan int, 3)
	cubes := make(chan int, 3)
	for i := 1; i <= 3; i++ {
		squares <- int(math.Pow(float64(i), 2))
		cubes <- int(math.Pow(float64(i), 3))
	}
	close(squares)
	close(cubes)

	for i := 1; i <= 10; i++ {
		select {
		case square := <-squares:
			fmt.Println("squared:", square)
		case cube := <-cubes:
			fmt.Println("cubes:", cube)
		default:
			fmt.Println("NO DATA!")
		}
	}
}
