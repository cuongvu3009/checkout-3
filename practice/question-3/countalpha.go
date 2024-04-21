// https://github.com/01-edu/public/tree/master/subjects/countalpha

package piscine

func CountAlpha(s string) int {
	count := 0
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			count++
		}
	}
	return count
}
