package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	var perMap map[string]string
	perMap = make(map[string]string)

	var naming string
	var addressing string

	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter the name: ")
	scanner.Scan()
	naming = scanner.Text()

	fmt.Println("Please enter the address: ")
	scanner.Scan()
	addressing = scanner.Text()

	perMap["name"] = naming

	perMap["address"] = addressing

	person, _ := json.Marshal(perMap)

	fmt.Println("JSON: ", string(person))
}
