package codewarsSolutions

import "fmt"

func RotatingArray(testArray [][]int) {
	// fmt.Printf("Hello World")
	finalArray := []int{}

	// loop through the array
	// if i == 0 then push all variables to the finalArray
	// if i > 1 but < len(testArray) then push the last variable to the finalArray
	// if i == len(testArray) then push all variables to the finalArray in reverse order
	// if i < len(testArray) but i > 0 then push the first variable to the finalArray in reverse order
	// if i == 1 then push all variables to the finalArray in reverse order except the last variable

	fmt.Println(testArray)
	for i := 0; i < len(testArray); i++ {
		for j := 0; j < len(testArray[i]); j++ {
			if i == 0 {
				finalArray = append(finalArray, testArray[i][j])
			} else if i > 0 && i < len(testArray)-1 {

			} else if i == len(testArray)-1 {

			}

		}
	}
}
