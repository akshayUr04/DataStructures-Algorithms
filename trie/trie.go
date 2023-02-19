package trie

const Alphabets = 26

type Node struct {
	Children [Alphabets]*Node
	isEnd    bool
}

type Trie struct {
	Root *Node
}

// This create a new Trie with a empty node
func InitTrie() *Trie {
	result := &Trie{Root: &Node{}}
	return result
}

func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.Root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.Children[charIndex] == nil {
			currentNode.Children[charIndex] = &Node{}
		}
		currentNode = currentNode.Children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.Root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.Children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.Children[charIndex]
	}
	return currentNode.isEnd
}
