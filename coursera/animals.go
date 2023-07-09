package coursera

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name       string
	foodEaten  string
	locomotion string
	sound      string
}

func (c Cow) Eat() {
	fmt.Println(c.foodEaten)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.sound)
}

type Bird struct {
	name       string
	foodEaten  string
	locomotion string
	sound      string
}

func (b Bird) Eat() {
	fmt.Println(b.foodEaten)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.sound)
}

type Snake struct {
	name       string
	foodEaten  string
	locomotion string
	sound      string
}

func (s Snake) Eat() {
	fmt.Println(s.foodEaten)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.sound)
}

func animalCommand(command string, animal Animal) {
	switch command {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Invalid action")
	}
}

func ModuleFour() {
	var command, animalName, animalType string
	var animals []Animal

	fmt.Println("Enter q to quit.")
	for {
		fmt.Print("> ")
		fmt.Scanln(&command, &animalName, &animalType)

		if command == "q" {
			break
		}

		switch command {
		case "newanimal":
			switch animalType {
			case "cow":
				animals = append(animals, Cow{animalName, "grass", "walk", "moo"})
				fmt.Println("Created it!")
			case "bird":
				animals = append(animals, Bird{animalName, "worms", "fly", "peep"})
				fmt.Println("Created it!")
			case "snake":
				animals = append(animals, Snake{animalName, "mice", "slither", "hsss"})
				fmt.Println("Created it!")
			default:
				fmt.Println("Invalid animal type")
			}
		case "query":
			for _, animal := range animals {
				switch animal.(type) {
				case Cow:
					if animal.(Cow).name == animalName {
						animalCommand(animalType, animal.(Cow))
					}
				case Bird:
					if animal.(Bird).name == animalName {
						animalCommand(animalType, animal.(Bird))
					}
				case Snake:
					if animal.(Snake).name == animalName {
						animalCommand(animalType, animal.(Snake))
					}
				default:
					fmt.Println("Invalid animal type")
				}
			}

		default:
			fmt.Println("Invalid command")
		}
	}

}
