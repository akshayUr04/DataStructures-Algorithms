package linkedlist

import "fmt"

type DNode struct {
	prev *DNode
	data int
	next *DNode
}
type DoubleLinkedList struct {
	head   *DNode
	tail   *DNode
	length int
}

func (dl *DoubleLinkedList) AddAtBigin(value int) {
	newNode := &DNode{data: value}
	if dl.head == nil && dl.tail == nil {
		dl.head = newNode
		dl.tail = newNode
	} else {
		newNode.next = dl.head
		dl.head.prev = newNode
		dl.head = newNode

	}
	dl.length++
}

func (dl *DoubleLinkedList) AddAtEnd(value int) {
	newNode := &DNode{data: value}
	if dl.head == nil && dl.tail == nil {
		dl.head = newNode
		dl.tail = newNode
	} else {
		newNode.prev = dl.tail
		dl.tail.next = newNode
		dl.tail = newNode
	}
	dl.length++
}

func (dl *DoubleLinkedList) Display() {
	temp := dl.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}
func (dl *DoubleLinkedList) DisplayReverse() {
	temp := dl.tail
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.prev
	}
}

func (dl *DoubleLinkedList) DltAtBegin() {
	if dl.head == dl.tail {
		dl.head = nil
		dl.tail = nil
	} else {
		dl.head = dl.head.next
		dl.head.prev = nil
	}
	dl.length--
}
func (dl *DoubleLinkedList) DltAtEnd() {
	if dl.tail == dl.head {
		dl.head = nil
		dl.tail = nil
	} else {
		dl.tail = dl.tail.prev
		dl.tail.next = nil
	}
	dl.length--
}

func (dl *DoubleLinkedList) dltWithVal(val int) {
	temp := dl.head
	if temp == nil {
		return
	}
	if dl.head.data == val {
		dl.head = dl.head.next
		dl.head.next.prev = nil
		dl.length--
		return
	}
	for temp.next.data != val {
		if temp.next.next == nil {
			return
		}
		temp = temp.next
	}
	if temp.next == dl.tail {
		dl.tail = temp
		dl.tail.next = nil
		dl.length--
		return
	}
	temp.next = temp.next.next
	temp.next.prev = temp
	dl.length--
}
