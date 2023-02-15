package stack

import "fmt"

type ArrayStack struct {
	Item []int
	Top  int
}

func (a *ArrayStack) Push(val int) {
	if a.Top == len(a.Item)-1 {
		fmt.Println("stack is full")
		return
	}
	a.Top++
	a.Item[a.Top] = val
}

func (a *ArrayStack) Pop() int {
	if a.Top == -1 {
		fmt.Println("stack is empty")
		return -1
	}
	value := a.Item[a.Top]
	a.Item = append(a.Item[:a.Top])
	a.Top--
	return value
}
