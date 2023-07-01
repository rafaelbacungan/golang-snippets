package coursera

import "fmt"

/*
	Test the program by running it twice and test it with two different sets of values
	for acceleration, initial velocity, initial displacement, and time.

	The code must contain a function called GenDisplaceFn() which takes three
	float64 arguments, acceleration a, initial velocity vo, and initial displacement
	so and returns a function, which computes displacement as a function of time,
*/

func ModuleTwo() {
	var a, vo, so, t float64

	fmt.Println("enter acceleration: ")
	fmt.Scan(&a)
	fmt.Println("enter initial velocity: ")
	fmt.Scan(&vo)
	fmt.Println("enter initial displacement: ")
	fmt.Scan(&so)
	fmt.Println("enter time: ")
	fmt.Scan(&t)

	testFn := GenDisplaceFn(a, vo, so)
	fmt.Println("Displacement: ", testFn(t))
}

func GenDisplaceFn(a float64, vo float64, so float64) func(float64) float64 {
	var s float64
	var fn func(float64) float64

	fn = func(t float64) float64 {
		s = ((vo * t) + (0.5 * a * t * t)) + so
		return s
	}

	return fn
}
