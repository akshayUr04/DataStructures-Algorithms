package search

func LinearSearch(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i // Found, return the index
		}
	}
	return -1 // Not found
}
