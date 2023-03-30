package main

import (
	"fmt"
	"unicode"
)

func main() {
	val := Rot13("NSA rknzcyr.")
	fmt.Println(val)
}

func Rot13(msg string) string {
	var encode string
	for _, character := range msg {
		var currentChar rune
		if unicode.IsLetter(character) {
			if unicode.IsUpper(character) {
				loweredChar := unicode.ToLower(character)
				currentChar = unicode.ToUpper(rune(int(loweredChar-'a'+13)%26 + 'a'))
			} else {
				currentChar = rune(int(character-'a'+13)%26 + 'a')
			}
			encode += string(currentChar)
		} else {
			encode += string(character)
		}

	}
	return encode
}
