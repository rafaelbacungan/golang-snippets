package codewarsSolutions

import (
	"fmt"
	"strings"
)

func CountingDuplicates(s1 string) {
	charArray := []string{}

	//for _, character := range s1 {
	//	if character >= 'a' && character <= 'z' || character >= 'A' && character <= 'Z' {
	//		// check if the character is in the array
	//		// if it is not in the array, add it to the array
	//		loweredChar := strings.ToLower(string(character))
	//		if charArray.len === 0 {
	//			continue
	//		} else {
	//			append(charArray, loweredChar)
	//		}
	//	} else {
	//		if character
	//	}
	//}

	fmt.Println(s1)
	dupliCount := 0

	for _, character := range s1 {
		currentCount := 0
		if character >= 'a' && character <= 'z' || character >= 'A' && character <= 'Z' {
			loweredChar := strings.ToLower(strings.ToLower(string(character)))
			for _, characterLetter := range s1 {
				if strings.ToLower(string(characterLetter)) == loweredChar {
					currentCount++
				}
			}
			charArray = append(charArray, loweredChar)
			if currentCount > 1 {
				fmt.Println("duplicount letter")
				dupliCount++
			}
		} else {
			for _, characterOther := range s1 {
				if string(characterOther) == string(character) {
					currentCount++
				}
			}
			if currentCount > 1 {
				fmt.Println("duplicount other")
				dupliCount++
			}
		}
	}

	fmt.Println(dupliCount)
}
