// https://github.com/01-edu/public/tree/master/subjects/saveandmiss

package piscine

func SaveAndMiss(arg string, num int) string {
	// If num is 0 or negative, return the original string
	if num <= 0 {
		return arg
	}

	// Initialize a variable to store the result
	result := ""

	// Iterate over the characters of the string
	for i, char := range arg {
		// Check if the current index modulo num is not equal to 1
		if (i+1)%num != 1 {
			continue // Skip the character if it's not in a "save" position
		}
		result += string(char) // Otherwise, save the character
	}

	return result
}
