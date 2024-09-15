package main

import (
	"fmt"
)

func bubbleSort(array []int, size int) {

	for i := 0; i < size; i++ {

		for j := i; j < size; j++ {

			if array[i] > array[j] {

				swap(array, i, j)

			}
		}
	}
}

func swap(array []int, major int, minor int) {

	tmp := array[major]
	array[major] = array[minor]
	array[minor] = tmp

}

func main() {

	array := make([]int, 10)

	for i := 0; i < 10; i++ {
		fmt.Println("Enter integer element: ")
		fmt.Scan(&array[i])
	}

	bubbleSort(array, 10)

	fmt.Println("Sorted Array: ")

	for i := 0; i < 10; i++ {
		fmt.Println(array[i])
	}

}
