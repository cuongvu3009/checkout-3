// https://github.com/01-edu/public/tree/master/subjects/iscapitalized

package piscine

func IsCapitalized(s string) bool {
	// Check if the string is empty
	if s == "" {
		return false
	}

	// Initialize a variable to track if each word is capitalized
	isCapitalized := true

	// Iterate over each word in the string
	for i := 0; i < len(s); i++ {
		// Skip leading whitespace
		for i < len(s) && s[i] == ' ' {
			i++
		}

		// Check if we reached the end of the string
		if i == len(s) {
			break
		}

		// Check if the first character of the word is lowercase
		if s[i] >= 'a' && s[i] <= 'z' {
			isCapitalized = false
			break
		}

		// Move to the end of the current word
		for i < len(s) && s[i] != ' ' {
			i++
		}
	}

	return isCapitalized
}
