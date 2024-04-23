// https://github.com/01-edu/public/tree/master/subjects/iscapitalized

package piscine

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' && s[i-1] == ' ' {
			return false
		}
	}
	return true
}
