package stack

import (
	"fmt"
)

type SliceStack struct {
	Top   int
	items []int
}

func (s *SliceStack) SlicePush(i int) {
	s.Top++
	s.items = append(s.items, i)
}

func (s *SliceStack) SlicePop() int {
	if s.Top == -1 {
		fmt.Println("stack is empty")
		return -1
	}
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}
