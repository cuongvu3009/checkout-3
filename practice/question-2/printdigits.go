// https://github.com/01-edu/public/tree/master/subjects/printdigits

package piscine

import (
	"github.com/01-edu/z01"
)

func PrintDigits() {
	// Loop through the ASCII values for digits ('0' to '9')
	for c := '0'; c <= '9'; c++ {
		z01.PrintRune(c) // Print each character
	}

	z01.PrintRune('\n') // Print a newline at the end
}
