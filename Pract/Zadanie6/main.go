package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ch := make(chan int)

	go Generator(ch)

	for {
		fmt.Println(<-ch)
		time.Sleep(time.Second * 1)
	}
}

func Generator(ch chan<- int) {
	rand.NewSource(time.Now().Unix())
	for {
		ch <- rand.Int()
	}
}
