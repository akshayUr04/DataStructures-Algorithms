package hashtable

import "fmt"

const arraySize = 7

// Hash table will hold an array
type HashTable struct {
	arr [arraySize]*Bucket
}

// linked list node that holds the key
type BucketNode struct {
	key  string
	next *BucketNode
}

type Bucket struct {
	head *BucketNode
}

// inserting in hash table
func (h *HashTable) HashInsert(key string) {
	index := hash(key)
	h.arr[index].BucketInsert(key)
}

// deleting in hash table
func (h *HashTable) HashDlet(key string) {
	index := hash(key)
	h.arr[index].BucketDlt(key)

}

// searching in hash table
func (h *HashTable) HashSearch(key string) bool {
	index := hash(key)
	return h.arr[index].BucketSearch(key)
}

// insert will take in a key, create a node with the key and insert the node in the bucket
func (b *Bucket) BucketInsert(k string) {
	if !b.BucketSearch(k) {
		newNode := &BucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("alredy exists")
	}

}

func (b *Bucket) BucketSearch(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *Bucket) BucketDlt(k string) {
	previousNode := b.head

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	for previousNode.next != nil {
		if previousNode.next.next == nil {
			return
		}
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
			return
		}
		previousNode = previousNode.next
	}

}

func hash(key string) int {
	sum := 0
	for _, val := range key {
		sum += int(val)
	}
	return sum % arraySize
}

// init will create a bucket in each slot of the table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.arr {
		result.arr[i] = &Bucket{}
	}
	return result
}
