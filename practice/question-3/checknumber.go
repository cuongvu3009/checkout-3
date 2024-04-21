// https://github.com/01-edu/public/tree/master/subjects/checknumber

package piscine

func CheckNumber(str string) bool {
	for _, c := range str {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}
