package main

import (
	"fmt"
	"sync"
)

func bubbleSort(array []int, chans chan int, wg *sync.WaitGroup) {

	size := len(array)

	for i := 0; i < size; i++ {

		for j := i; j < size; j++ {

			if array[i] > array[j] {

				swap(array, i, j)

			}
		}

		chans <- array[i]
	}

	fmt.Println("Sorted sliced array: ", array)
	wg.Done()
}

func swap(array []int, major int, minor int) {

	tmp := array[major]
	array[major] = array[minor]
	array[minor] = tmp

}

func main() {

	var quantity int

	fmt.Println("Enter how many integer element do you want to insert: ")
	fmt.Scan(&quantity)

	array := make([]int, quantity)

	for i := 0; i < quantity; i++ {
		fmt.Println("Enter the #", i, " integer element: ")
		fmt.Scan(&array[i])
	}

	tempQuant := quantity / 4

	slice1 := array[0:tempQuant]

	slice2 := array[tempQuant : tempQuant*2]

	slice3 := array[tempQuant*2 : tempQuant*3]

	slice4 := array[tempQuant*3 : quantity]

	tempsize := len(slice4)

	var chans = []chan int{make(chan int, tempQuant), make(chan int, tempQuant), make(chan int, tempQuant), make(chan int, tempsize)}

	var wg sync.WaitGroup

	wg.Add(4)

	go bubbleSort(slice1, chans[0], &wg)

	for i := len(slice1) - 1; i == 0; i-- {

		a := <-chans[0]

		slice1[i] = a
	}

	go bubbleSort(slice2, chans[1], &wg)

	for i := len(slice2) - 1; i == 0; i-- {

		b := <-chans[1]

		slice2[i] = b
	}

	go bubbleSort(slice3, chans[2], &wg)

	for i := len(slice3) - 1; i == 0; i-- {

		c := <-chans[2]

		slice3[i] = c
	}

	go bubbleSort(slice4, chans[3], &wg)

	for i := len(slice4) - 1; i == 0; i-- {

		d := <-chans[3]

		slice4[i] = d
	}

	wg.Wait()

	arrayTot := slice1
	arrayTot = append(arrayTot, slice2...)
	arrayTot = append(arrayTot, slice3...)
	arrayTot = append(arrayTot, slice4...)

	fmt.Printf("Sorted complete array: %v", arrayTot)
}
