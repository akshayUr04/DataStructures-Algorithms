package heep

import (
	"fmt"
)

type MaxHeap struct {
	array []int
}

func (m *MaxHeap) Insert(val int) {
	m.array = append(m.array, val)
	m.maxHeapifyUp(len(m.array) - 1)
}

func (m *MaxHeap) Extract() int {
	extracted := m.array[0]
	if len(m.array) == 0 {
		fmt.Println("array is empty")
		return -1
	}
	l := len(m.array) - 1
	m.array[0] = m.array[l]
	m.array = m.array[:l]
	m.maxHeapifyDown(0)
	return extracted
}

func (m *MaxHeap) maxHeapifyUp(index int) {
	for m.array[parent(index)] < m.array[index] {
		m.swap(parent(index), index)
		index = parent(index)
	}
}

func (m *MaxHeap) maxHeapifyDown(index int) {

	lastIndex := len(m.array) - 1
	l, r := leftChild(index), rightChild(index)
	childToCompare := 0

	// loop while index has least one child
	for l <= lastIndex { //when left child is the only child
		if l == lastIndex {
			childToCompare = l
		} else if m.array[l] > m.array[r] { //when left child is greter than right child
			childToCompare = l
		} else { //when right child is larger
			childToCompare = r
		}
		//compare value of current index to larger child amd swap if smaller
		if m.array[index] < m.array[childToCompare] {
			m.swap(index, childToCompare)
			index = childToCompare
			l, r = leftChild(index), rightChild(index)
		} else {
			return
		}
	}

}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func (m *MaxHeap) swap(i1, i2 int) {
	m.array[i1], m.array[i2] = m.array[i2], m.array[i1]
}
