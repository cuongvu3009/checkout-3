// https://github.com/01-edu/public/tree/master/subjects/concatslice

package piscine

func ConcatSlice(slice1, slice2 []int) []int {
	// Create a new slice to store the concatenation
	concatenated := make([]int, 0, len(slice1)+len(slice2))

	// Append elements from the first slice
	concatenated = append(concatenated, slice1...)

	// Append elements from the second slice
	concatenated = append(concatenated, slice2...)

	return concatenated
}
