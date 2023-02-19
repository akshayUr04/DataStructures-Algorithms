package tree

func (t *Tree) BfsPreOrder() []int {
	if t.Root == nil {
		return nil
	}
	var result []int
	t.Root.bfsPreOrder(&result)
	return result
}

func (n *Node) bfsPreOrder(result *[]int) {
	*result = append(*result, n.Key)
	if n.Left != nil {
		n.Left.bfsPreOrder(result)
	}
	if n.Right != nil {
		n.Right.bfsPreOrder(result)
	}
}
