package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	//val := Chaser(2, 3)
	//fmt.Println(val)
	// testWebScrape()
	fmt.Println(ModifyString("oll123eH56"))

}

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
