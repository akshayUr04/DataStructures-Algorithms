package sorting

func MergeSort(arr []int) []int {
	//Base condition : If the size of the array is 1, return the array back
	if len(arr) <= 1 {
		return arr
	}
	//splitting the array into two from mid and using merge sort again
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

// merge function helps to merge two sorted arrays
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	//If anything is left, appending them to the result array. Only one of the below lines of code will get executed
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// MergeSort([5, 4, 10, 6, 2, 1])
// 	-> MergeSort([5, 4, 10])
// 		-> MergeSort([5])
// 			return [5]
// 		-> MergeSort([4, 10])
// 			-> MergeSort([4])
// 				return [4]
// 			-> MergeSort([10])
// 				return [10]
// 			-> merge([4], [10])
// 				return [4, 10]
// 		-> merge([5], [4, 10])
// 			return [4, 5, 10]
// 	-> MergeSort([6, 2, 1])
// 		-> MergeSort([6])
// 			return [6]
// 		-> MergeSort([2, 1])
// 			-> MergeSort([2])
// 				return [2]
// 			-> MergeSort([1])
// 				return [1]
// 			-> merge([2], [1])
// 				return [1, 2]
// 		-> merge([6], [1, 2])
// 			return [1, 2, 6]
// 	-> merge([4, 5, 10], [1, 2, 6])
// 		return [1, 2, 4, 5, 6, 10]
