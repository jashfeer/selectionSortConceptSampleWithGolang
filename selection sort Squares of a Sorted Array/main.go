package main

import "fmt"

func selectionSort(array []int) []int {
	l := len(array)
	for i := 0; i < l; i++ {
		array[i] = array[i] * array[i]
	}
	for i := 0; i < l; i++ {
		min := i
		for j := i; j < l; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]
	}

	return array
}

func main() {
	nums := []int{-2, 4, 5, 2, -6, 8}
	fmt.Println(selectionSort(nums))
}
