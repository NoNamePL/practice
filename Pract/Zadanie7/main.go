package main

import (
	"fmt"
	"sync"
)

func main() {

	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	aArr := []int{1, 2, 3}
	bArr := []int{10, 20, 30}
	cArr := []int{100, 200, 300}

	go func() {
		for _, val := range aArr {
			a <- val
		}
		close(a)
	}()

	go func() {
		for _, val := range bArr {
			b <- val
		}
		close(b)
	}()
	go func() {
		for _, val := range cArr {
			c <- val
		}
		close(c)
	}()

	res := joinChannels(a, b, c)

	for val := range res {
		fmt.Println(val)
	}

}

func joinChannels(chs ...chan int) chan int {
	res := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}
		mg := &sync.Mutex{}
		wg.Add(len(chs))
		for _, c := range chs {
			c := c
			go func() {
				defer wg.Done()
				for val := range c {
					mg.Lock()
					res <- val
					mg.Unlock()
				}
			}()
		}

		wg.Wait()

		close(res)
	}()

	return res
}
