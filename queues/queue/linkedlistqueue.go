package queue

import "fmt"

type Node struct {
	data int
	next *Node
}

type Heap struct {
	Head   *Node
	length int
}

// enqueue

func (h *Heap) LinkedListEnqueue(val int) {
	newNode := &Node{data: val}
	temp := h.Head
	if h.Head == nil {
		h.Head = newNode
		h.length++
		return
	}
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
	h.length++

}

// dequeue

func (h *Heap) LinkedListDequeue() (int, bool) {
	if h.length == 0 {
		fmt.Println("queue is empty")
		return -1, false
	}
	valToDlt := h.Head.data
	h.Head = h.Head.next
	return valToDlt, true
}

//Display

func (h *Heap) Display() {
	temp := h.Head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}
