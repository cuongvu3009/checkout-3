// https://github.com/01-edu/public/tree/master/subjects/printrevcomb

package piscine

import (
	"github.com/01-edu/z01"
)

func PrintRevComb() {
	// Iterate through all possible combinations of three different digits
	for i := '9'; i >= '0'; i-- {
		for j := i - 1; j >= '0'; j-- {
			for k := j - 1; k >= '0'; k-- {
				// Print the combination if it meets the specified conditions
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)

				// Check if it's not the last combination, then print comma and space
				if i != '9' || j != '1' || k != '0' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
