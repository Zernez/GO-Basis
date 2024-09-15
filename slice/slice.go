package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var arr [4]int

	slice := arr[0:1]

	for i := 0; i < 4; i++ {

		fmt.Println("Enter a number or \"X\" for exit: ")

		var insval string

		fmt.Scan(&insval)

		if insval == "X" {

			break
		} else {

			arr[i], _ = strconv.Atoi(insval)
			slice = arr[0 : i+1]

			sort.Ints(slice)

			fmt.Println(slice)
		}
	}
}
