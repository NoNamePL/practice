package main

import (
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

}