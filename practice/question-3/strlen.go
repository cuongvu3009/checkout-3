// https://github.com/01-edu/public/tree/master/subjects/strlen

package piscine

func StrLen(s string) int {
	return len([]rune(s))
}
