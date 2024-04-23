// https://github.com/01-edu/public/tree/master/subjects/printmemory

package main

import (
	"github.com/01-edu/z01"
)

func PrintMemory(arr [10]byte) {
	// Print hexadecimal memory representation

	//First part
	for i, c := range arr {

		hexDigit1 := c / 16
		hexDigit2 := c % 16

		z01.PrintRune(hexMapper(hexDigit1))
		z01.PrintRune(hexMapper(hexDigit2))

		z01.PrintRune(' ')

		if i%4 == 3 {
			z01.PrintRune('\n')
		}
	}

	z01.PrintRune('\n')
	printInputCustom(arr)
}

func PrintString(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

// print input string as required
func printInputCustom(arr [10]byte) {
	for _, char := range arr {
		if char >= ' ' && char <= '~' {
			z01.PrintRune(rune(char))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func hexMapper(digit byte) rune {
	switch digit {
	case 0:
		return '0'
	case 1:
		return '1'
	case 2:
		return '2'
	case 3:
		return '3'
	case 4:
		return '4'
	case 5:
		return '5'
	case 6:
		return '6'
	case 7:
		return '7'
	case 8:
		return '8'
	case 9:
		return '9'
	case 10:
		return 'a'
	case 11:
		return 'b'
	case 12:
		return 'c'
	case 13:
		return 'd'
	case 14:
		return 'e'
	case 15:
		return 'f'
	default:
		return ' '
	}
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

// Output:
// 68 65 6c 6c$
// 6f 10 15 2a$
// 00 00$
// hello..*..$
