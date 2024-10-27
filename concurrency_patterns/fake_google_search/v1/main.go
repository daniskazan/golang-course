package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type SearchFunc func(query string) Result

func fakeSearch(kind string) SearchFunc {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%v result for %v", kind, query))
	}
}

func Google(query string) []Result {
	Image := fakeSearch("Images")
	Video := fakeSearch("Videos")
	Web := fakeSearch("Web")
	results := make([]Result, 3)
	backends := []SearchFunc{Image, Video, Web}
	for _, backend := range backends {
		results = append(results, backend(query))
	}
	return results
}

func main() {
	start := time.Now()
	results := Google("golang")
	fmt.Println(results)
	fmt.Println(time.Since(start))
}
