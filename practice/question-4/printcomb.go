// https://github.com/01-edu/public/tree/master/subjects/printcomb

package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
				z01.PrintRune(rune(i + '0'))
				z01.PrintRune(rune(j + '0'))
				z01.PrintRune(rune(k + '0'))
				if i != 7 || j != 8 || k != 9 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
