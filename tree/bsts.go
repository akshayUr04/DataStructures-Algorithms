package tree

import "fmt"

type Node struct {
	Left  *Node
	Key   int
	Right *Node
}

type Tree struct {
	Root *Node
}

func (n *Node) Insert(val int) {
	if n.Key < val {
		if n.Right == nil {
			n.Right = &Node{Key: val}
		} else {
			n.Right.Insert(val)
		}
	} else if n.Key > val {
		if n.Left == nil {
			n.Left = &Node{Key: val}
		} else {
			n.Left.Insert(val)
		}
	}
}

func (n *Node) Search(val int) bool {
	if n == nil {
		return false
	}
	if n.Key < val {
		return n.Right.Search(val)
	} else if n.Key > val {
		return n.Left.Search(val)
	}
	return true
}

func (n *Node) Delete(val int) *Node {
	if n == nil {
		return nil
	}
	fmt.Println("node", *n)

	//search for the node to dlt
	fmt.Println(n.Key)
	if val < n.Key {
		n.Left = n.Left.Delete(val)
		fmt.Println("nL", *n)
		return n
	} else if val > n.Key {
		n.Right = n.Right.Delete(val)
		fmt.Println("nR", *n)
		return n
	}

	//if node has no childeren dlt it
	if n.Left == nil && n.Right == nil {
		return nil
	}

	//if the node has one childe replace it with child node
	if n.Left == nil {
		return n.Right
	} else if n.Right == nil {
		return n.Left
	}

	// if the node has 2 child node find the successor node
	successor := n.Right
	for successor.Left != nil {
		fmt.Println(successor.Key)
		successor = successor.Left
	}

	// copy the successor's value into the current node
	n.Key = successor.Key

	//delete the successor node by recursivly calling the delete with successor key
	n.Right = n.Right.Delete(successor.Key)

	return n

}

func (t *Tree) Insert(k int) {
	if t.Root == nil {
		t.Root = &Node{Key: k}
	} else {
		t.Root.Insert(k)
	}
}

func (t *Tree) Search(val int) bool {
	return t.Root.Search(val)
}

func (t *Tree) Delete(val int) {
	t.Root = t.Root.Delete(val)
}

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
