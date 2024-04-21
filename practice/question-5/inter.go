//https://github.com/01-edu/public/tree/master/subjects/inter

package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func Inter() {
	// Check if the number of arguments is not equal to 3 (including the program name)
	if len(os.Args) != 3 {
		return
	}

	// Extract the two strings from the command-line arguments
	str1 := os.Args[1]
	str2 := os.Args[2]

	// Create a map to store the characters of the first string
	seen := make(map[rune]bool)

	// Iterate over the characters of the first string and mark them as seen
	for _, char := range str1 {
		seen[char] = true
	}

	// Iterate over the characters of the second string and print those that appear in the first string
	for _, char := range str2 {
		if seen[char] {
			// Print the character if it's present in the first string
			z01.PrintRune(char)
			// Mark the character as seen to avoid duplicates
			delete(seen, char)
		}
	}

	// Print a newline at the end
	z01.PrintRune('\n')
}
