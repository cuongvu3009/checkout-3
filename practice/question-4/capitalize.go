// https://github.com/01-edu/public/tree/master/subjects/capitalize

package piscine

func Capitalize(s string) string {
	// Convert the string to a rune slice
	r := []rune(s)

	// Check if the first character is a lowercase letter
	if r[0] >= 'a' && r[0] <= 'z' {
		// Convert the first character to uppercase
		r[0] -= 32
	}

	// Iterate over the remaining characters
	for i := 1; i < len(r); i++ {
		// Check if the previous character is not a digit and not a lowercase letter
		if !(r[i-1] >= '0' && r[i-1] <= '9') && !(r[i-1] >= 'a' && r[i-1] <= 'z') {
			// Check if the current character is a lowercase letter
			if r[i] >= 'a' && r[i] <= 'z' {
				// Convert the current character to uppercase
				r[i] -= 32
			}
		}
	}

	// Convert the rune slice back to a string and return it
	return string(r)
}
