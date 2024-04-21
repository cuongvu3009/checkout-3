// https://github.com/01-edu/public/tree/master/subjects/wdmatch

package main

import (
	"os"
)

func main() {
	// Check if there are exactly two command-line arguments
	if len(os.Args) != 3 {
		return
	}

	// Extract the first and second strings from command-line arguments
	firstString := os.Args[1]
	secondString := os.Args[2]

	// Check if it is possible to write the first string with characters from the second string
	if canWriteWithCharacters(firstString, secondString) {
		// If possible, display the first string followed by a newline
		println(firstString)
	}
}

// Function to check if it is possible to write the first string with characters from the second string
func canWriteWithCharacters(first, second string) bool {
	// Initialize a variable to track the index in the second string
	index := 0

	// Iterate over each character of the first string
	for _, char := range first {
		// Search for the current character in the second string
		found := false
		for index < len(second) {
			if rune(second[index]) == char {
				found = true
				break
			}
			index++
		}

		// If the current character is not found in the second string, return false
		if !found {
			return false
		}
	}

	// If all characters of the first string are found in the second string in the correct order, return true
	return true
}
