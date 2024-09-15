package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	var users []name
	var filename string
	fmt.Println("Enter Filename: ")
	fmt.Scan(&filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() //close the file just before exiting

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { //scanning line by line
		username := scanner.Text()
		nameslice := strings.Split(username, " ")
		user := name{
			fname: nameslice[0],
			lname: nameslice[1],
		}
		users = append(users, user)
	}

	for _, u := range users {
		fmt.Println("First name: ", u.fname, "|", "Last name: ", u.lname)
	}
}
