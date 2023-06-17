package coursera

import "fmt"

func ModuleOne() {
	var input int
	var numbers []int

	fmt.Println("Enter integers: ")
	for i := 0; i < 10; i++ {
		fmt.Scan(&input)
		numbers = append(numbers, input)
	}

	BubbleSort(numbers)
}

func BubbleSort(slice []int) {
	fmt.Println("presorted slice: ", slice)
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			Swap(slice, j)
		}
	}
	fmt.Println("sorted slice", slice)
}

func Swap(swapSlice []int, i int) []int {
	if swapSlice[i] > swapSlice[i+1] {
		swapSlice[i], swapSlice[i+1] = swapSlice[i+1], swapSlice[i]
	}
	return swapSlice
}
