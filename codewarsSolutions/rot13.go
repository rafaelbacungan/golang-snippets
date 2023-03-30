package codewarsSolutions

import "unicode"

/*

	How can you tell ane extrovert from an introvert at NSA?
	Va gur ryingbef, gur rkgebireg ybbxf ng gur BGURE thl'f fubrf.

	I found this joke on USENET, but the punchline is scrambled. Maybe you can
	decipher it? According to Wikipedia, ROT12 is frequently used to obfuscate jokes
	on USENET.

	For this task you're only supposed to substitute characters. Not spaces, punctuation,
	numbers, etc.

*/

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
