package main

import (
	"fmt"
	"sync"
)

type SemaphoreWaitGroup struct {
	count int
	// sem   chan int
	mu    sync.Mutex
}

func (wg *SemaphoreWaitGroup) Add(delta int) {
	wg.mu.Lock()
	wg.count += delta
	wg.mu.Unlock()
}

func (wg *SemaphoreWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *SemaphoreWaitGroup) Wait() {
	for {
		wg.mu.Lock()
		if wg.count <= 0 {
			wg.mu.Unlock()
			break
		}
		wg.mu.Unlock()

	}
}

func NewSemaphoreWaitGroup() *SemaphoreWaitGroup {
	return &SemaphoreWaitGroup{
		count: 0,
		// sem: make(chan int),
		mu: sync.Mutex{},
	}
}

func main() {
	wg := NewSemaphoreWaitGroup()

	wg.Add(5)

	for idx := range 5 {
		go func(idx int) {
			defer wg.Done()
			fmt.Println(idx)

		}(idx)
	}

	wg.Wait()
	fmt.Println("yes")
}
