package codewarsSolutions

import (
	"unicode"
)

func ToNato(words string) string {
	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu"}
	var code string

	for i := 0; i < len(words); i++ {
		if unicode.IsLetter(rune(words[i])) {
			for _, codeName := range nato {
				if string(unicode.ToLower(rune(words[i]))) == string(unicode.ToLower(rune(codeName[0]))) {
					code += codeName
					if i != len(words)-1 {
						code += " "
					}
				}
			}

		} else {
			if string(words[i]) == " " {
				continue
			} else {
				code = code + string(words[i])
				if i != len(words)-1 {
					code += " "
				}
			}
		}
	}

	return code
}
