package coursera

import "fmt"

func BubbleSort(arr []int) {
	// Create a bubble sort algorithm
	fmt.Println("this is a bubble sort algorithm")
	fmt.Println(arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			Swap(arr, j)
		}
	}
	fmt.Println(arr)
}

func Swap(swapSlice []int, i int) []int {
	if swapSlice[i] > swapSlice[i+1] {
		swapSlice[i], swapSlice[i+1] = swapSlice[i+1], swapSlice[i]
	}
	return swapSlice
}
