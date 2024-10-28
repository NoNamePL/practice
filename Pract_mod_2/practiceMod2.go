package main

import (
	"fmt"
	"time"
)

// GO 1.23.2
func main() {

	// Задание 1

	// Пустой слайс и слайс с предопределнными значениями
	var emptySlice []int
	predefinedSliec := []int{1, 2, 3}

	_ = emptySlice

	// Создание слайса через make
	sizedSlice := make([]int, 0, 5)

	_ = sizedSlice

	// Инициализация слайса с литералом
	literalSlice := []int{4, 5, 6, 7, 8}

	_ = literalSlice

	// Задание 2
	fmt.Println("Задание 2\n")
	// Добавление элемента
	nums := append(predefinedSliec, 4, 5, 6)
	fmt.Println("Добавление элемента ", nums)

	// Удаление элемента(сохранение порядка)
	nums = append(nums[:1], nums[2:]...)
	fmt.Println("Удаление элемента(сохранение порядка) ", nums)

	// Удаление элемента(без сохранения порядка)
	nums[2] = nums[len(nums)-1]
	nums = nums[:len(nums)-1]

	fmt.Println("Удаление элемента(без сохранения порядка) ", nums)

	// вставка элемента в середину
	middleIndex := len(nums) / 2
	nums = append(nums[:middleIndex], append([]int{100}, nums[middleIndex:]...)...)
	fmt.Println("вставка элемента в середину ", nums)

	// Задание 3
	fmt.Println("\nЗадание 3\n")
	numbers := make([]int, 0, 1)
	for i := range 100 {
		numbers = append(numbers, i)
		fmt.Println("Длина: ", len(numbers), "Емкость: ", cap(numbers))
	}

	// Проверка на производительность
	n := 10000000
	checkArr := make([]int, 0, n)
	for i := range n {
		checkArr = append(checkArr, i)
	}
	checkTime := time.Now()
	checkArr = append(checkArr, 1)
	duration := time.Since(checkTime)
	fmt.Println("Проверка на производительность n = 10000000", duration)

	// при маленьком объеме слайса
	numbers = make([]int, 0)
	checkTime = time.Now()
	numbers = append(numbers, 1)
	duration = time.Since(checkTime)
	fmt.Println("Проверка на производительность пустой слайс", duration)

	// Задание 4

	fmt.Println("\nЗадание 4\n")
	testSlice := []int{}
	testSlice = modSlice(testSlice, 6)
	fmt.Println("Модифицированный слайс: ", testSlice)

	// Проверка на кратность 2
	testSlice = filterSlice(testSlice, filter)
	fmt.Println(testSlice)

}

func filter(i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}

func modSlice(slice []int, elem int) []int {
	return append(slice, elem)
}

func filterSlice(slice []int, filterFunc func(int) bool) []int {
	var result []int
	for _, v := range slice {
		if filterFunc(v) {
			result = append(result, v)
		}
	}

	return result
}
