// https://github.com/01-edu/public/tree/master/subjects/printmemory

package main

import (
	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	// Print hexadecimal memory representation
	for i, c := range arr {
		hexDigit := c >> 4
		z01.PrintRune(hexDigitToRune(hexDigit))
		hexDigit = c & 0x0F
		z01.PrintRune(hexDigitToRune(hexDigit))

		z01.PrintRune(' ')

		if i%4 == 3 {
			z01.PrintRune('\n')
		}

	}
	z01.PrintRune('\n')

	// Print ASCII characters
	for i := 0; i < len(arr); i++ {
		char := arr[i]
		if char >= 32 && char <= 126 {
			z01.PrintRune(rune(char))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func hexDigitToRune(digit byte) rune {
	if digit < 10 {
		return rune('0' + digit)
	}
	return rune('A' + digit - 10)
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
