// https://github.com/01-edu/public/tree/master/subjects/findprevprime

package main

import "fmt"

func main() {

	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}

func FindPrevPrime(nb int) int {
	for i := nb; i > 1; i-- {
		if isprime(i) {
			return i
		}
	}

	return 0
}

func isprime(n int) bool {
	if n < 2 || n%2 == 0 {
		return false
	}

	if n == 2 {
		return true
	}

	for i := 3; i*i < n; i = i + 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
