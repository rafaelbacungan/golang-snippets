package main

import (
	"GolangSnippets/codewarsSolutions"
	"fmt"
)

func main() {
	val := codewarsSolutions.Beeramid(4213414124, 23)
	val2 := codewarsSolutions.BeeramidRefactored(4213414124, 23)
	val3 := codewarsSolutions.BeeramidRefactored(21, 1.5)
	val4 := codewarsSolutions.BeeramidRefactored(-1, 4)
	fmt.Println(val)
	fmt.Println(val2)
	fmt.Println(val3)
	fmt.Println(val4)
}
