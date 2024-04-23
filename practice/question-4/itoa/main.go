// https://github.com/01-edu/public/tree/master/subjects/itoa

package main

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var result string
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	for n > 0 {
		digit := n % 10
		result = string('0'+rune(digit)) + result
		n /= 10
	}

	return sign + result
}

func main() {
	// Test the function
	println(Itoa(123))  // Output: "123"
	println(Itoa(-456)) // Output: "-456"
	println(Itoa(0))    // Output: "0"
	println(Itoa(12345))
	println(Itoa(-1234))
	println(Itoa(987654321))
}
