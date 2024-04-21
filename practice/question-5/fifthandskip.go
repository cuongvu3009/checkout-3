// https://github.com/01-edu/public/tree/master/subjects/fifthandskip

package piscine

import "strings"

func FifthAndSkip(str string) string {
	// If the string is empty, return a newline
	if len(str) == 0 {
		return "\n"
	}

	// If the string is less than 5 characters, return "Invalid Input" followed by a newline
	if len(str) < 5 {
		return "Invalid Input\n"
	}

	var result strings.Builder

	// Iterate over the string in steps of 5 characters
	for i := 0; i < len(str); i += 5 {
		// Add the current 5 characters to the result
		if i+5 <= len(str) {
			result.WriteString(str[i : i+5])
		} else {
			result.WriteString(str[i:])
		}

		// If there's a character after the fifth character, add a space
		if i+6 < len(str) {
			result.WriteByte(' ')
			// Move the index to the next character after the skipped sixth character
			i++
		}
	}

	return result.String()
}
