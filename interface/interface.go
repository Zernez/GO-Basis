package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type animalattr struct {
	Food       string
	Locomotion string
	Noise      string
}

type cow struct {
	Name      string
	Attribute animalattr
}

func (data cow) Eat() {

	fmt.Println("This cow eat", data.Attribute.Food)
}
func (data cow) Move() {

	fmt.Println("This cow move by", data.Attribute.Locomotion)
}
func (data cow) Speak() {

	fmt.Println("This cow speak with a", data.Attribute.Noise)
}

type bird struct {
	Name      string
	Attribute animalattr
}

func (data bird) Eat() {

	fmt.Println("This bird eat", data.Attribute.Food)
}
func (data bird) Move() {

	fmt.Println("This bird move by", data.Attribute.Locomotion)
}
func (data bird) Speak() {

	fmt.Println("This bird speak with a", data.Attribute.Noise)
}

type snake struct {
	Name      string
	Attribute animalattr
}

func (data snake) Eat() {

	fmt.Println("This snake eat", data.Attribute.Food)
}
func (data snake) Move() {

	fmt.Println("This snake move by", data.Attribute.Locomotion)
}
func (data snake) Speak() {

	fmt.Println("This snake speak with a", data.Attribute.Noise)
}

func main() {

	var cowlist = make(map[string]cow)
	var birdlist = make(map[string]bird)
	var snakelist = make(map[string]snake)
	var requestA string
	var requestB string
	var requestC string

	cowattr := animalattr{
		Food:       "grass",
		Locomotion: "walk",
		Noise:      "moo",
	}

	birdattr := animalattr{
		Food:       "worms",
		Locomotion: "fly",
		Noise:      "beep",
	}

	snakeattr := animalattr{
		Food:       "mice",
		Locomotion: "slither",
		Noise:      "hsss",
	}

	for {

		fmt.Println("Type \"newanimal\" followed by a name and a type from \"cow\" \"bird\" and \"snake\" distantiated by a space each other, for creating a new animal. Or you can type \"query\" followed by name and a information from \"eat\" \"move\" and \"speak\" distantiated by a space each other for query the list:")
		fmt.Scan(&requestA, &requestB, &requestC)

		switch requestA {

		case "newanimal":

			switch requestC {

			case "cow":

				cowlist[requestB] = cow{Name: requestB, Attribute: cowattr}
				fmt.Println("Cow created")

			case "bird":

				birdlist[requestB] = bird{Name: requestB, Attribute: birdattr}
				fmt.Println("Bird created")

			case "snake":

				snakelist[requestB] = snake{Name: requestB, Attribute: snakeattr}
				fmt.Println("Snake created")
			}

		case "query":

			valc, okc := cowlist[requestB]
			valb, okb := birdlist[requestB]
			vals, oks := snakelist[requestB]

			var animalreq Animal

			if okc {

				animalreq = valc

				switch requestC {

				case "eat":
					animalreq.Eat()
				case "move":
					animalreq.Move()
				case "speak":
					animalreq.Speak()

				}
			}

			if okb {

				animalreq = valb

				switch requestC {

				case "eat":
					animalreq.Eat()
				case "move":
					animalreq.Move()
				case "speak":
					animalreq.Speak()

				}
			}

			if oks {

				animalreq = vals

				switch requestC {

				case "eat":
					animalreq.Eat()
				case "move":
					animalreq.Move()
				case "speak":
					animalreq.Speak()

				}
			}
		}

	}
}
