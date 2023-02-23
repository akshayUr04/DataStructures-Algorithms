package recursion

func Binary(arr []int, val int) int {
	star := 0
	end := len(arr) - 1
	mid := (star + end) / 2
	if val == arr[mid] {
		return mid + 1
	} else if arr[mid] < val {
		return Binary(arr[mid+1:], val)
	} else {
		return Binary(arr[:mid-1], val)
	}
}
