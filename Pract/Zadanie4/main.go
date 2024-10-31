package main

import (
	"fmt"
	"slices"
)

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	res := misMatch(slice1, slice2)

	fmt.Println(res)
}

func misMatch(slice1, slice2 []string) []string {

	res := make([]string, 0, len(slice1))

	for _, val := range slice1 {
		if !(slices.Contains(slice2, val)) {
			res = append(res, val)
		}
	}

	return res
}
