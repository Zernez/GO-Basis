package main

import (
	"fmt"
)

type Animal struct {
	Food       string
	Locomotion string
	Noise      string
}

func (data Animal) Eat() {

	fmt.Println("This animal eat", data.Food)
}

func (data Animal) Move() {

	fmt.Println("This animal move", data.Locomotion)
}

func (data Animal) Speak() {

	fmt.Println("This animal speak with a", data.Noise)
}

func main() {

	var requestA string
	var requestQ string
	var reqAnimal Animal

	cow := Animal{
		Food:       "grass",
		Locomotion: "walk",
		Noise:      "moo",
	}

	bird := Animal{
		Food:       "worms",
		Locomotion: "fly",
		Noise:      "beep",
	}

	snake := Animal{
		Food:       "mice",
		Locomotion: "slither",
		Noise:      "hsss",
	}

	fmt.Println("Type a animal from \"cow, bird, snake\" and a request from \"eat, move, speak\" separated by a space")
	fmt.Scan(&requestA, &requestQ)

	switch requestA {

	case "cow":
		reqAnimal = cow

	case "bird":
		reqAnimal = bird

	case "snake":
		reqAnimal = snake
	}

	switch requestQ {

	case "eat":
		reqAnimal.Eat()

	case "move":
		reqAnimal.Move()

	case "speak":
		reqAnimal.Speak()
	}
}
