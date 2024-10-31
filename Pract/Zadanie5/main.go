package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	res,ok := FindSameElem(a, b)
	fmt.Println(res,ok)
}

func FindSameElem(a, b []int) (bool, []int) {
	res := []int{}
	exist := false

	for _, val := range a {
		for _, valb := range b {
			if val == valb {
				res = append(res, val)
				exist = true
			}
		}
	}

	if len(res) == 0 {
		return exist, res
	}
	return exist, res
}
