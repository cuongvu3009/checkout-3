// https://github.com/01-edu/public/tree/master/subjects/fromto

package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	var result string

	for i := from; i <= to; i++ {
		if i < 10 {
			z01.PrintRune('0')
		}
		for _, digit := range fmt.Sprintf("%d", i) {
			z01.PrintRune(digit)
		}
		if i != to {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')

	return result
}
