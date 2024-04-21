// https://github.com/01-edu/public/tree/master/subjects/slice

package piscine

func Slice(a []string, nbrs ...int) []string {
	// Check if the number of integers is valid
	if len(nbrs) == 0 {
		return nil
	}

	// Extract the start and end indices from the integers
	start := nbrs[0]
	end := len(a)

	if len(nbrs) > 1 {
		end = nbrs[1]
	}

	// Handle negative indices
	if start < 0 {
		start += len(a)
	}
	if end < 0 {
		end += len(a)
	}

	// Clamp start and end indices to valid range
	start = max(0, min(start, len(a)))
	end = max(0, min(end, len(a)))

	// Return the sliced slice
	return a[start:end]
}

// Helper function to return the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function to return the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
