// Package caesar provides a caesar cipher for encoding & decoding strings
package caesar

// Encodes shifts each alphabetic character in the string n places to the right.
// If the character is not a letter it does nothing
func Encode(text string, shift int32) string {
	letters := []rune(text)

	for i := 0; i < len(letters); i++ {
		letter := letters[i]
		if letter >= 65 && letter <= 90 {
			letters[i] = (letters[i]+shift-65)%26 + 65
		} else if letter >= 97 && letter <= 122 {
			letters[i] = (letters[i]+shift-97)%26 + 97
		}
	}
	return string(letters)
}

// Decode shifts each alphabetic character in the string n places to the left.
// If the character is not a letter it does nothing
func Decode(text string, shift int32) string {
	letters := []rune(text)

	for i := 0; i < len(letters); i++ {
		letter := letters[i]
		if letter >= 65 && letter <= 90 {
			letters[i] = (letters[i]-shift-65)%26 + 65
		} else if letter >= 97 && letter <= 122 {
			letters[i] = (letters[i]-shift-97)%26 + 97
		}
	}
	return string(letters)
}
