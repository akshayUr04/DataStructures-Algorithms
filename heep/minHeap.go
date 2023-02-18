package heep

type MinHeap struct {
	array []int
}

func (min *MinHeap) Insert(val int) {
	min.array = append(min.array, val)
	min.minHeapifyUp(len(min.array) - 1)
}

func (min *MinHeap) Extract() int {
	l := len(min.array) - 1
	extracted := min.array[l]
	min.array[0] = min.array[l]
	min.array = min.array[:l]
	min.minHeapifydown(0)
	return extracted
}

func (min *MinHeap) minHeapifyUp(index int) {
	for min.array[parent(index)] > min.array[index] {
		min.swap(parent(index), index)
		index = parent(index)
	}
}

func (min *MinHeap) minHeapifydown(index int) {
	lastIndex := len(min.array) - 1
	l, r := leftChild(index), rightChild(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if min.array[l] < min.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if min.array[childToCompare] < min.array[index] {
			min.swap(childToCompare, index)
			index = childToCompare
			l, r = leftChild(index), rightChild(index)
		} else {
			return
		}
	}
}

func (min *MinHeap) swap(i1, i2 int) {
	min.array[i1], min.array[i2] = min.array[i2], min.array[i1]
}
