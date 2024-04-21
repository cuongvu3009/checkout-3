// https://github.com/01-edu/public/tree/master/subjects/itoa

package piscine

func Itoa(n int) string {
	// Handle the special case when n is 0
	if n == 0 {
		return "0"
	}

	// Initialize an empty string to store the result
	var result string

	// Handle negative numbers
	if n < 0 {
		// Add the negative sign to the result
		result += "-"

		// Change the sign of n to positive
		n = -n
	}

	// Convert each digit of n to its corresponding character and prepend to the result
	for n > 0 {
		digit := n % 10
		// Convert the digit to its character representation and prepend it to the result
		result = string(rune('0'+digit)) + result
		// Move to the next digit
		n /= 10
	}

	return result
}
