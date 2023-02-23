package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (l *LinkedList) Prepend(value int) { //inserting at the begining
	newNode := Node{data: value}
	if l.head != nil {
		newNode.next = l.head //set the next pointer of the new node to current head
		l.head = &newNode     //set the head pointer to the new node
		l.length++
	} else {
		l.head = &newNode
		l.tail = &newNode
		l.length++
	}

}

func (l *LinkedList) Add(value int) {
	newNode := Node{data: value}
	if l.tail != nil {
		l.tail.next = &newNode
		l.tail = &newNode
	} else {
		l.head = &newNode
		l.tail = &newNode
	}
	l.length++

}

func (l *LinkedList) AddAfter(nextTo int, value int) {
	newNode := &Node{data: value}
	temp := l.head
	for temp.data != nextTo && temp != nil {
		temp = temp.next
	}
	if temp == nil {
		return
	}
	if temp == l.tail {
		l.tail.next = newNode
		l.tail = newNode
		l.length++
		return
	}
	newNode.next = temp.next
	temp.next = newNode
	l.length++
}

func (l *LinkedList) AddBefore(before, value int) {
	newNode := &Node{data: value}
	if l.head.data == before {
		newNode.next = l.head
		l.head = newNode
		l.length++
		return
	}
	temp := l.head
	for temp != nil {
		if temp.next.data == before && temp.next != nil {
			newNode.next = temp.next
			temp.next = newNode
			l.length++
			return
		}
		temp = temp.next
	}
	if temp == nil {
		return
	}

}

func (l *LinkedList) PrintList() {
	if l.head == nil {
		return
	}
	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

func (l *LinkedList) deletewithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDlt := l.head
	for previousToDlt.next.data != value {
		if previousToDlt.next.next == nil {
			return
		}
		previousToDlt = previousToDlt.next
	}
	if l.tail.data == value {
		l.tail = previousToDlt
		l.length--
	}
	previousToDlt.next = previousToDlt.next.next
	l.length--
}

func (l *LinkedList) AddAfterPosition(pos, value int) {
	newNode := &Node{data: value}
	temp := l.head
	for i := 1; i < pos; i++ {
		temp = temp.next
	}
	newNode.next = temp.next
	temp.next = newNode
	l.length++
}

func (l *LinkedList) AddBeforePosition(pos, value int) {
	newNode := &Node{data: value}
	temp := l.head
	for i := 1; i < pos-1; i++ {
		temp = temp.next
	}
	newNode.next = temp.next
	temp.next = newNode
	l.length++

}

func (l *LinkedList) DltFromPosition(pos int) {

	if l.length == 0 {
		return
	}
	if l.length < pos {
		return
	}
	if pos == 1 {
		l.head = l.head.next
		l.length--
		return
	}
	temp := l.head
	for i := 1; i < pos-1; i++ {
		temp = temp.next
	}
	if temp.next == l.tail {
		temp.next = nil
		l.tail = temp
		l.length--
		return
	}

	temp.next = temp.next.next
	l.length--

}

func (l *LinkedList) lLargestValues(val LinkedList) {
	temp := val.head.data
	head := val.head

	for head != nil {
		if head.data > temp {
			temp = head.data
		}
		if head.next == nil {
			break
		}
		head = head.next
	}
	fmt.Println(temp)

}

func (l *LinkedList) ArrayToLinkedList(arr []int) {
	for _, val := range arr {
		l.Add(val)
	}

}

func (l *LinkedList) Reverse() {
	var prev *Node
	current := l.head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
}
