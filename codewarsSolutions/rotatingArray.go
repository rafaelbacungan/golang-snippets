package codewarsSolutions

import "fmt"

// TODO: Find a way to refactor this
func RotatingArray(testArray [][]int) {
	up := 0
	down := len(testArray) - 1
	left := 0
	right := len(testArray[down]) - 1
	dir := 0
	finalArr := []int{}

	fmt.Println(down, right)

	for up <= down && left <= right {
		if dir == 0 {
			for i := left; i <= right; i++ {
				finalArr = append(finalArr, testArray[up][i])
			}
			up++
		} else if dir == 1 {
			for i := up; i <= down; i++ {
				finalArr = append(finalArr, testArray[i][right])
			}
			right--
		} else if dir == 2 {
			for i := right; i >= left; i-- {
				finalArr = append(finalArr, testArray[down][i])
			}
			down--
		} else if dir == 3 {
			for i := down; i >= up; i-- {
				finalArr = append(finalArr, testArray[i][left])
			}
			left++
		}
		dir = (dir + 1) % 4
	}

	fmt.Println(finalArr)

}
