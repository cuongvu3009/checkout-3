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

	slicy := []rune(s)
	var sign int = 1
	var answer int

	if slicy[0] == '+' {
		sign = 1
		slicy = slicy[1:]
	}

	if slicy[0] == '-' {
		sign = -1
		slicy = slicy[1:]
	}

	for i := 0; i < len(slicy); i++ {
		if slicy[i] < '0' || slicy[i] > '9' {
			return 0
		}

		answer = answer*10 + int(slicy[i]-'0')
	}

	answer = answer * sign

	return answer
}
