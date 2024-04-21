// https://github.com/01-edu/public/tree/master/subjects/atoi

package piscine

func Atoi(s string) int {
	sign := 1
	result := 0
	// Check if the string starts with a sign
	if s != "" && (s[0] == '+' || s[0] == '-') {
		if s[0] == '-' {
			sign = -1
		}
		s = s[1:] // Skip the sign character
	}
	// Iterate over each character in the remaining string
	for _, char := range s {
		// Check if the character is not a digit
		if char < '0' || char > '9' {
			return 0 // Non-digit character encountered, return 0
		}
		result = result*10 + int(char-'0')
	}
	return sign * result
}
