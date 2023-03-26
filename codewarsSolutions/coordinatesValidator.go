package codewarsSolutions

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

func IsValidCoordinates(coordinates string) bool {
	regexNumber := regexp.MustCompile("[0-9]+[.]?[0-9]*[[:alnum:]]*")
	regexLetter := regexp.MustCompile("[a-zA-Z]+")

	coordinatesArr := regexNumber.FindAllString(coordinates, -1)
	letters := regexLetter.FindAllString(coordinates, -1)

	commaCount := 0

	if len(coordinatesArr) != 2 || len(letters) > 0 {
		return false
	}

	for i := 0; i < len(coordinatesArr); i++ {
		n, err := strconv.ParseFloat(coordinatesArr[i], 64)
		if err != nil {
		}

		if i == 0 {
			if n > 90 || n < -90 {
				return false
			}
		}

		if i == 1 {
			if n > 180 || n < -180 {
				fmt.Println(n)
				return false
			}
		}
	}

	for j := 0; j < len(coordinates); j++ {
		if coordinates[j] == '-' {
			nextVal := rune(coordinates[j+1])
			if !unicode.IsNumber(nextVal) {
				return false
			}
		}
		if coordinates[j] == ',' {
			commaCount++
			if commaCount > 1 {
				return false
			}
		}
	}

	return true
}
