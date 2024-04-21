// https://github.com/01-edu/public/tree/master/subjects/printreversealphabet

package piscine

import (
	"github.com/01-edu/z01"
)

func Print() {
	for i := 'z'; i >= 'a'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
