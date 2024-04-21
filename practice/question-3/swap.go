// https://github.com/01-edu/public/tree/master/subjects/swap

package piscine

func Swap(a *int, b *int) {
	// Temporarily store the value pointed to by a
	temp := *a
	// Assign the value pointed to by b to the memory location pointed to by a
	*a = *b
	// Assign the temporarily stored value to the memory location pointed to by b
	*b = temp
}
