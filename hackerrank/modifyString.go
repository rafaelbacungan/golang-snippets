package hackerrank

import (
	"regexp"
	"strings"
)

func ModifyString(str string) string {
	regexLetter := regexp.MustCompile("[a-zA-Z]+")
	clearedNumbers := regexLetter.FindAllString(str, -1)
	joinedString := strings.Join(clearedNumbers, "")
	var newString string

	for i := len(joinedString); i > 0; i-- {
		newString += string(joinedString[i-1])
	}

	return newString
}
