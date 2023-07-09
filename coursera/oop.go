package coursera

import (
	"fmt"
)

type Animals struct {
	foodEaten  string
	locomotion string
	sound      string
}

func (a Animals) Eat() string {
	return a.foodEaten
}

func (a Animals) Move() string {
	return a.locomotion
}

func (a Animals) Speak() string {
	return a.sound
}

func ModuleThree() {
	var animal, action string
	animals := map[string]Animals{
		"cow":   Animals{"grass", "walk", "moo"},
		"bird":  Animals{"worms", "fly", "peep"},
		"snake": Animals{"mice", "slither", "hsss"},
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
