// https://github.com/01-edu/public/tree/master/subjects/union

package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func Union() {
	// Check if the number of arguments is not equal to 3 (including the program name)
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}

	// Create a map to store encountered characters
	encountered := make(map[rune]bool)

	// Iterate over the characters of the first string
	for _, char := range os.Args[1] {
		encountered[char] = true
	}

	// Iterate over the characters of the second string
	for _, char := range os.Args[2] {
		encountered[char] = true
	}

	// Print the unique characters in the order they appear
	for _, char := range os.Args[1] {
		if encountered[char] {
			z01.PrintRune(char)
			encountered[char] = false // Mark the character as printed
		}
	}

	// Iterate over the characters of the second string
	for _, char := range os.Args[2] {
		if encountered[char] {
			z01.PrintRune(char)
			encountered[char] = false // Mark the character as printed
		}
	}

	// Print a newline
	z01.PrintRune('\n')
}
