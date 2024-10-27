package main

import (
	"fmt"
	"time"
)

type printer struct {
}

func (p *printer) printHello() {
	fmt.Println("Hi from struct method")
}

func printHello() {
	fmt.Println("Hi from named func")
}

func main() {
	go func() {
		fmt.Println("Hi from anonymous func")
		time.Sleep(time.Second * 2)
		fmt.Println("Hi2 from anonymous func")
	}()
	go printHello()

	var p printer
	p.printHello()

	time.Sleep(time.Second)

}
