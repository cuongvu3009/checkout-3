// https://github.com/01-edu/public/tree/master/subjects/countcharacter

package piscine

func CountCharacter(str string, c rune) int {
	count := 0
	for _, char := range str {
		if char == c {
			count++
		}
	}
	return count
}
