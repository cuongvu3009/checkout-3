// https://github.com/01-edu/public/tree/master/subjects/fprime

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if the number of arguments is exactly 2 (including the program name)
	if len(os.Args) != 2 {
		return
	}

	// Parse the argument as an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 1 {
		return
	}

	// Find and display the prime factors
	primeFactors := findPrimeFactors(n)
	if len(primeFactors) > 0 {
		fmt.Println(joinWithAsterisk(primeFactors))
	}
}

// Function to find prime factors of a positive integer
func findPrimeFactors(n int) []int {
	factors := make([]int, 0)

	// Find all factors of 2
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	// Find remaining prime factors
	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	// If n is a prime number greater than 2
	if n > 2 {
		factors = append(factors, n)
	}

	return factors
}

// Function to join integers with asterisks
func joinWithAsterisk(nums []int) string {
	result := ""
	for i, num := range nums {
		if i > 0 {
			result += "*"
		}
		result += strconv.Itoa(num)
	}
	return result
}
