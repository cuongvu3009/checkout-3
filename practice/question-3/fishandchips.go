// https://github.com/01-edu/public/tree/master/subjects/fishandchips

package piscine

func FishAndChips(n int) string {
	if n < 0 {
		return "error: number is negative"
	}

	if (n%2 == 0) && (n%3 == 0) {
		return "fish and chips"
	}
	if n%2 == 0 {
		return "fish"
	}
	if n%3 == 0 {
		return "chips"
	}
	return "error: non divisible"
}
