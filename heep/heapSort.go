package heep

import "fmt"

func (mx *MaxHeap) HeapSort(arr []int) {
	fmt.Println(mx.array)
	n := len(mx.array)
	for i := 0; i < n; i++ {
		max := mx.Extract()
		arr[n-i-1] = max
		mx.array = mx.array[:n-i-1]
		mx.maxHeapifyDown(0)
	}
}
