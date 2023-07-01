package coursera

import (
	"fmt"
)

type Animal struct {
	foodEaten  string
	locomotion string
	sound      string
}

func (a Animal) Eat() string {
	return a.foodEaten
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.sound
}

func ModuleThree() {
	var animal, action string
	animals := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	fmt.Print(">")
	fmt.Scanln(&animal, &action)

	switch action {
	case "eat":
		fmt.Println(animals[animal].Eat())
	case "move":
		fmt.Println(animals[animal].Move())
	case "speak":
		fmt.Println(animals[animal].Speak())
	default:
		fmt.Println("Invalid input")
	}
}
