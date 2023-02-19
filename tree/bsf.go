package tree

func (t *Tree) Bsf() []int {
	if t.Root == nil {
		return []int{}
	}
	result := make([]int, 0)
	queue := []*Node{t.Root}
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		result = append(result, curNode.Key)

		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
		}

		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
		}
	}
	return result
}
