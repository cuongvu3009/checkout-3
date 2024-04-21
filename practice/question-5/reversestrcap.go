// https://github.com/01-edu/public/tree/master/subjects/reversestrcap

package piscine

import (
	"os"
	"strings"

	"github.com/01-edu/z01"
)

func ReverseStrCap() {
	if len(os.Args) < 2 {
		return
	}

	for _, arg := range os.Args[1:] {
		words := strings.Fields(arg)
		for _, word := range words {
			for i, char := range word {
				if i == len(word)-1 {
					z01.PrintRune(char - 'a' + 'A') // Convert the last letter to uppercase
				} else {
					z01.PrintRune(char) // Print other letters as they are
				}
			}
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')
}
