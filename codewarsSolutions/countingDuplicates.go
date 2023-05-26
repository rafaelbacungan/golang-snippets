package codewarsSolutions

import (
	"fmt"
	"unicode"
)

// TODO: Read more about maps and refactor this code
func CountingDuplicates(s1 string) {
	//charArray := []string{}
	fmt.Println(s1)
	dupliCount := make(map[rune]int)
	finalCount := 0

	for _, character := range s1 {
		if !unicode.IsLetter(character) && !unicode.IsNumber(character) {
			continue
		}

		if unicode.IsLetter(character) {
			character = unicode.ToLower(character)
		}

		dupliCount[character]++

	}

	for _, char := range dupliCount {
		if char > 1 {
			finalCount++
		}
	}

	fmt.Println(len(dupliCount))
}
