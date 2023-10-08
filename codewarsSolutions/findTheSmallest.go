package codewarsSolutions

import "fmt"

func Smallest(n int64) []int64 {

	// create the smallest possible number from an integer by moving only one number
	// to a different position
	// return an array of 3 items, consisting of the smallest number, the number to remove,
	// and the index from where its from

	digits := n
	var smallestDigit int64
	digitsSlice := []int64{}

	for digits > 0 {
		digit := digits % 10

		if digits == n {
			smallestDigit = digit
		}

		if digit < smallestDigit {
			smallestDigit = digit
		}

		fmt.Println("digit: ", digit)

		digitsSlice = append(digitsSlice, digit)
		digits = digits / 10
	}

	fmt.Println("digitsSlice: ", digitsSlice)
	fmt.Println("")

	return []int64{1}
}
