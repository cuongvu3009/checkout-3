// https://github.com/01-edu/public/tree/master/subjects/findprevprime

package main

import "fmt"

func main() {

	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}

func FindPrevPrime(nb int) int {
	for i := nb; i > 1; i-- {
		if isPrime(i) {
			return i
		}
	}
	return 0
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}

	}
	return true
}
