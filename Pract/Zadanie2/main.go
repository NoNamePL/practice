package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 1
	fmt.Println("1")
	originalSlice := make([]int, 10)

	rand.NewSource(time.Now().Unix())

	for idx := range 10 {
		originalSlice[idx] = rand.Int()
	}

	fmt.Println(originalSlice)

	// 2
	fmt.Println(2)

	res := sliceExample(originalSlice)

	fmt.Println(res)

	// 3
	fmt.Println(3)

	res = addElements(originalSlice, 12)
	fmt.Println(res)

	// 4
	fmt.Println(4)

	res = copySlice(originalSlice)
	fmt.Println(res)
	fmt.Println(originalSlice)

	// 5
	fmt.Println(5)

	res = removeElement(originalSlice, 15)
	fmt.Println(res)

}

func sliceExample(slice []int) []int {

	res := make([]int, 0, len(slice))

	for _, val := range slice {
		if val%2 == 0 {
			res = append(res, val)
		}
	}

	return res

}

func addElements(slice []int, element int) []int {
	res := make([]int, len(slice), len(slice)+1)

	copy(res, slice)
	res = append(res, element)

	return res
}

func copySlice(slice []int) []int {
	res := make([]int, len(slice))
	copy(res, slice)
	return res
}

func removeElement(slice []int, index int) []int {
	if index > (len(slice)-1) || index < 0 {
		return nil
	}
	res := make([]int, len(slice))

	res = append(slice[:index], slice[index+1:]...)
	return res
}
