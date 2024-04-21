// https://github.com/01-edu/public/tree/master/subjects/digitlen

package piscine

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func DigitLen(n, base int) int {
	// Check if the base is within the valid range (2 to 36)
	if base < 2 || base > 36 {
		return -1
	}

	// Convert n to its absolute value
	n = abs(n)

	// Initialize a counter to count the divisions
	count := 0

	// Divide n by base until it becomes zero
	for n > 0 {
		n /= base
		count++
	}

	return count
}
