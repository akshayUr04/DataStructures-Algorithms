package stack

import "fmt"

type node struct {
	value int
	next  *node
}

type LinkedList struct {
	Head   *node
	length int
}

func (l *LinkedList) Push(val int) {
	newNode := &node{value: val}
	if l.Head == nil {
		l.Head = newNode
	} else {
		newNode.next = l.Head
		l.Head = newNode
	}
	l.length++
}

func (l *LinkedList) Pop() (int, bool) {
	if l.length == 0 {
		fmt.Println("stack is empty")
		return -1, false
	}
	val := l.Head.value
	l.Head = l.Head.next
	l.length--
	return val, true
}

func (l *LinkedList) Display() {

	temp := l.Head
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}

}
