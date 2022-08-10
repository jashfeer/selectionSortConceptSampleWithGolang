package main

import "fmt"

func sort(array []int) []int{
	l := len(array)
	for i := 0; i < l; i++ {
		min := i
		for j := i; j < l; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		array[i],array[min]=array[min],array[i]
	}
	return array
}

func main() {
	array:=[]int{2,5,7,3,5,6,8,0,1,2,3,4}
	fmt.Println(sort(array))

}



