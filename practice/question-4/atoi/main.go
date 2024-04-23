// https://github.com/01-edu/public/tree/master/subjects/atoi
package main

import (
	"fmt"
)

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}

func Atoi(s string) int {

	/*
		First, check the sign. Store the sign
		Check if the input is only numbers. Return 0
		now, collect the numbers
		create int
		multiply it with sign
		Tadaaa!!!!
	*/

	sr := []rune(s)
	var sign int = 1
	var answer int

	if sr[0] == '+' {
		sign = 1
		sr = sr[1:]
	}

	if sr[0] == '-' {
		sign = -1
		sr = sr[1:]
	}

	for i := 0; i < len(sr); i++ {
		if sr[i] < '0' || sr[i] > '9' {
			return 0
		}

		answer = answer*10 + int(sr[i]-'0')
	}

	answer = answer * sign

	return answer
}
