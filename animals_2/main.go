package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"mylearning.com.br/data"
)

func main() {

	animals := make(map[string]data.Animal)

	fmt.Println("You can create a new animal using unce of the types below or request information about the animals created.\n * cow\n * bird\n * snake")
	orientation()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n> ")
		scanner.Scan()
		command := scanner.Text()

		pieces := strings.Split(command, " ")

		switch {
		case len(pieces) == 3:
			takeAction(animals, pieces[0], pieces[1], pieces[2])

		case strings.EqualFold(command, "list"):
			printAnimalsList(animals)

		case strings.EqualFold(command, "X"):
			fmt.Println("Bye!")
			return

		default:
			orientation()
		}
	}
}

func takeAction(animals map[string]data.Animal, param1 string, param2 string, param3 string) {
	switch {
	case strings.EqualFold(param1, "newanimal"):
		create(animals, param2, param3)
	case strings.EqualFold(param1, "query"):
		request(animals, param2, param3)
	}
}

func create(animals map[string]data.Animal, name string, animalType string) {
	var animal data.Animal

	switch animalType {
	case "cow":
		animal = data.Cow{}
	case "bird":
		animal = data.Bird{}
	case "snake":
		animal = data.Snake{}
	default:
		fmt.Println("Type not found!")
		return
	}

	animals[name] = animal
	fmt.Println("Created it!")
}

func request(animals map[string]data.Animal, name string, animalInfo string) {
	if animal, ok := animals[name]; ok {
		switch animalInfo {
		case "eat":
			fmt.Println(animal.Eat())
		case "move":
			fmt.Println(animal.Move())
		case "speak":
			fmt.Println(animal.Speak())
		default:
			fmt.Println("Info not found!")
		}
	} else {
		fmt.Println("Animal not found!")
	}
}

func printAnimalsList(animals map[string]data.Animal) {
	for k := range animals {
		fmt.Println(" * ", k)
	}
}

func orientation() {
	fmt.Println("Use one of the commands below.")
	fmt.Println("> newanimal\n  * Create new animal\n  * Eg.: newanimal billy cow")
	fmt.Println("> query\n  * Request info about animal created\n  * Eg.: query billy eat")
	fmt.Println("> list\n  * List animal created\n  * Eg. list")
	fmt.Println("To exit, just type X.")
}
