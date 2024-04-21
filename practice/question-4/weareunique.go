// https://github.com/01-edu/public/tree/master/subjects/weareunique

package piscine

func WeAreUnique(str1, str2 string) int {
	// Check if both strings are empty
	if str1 == "" && str2 == "" {
		return -1
	}

	// Create a map to store unique characters from both strings
	charMap := make(map[rune]bool)

	// Iterate over the characters in str1 and add them to the map
	for _, char := range str1 {
		charMap[char] = true
	}

	// Iterate over the characters in str2 and remove them from the map if they exist
	for _, char := range str2 {
		delete(charMap, char)
	}

	// Return the number of unique characters left in the map
	return len(charMap)
}
