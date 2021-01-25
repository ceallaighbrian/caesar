// Package caesar provides a caesar cipher for encoding & decoding strings
package caesar

const Shift = 3

// Encodes shifts each alphabetic character in the string 3 places to the right.
// If the character is not a letter it does nothing
func Encode(text string) string {
	letters := []rune(text)

	for i := 0; i < len(letters); i++ {
		letter := letters[i]
		if letter >= 65 && letter <= 90 {
			letters[i] = (letters[i]+Shift-65)%26 + 65
		} else if letter >= 97 && letter <= 122 {
			letters[i] = (letters[i]+Shift-97)%26 + 97
		}
	}
	return string(letters)
}

// Decode shifts each alphabetic character in the string 3 places to the left.
// If the character is not a letter it does nothing
func Decode(text string) string {
	letters := []rune(text)

	for i := 0; i < len(letters); i++ {
		letter := letters[i]
		if letter >= 65 && letter <= 90 {
			letters[i] = (letters[i]-Shift-65)%26 + 65
		} else if letter >= 97 && letter <= 122 {
			letters[i] = (letters[i]-Shift-97)%26 + 97
		}
	}
	return string(letters)
}
