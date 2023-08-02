package search

func BinarSearch(arr []int, target int) int {
	first := 0
	last := len(arr) - 1
	for first <= last {
		mid := first + last
		if target == arr[mid] {
			return mid
		} else if target > arr[mid] {
			first = mid + 1
		} else {
			last = mid - 1
		}
	}
	return -1
}
