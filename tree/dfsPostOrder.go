package tree

func (t *Tree) BfsPostOrder() []int {
	if t.Root == nil {
		return nil
	}
	var result []int
	t.Root.bfsPostOrder(&result)
	return result
}

func (n *Node) bfsPostOrder(result *[]int) {
	if n.Left != nil {
		n.Left.bfsPostOrder(result)
	}
	if n.Right != nil {
		n.Right.bfsPostOrder(result)
	}

	*result = append(*result, n.Key)
}
