package main

import (
	"fmt"
	"strings"
)

func main() {

	var searchstr string

	fmt.Println("Enter a string: ")

	fmt.Scan(&searchstr)

	searchstr = strings.ToLower(searchstr)

	fmt.Println(searchstr)

	if strings.Contains(searchstr, "a") && strings.HasPrefix(searchstr, "i") && strings.HasSuffix(searchstr, "n") {

		fmt.Println("Found!")
	} else {

		fmt.Println("Not Found!")
	}
}
