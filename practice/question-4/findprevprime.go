// https://github.com/01-edu/public/tree/master/subjects/findprevprime

package piscine

func FindPrevPrime(nb int) int {
	// Edge case: if nb is less than 2, there is no prime number less than nb
	if nb < 2 {
		return 0
	}

	// Start iterating downward from nb
	for i := nb; i >= 2; i-- {
		// Check if the current number is prime
		if isPrime(i) {
			return i
		}
	}

	// If no prime number is found, return 0
	return 0
}

// Function to check if a number is prime
func isPrime(n int) bool {
	// Edge case: 1 is not a prime number
	if n == 1 {
		return false
	}

	// Check if n is divisible by any number from 2 to n-1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
