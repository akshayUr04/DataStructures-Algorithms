package tree

func (t *Tree) BfsInorder() []int {
	if t.Root == nil {
		return nil
	}
	var result []int
	t.Root.bfsInorder(&result)
	return result
}

func (n *Node) bfsInorder(result *[]int) {
	if n.Left != nil {
		n.Left.bfsInorder(result)
	}
	*result = append(*result, n.Key)
	if n.Right != nil {
		n.Right.bfsInorder(result)
	}
}
