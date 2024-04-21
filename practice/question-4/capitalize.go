// https://github.com/01-edu/public/tree/master/subjects/capitalize

package piscine

func isAlphaNumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func Capitalize(s string) string {
	runes := []rune(s)
	capitalizeNext := true
	for i, r := range runes {
		if isAlphaNumeric(r) {
			if capitalizeNext {
				if r >= 'a' && r <= 'z' {
					runes[i] = r - ('a' - 'A')
				}
				capitalizeNext = false
			} else {
				if r >= 'A' && r <= 'Z' {
					runes[i] = r + ('a' - 'A')
				}
			}
		} else {
			capitalizeNext = true
		}
	}
	return string(runes)
}
