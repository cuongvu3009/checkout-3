// https://github.com/01-edu/public/tree/master/subjects/countcharacter

package piscine

func CountChar(str string, c rune) int {
	var count int
	for _, v := range str {
		if rune(v) == c {
			count++
		}
	}
	return count
}
