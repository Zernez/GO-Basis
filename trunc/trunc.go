package main

import "fmt"

func main() {

	var number float32

	fmt.Println("Enter a float value : ")

	fmt.Scan(&number)

	fmt.Println("Truncated number", int(number))
}
