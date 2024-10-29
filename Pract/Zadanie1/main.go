package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	//1

	fmt.Println(1)

	var numDecimal int = 42 // Десятичная система
	fmt.Println(numDecimal, "имеет тип:", reflect.TypeOf(numDecimal))
	var numOctal int = 052 // Восьмеричная система
	fmt.Println(numOctal, "имеет тип:", reflect.TypeOf(numOctal))
	var numHexadecimal int = 0x2A // Шестнадцатиричная система
	fmt.Println(numHexadecimal, "имеет тип:", reflect.TypeOf(numHexadecimal))
	var pi float64 = 3.14
	fmt.Println(pi, "имеет тип:", reflect.TypeOf(pi))
	var name string = "Golang"
	fmt.Println(name, "имеет тип:", reflect.TypeOf(name))
	var isActive bool = true
	fmt.Println(isActive, "имеет тип:", reflect.TypeOf(isActive))
	var complexNum complex64 = 1 + 2i
	fmt.Println(complexNum, "имеет тип:", reflect.TypeOf(complexNum))

	// 2
	fmt.Println(2)
	str := ""

	str += strconv.Itoa(numDecimal) + strconv.Itoa(numOctal) +
		strconv.Itoa(numHexadecimal) + name + strconv.FormatBool(isActive) +
		strconv.FormatComplex(complex128(complexNum), 'f', -1, 64)
	fmt.Println(str)

	// 3
	fmt.Println(3)
	arrRun := []rune(str)
	fmt.Println(arrRun)

	// 4
	fmt.Println(4)
	arByte := []byte(string(arrRun))
	che := []byte("g0-2024")

	fmt.Println(arByte)
	for idx,jdx := 8, 0; jdx < len(che); idx,jdx = idx+1,jdx+1 {
		arByte[idx] = che[idx-8]
	}


	h := sha256.Sum256(arByte)

	fmt.Printf("%x", h)

}
