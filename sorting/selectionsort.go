package sorting

func SelectinSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		index := i
		for j := i + 1; j < len(arr); j++ {
			if arr[index] > arr[j] {
				index = j
			}
		}
		arr[i], arr[index] = arr[index], arr[i]
	}
	return arr
}
