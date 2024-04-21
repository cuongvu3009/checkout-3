// https://github.com/01-edu/public/tree/master/subjects/strlen

package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func StrLen(s string) int {
	// Initialize a variable to hold the length
	length := 0

	// Iterate over each rune in the string
	for range s {
		// Increment the length for each rune
		length++
	}

	// Print the length using PrintRune
	PrintStrLen(length)

	// Return the total length
	return length
}

// PrintStrLen prints the length using PrintRune
func PrintStrLen(length int) {
	for _, r := range fmt.Sprintf("%d", length) {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
