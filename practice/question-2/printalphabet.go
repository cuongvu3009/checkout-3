// https://github.com/01-edu/public/tree/master/subjects/printalphabet

package piscine

import (
	"github.com/01-edu/z01"
)

func PrintAlphabet() {
	// Loop through the ASCII values for lowercase letters ('a' to 'z')
	for c := 'a'; c <= 'z'; c++ {
		z01.PrintRune(c) // Print each character
	}

	z01.PrintRune('\n') // Print a newline at the end
}
