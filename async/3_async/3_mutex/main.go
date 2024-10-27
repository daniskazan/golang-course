package main

import (
	"sync"
	"time"
)

type storage struct {
	m  map[int]int
	mu sync.Mutex
}

func (s *storage) write(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[0] = i
}

func (s *storage) check() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.m[0] < 0 {
		// some action
	}
}

func main() {
	s := storage{
		m: make(map[int]int),
	}

	go func() {
		for i := 0; i < 1000000; i++ {
			s.write(i)
		}
	}()

	go func() {
		for i := 0; i < 1000000; i++ {
			s.check()
		}
	}()

	time.Sleep(time.Second / 2)
}
