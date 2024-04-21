// https://github.com/01-edu/public/tree/master/subjects/concatslice

package piscine

func ConcatSlice(slice1, slice2 []int) []int {
	var result []int
	for _, v := range slice1 {
		result = append(result, v)
	}
	for _, v := range slice2 {
		result = append(result, v)
	}
	return result
}
