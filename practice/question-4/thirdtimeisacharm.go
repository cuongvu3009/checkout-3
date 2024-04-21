// https://github.com/01-edu/public/tree/master/subjects/thirdtimeisacharm

package piscine

func ThirdTimeIsACharm(str string) string {
	// Check if the string is empty
	if len(str) == 0 {
		return "\n"
	}

	// Initialize a variable to store the result
	result := ""

	// Iterate over the characters of the string
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	// If there are no third characters, return a newline
	if result == "" {
		return "\n"
	}

	// Return the result followed by a newline
	return result + "\n"
}
