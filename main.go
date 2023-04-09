package main

import (
	"fmt"
)

func main() {
	val := Chaser(2, 3)
	fmt.Println(val)
}

func Chaser(speed, time int) int {

	var maxDistance int

	if time == 1 {
		return speed * 2
	}

	for i := 0; i < time; i++ {
		if i == time-1 {
			maxDistance += speed * 2
		} else {
			maxDistance += speed
		}
	}

	return maxDistance
}
