package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Generator(ch chan uint8) {
	rand.NewSource(time.Now().Unix())
	for {
		ch <- uint8(rand.Uint32())
	}
}

func Read(chU chan uint8, chF chan float64) {
	for {
		temp := float64(<-chU)
		chF <- temp * temp * temp
	}
}

func main() {
	chU := make(chan uint8)
	chF := make(chan float64)
	go Generator(chU)
	go Read(chU, chF)
	for {
		time.Sleep(time.Second * 1)
		fmt.Println(<-chF)
	}

}
