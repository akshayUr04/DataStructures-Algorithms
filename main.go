package main

import "golang-dsa/graph"

func main() {

	// -------------------------------------------------Sorting----------------------------------------------------------
	// arr := []int{5, 4, 10, 2, 3}
	// fmt.Println("array before sorting", arr)
	// a := len(arr) - 1
	// b := 0

	// fmt.Println("insertion sort", sorting.InsertionSort(arr))
	// fmt.Println("bouble sort", sorting.BoubleSort(arr))
	// fmt.Println("selction sort", sorting.SelectinSort(arr))
	// sorting.QuickSort(arr, b, a)
	// fmt.Println("MergeSort sort", sorting.MergeSort(arr))
	// fmt.Println(arr)

	// --------------------------------------------------HashTable----------------------------------------------------------

	// testBucket := &hashtable.Bucket{}
	// testBucket.BucketInsert("akshay")
	// testBucket.BucketInsert("hallo")
	// fmt.Println(testBucket.BucketSearch("akshay"))
	// testBucket.BucketDlt("akshay")
	// fmt.Println(testBucket.BucketSearch("akshay"))

	// Hashtable := hashtable.Init()
	// list := []string{
	// 	"Aswathi",
	// 	"Akshay",
	// 	"kshayA",
	// }

	// for _, val := range list {
	// 	Hashtable.HashInsert(val)
	// }
	// Hashtable.HashDlet("kshayA")
	// fmt.Println(Hashtable.HashSearch("Akshay"))
	// fmt.Println(Hashtable.HashSearch("Unni"))

	// -------------------------------------------------------Stack----------------------------------------------------------

	//Slice

	// myStack := &stack.SliceStack{Top: 10}
	// fmt.Println(myStack)
	// myStack.SlicePush(10)
	// myStack.SlicePush(11)
	// myStack.SlicePush(12)
	// myStack.SlicePush(13)
	// fmt.Println(myStack)
	// fmt.Println(myStack.SlicePop())

	//Array

	// maxSize := 4
	// myStackArray := stack.ArrayStack{Item: make([]int, maxSize), Top: -1}
	// myStackArray.Push(10)
	// myStackArray.Push(20)
	// myStackArray.Push(30)
	// myStackArray.Push(40)
	// fmt.Println(myStackArray.Item)
	// fmt.Println(myStackArray.Pop())
	// fmt.Println(myStackArray.Item)

	//LinkedList

	// newStack := &stack.LinkedList{}
	// newStack.Push(10)
	// newStack.Push(11)
	// newStack.Push(12)
	// newStack.Push(13)
	// newStack.Push(14)
	// newStack.Push(15)
	// newStack.Push(16)
	// newStack.Display()
	// fmt.Println("---------------")
	// fmt.Println(newStack.Pop())
	// fmt.Println("---------------")
	// newStack.Display()

	// --------------------------------------------------------Queue----------------------------------------------------------

	//Slice

	// newQueue := &queue.Queue{}
	// newQueue.SliceEnque(11)
	// newQueue.SliceEnque(12)
	// newQueue.SliceEnque(13)
	// newQueue.SliceEnque(14)
	// newQueue.SliceEnque(15)
	// newQueue.SliceEnque(16)
	// newQueue.SliceEnque(17)
	// fmt.Println(newQueue)
	// fmt.Println(newQueue.SliceDequeue())
	// fmt.Println(newQueue.SliceDequeue())
	// fmt.Println(newQueue.SliceDequeue())
	// fmt.Println(newQueue.SliceDequeue())
	// fmt.Println(newQueue)

	//Linked list

	// newQueue := queue.Heap{}
	// newQueue.LinkedListEnqueue(10)
	// newQueue.LinkedListEnqueue(11)
	// newQueue.LinkedListEnqueue(12)
	// newQueue.LinkedListEnqueue(13)
	// newQueue.Display()
	// fmt.Println(newQueue.LinkedListDequeue())
	// newQueue.Display()

	//Array
	// queue := queue.Init(4)
	// fmt.Println(queue)
	// queue.Enqueue(10)
	// queue.Enqueue(12)
	// queue.Enqueue(13)
	// queue.Enqueue(14)
	// queue.Display()
	// fmt.Println(queue)
	// queue.Dequeue()
	// queue.Display()
	// fmt.Println(queue)
	// queue.Enqueue(15)

	// circular := circularqueue.Init(4)
	// circular.Enqueue(11)
	// circular.Enqueue(12)
	// circular.Enqueue(13)
	// circular.Enqueue(14)
	// fmt.Println(circular)
	// circular.Dequeue()
	// fmt.Println(circular)
	// circular.Enqueue(15)
	// fmt.Println(circular)

	// ---------------------------------------------------Tree-------------------------------------------------------------------

	// tree := &tree.Tree{}
	// values := []int{10, 15, 12, 8, 13, 19, 18, 11}
	// for _, val := range values {
	// 	tree.Insert(val)
	// }
	// fmt.Println(tree.Search(10))
	// fmt.Println(tree.Search(485))
	// tree.Delete(15)
	// fmt.Println(tree.Bsf())
	// fmt.Println(tree.BfsPreOrder())
	// fmt.Println(tree.BfsPostOrder())
	// fmt.Println(tree.BfsInorder())

	// ---------------------------------------------------Heap-------------------------------------------------------------------

	// maxHeap := &heep.MaxHeap{}
	// values := []int{10, 23, 45, 73, 86, 19, 36, 5}
	// for _, val := range values {
	// 	maxHeap.Insert(val)
	// }
	// fmt.Println(maxHeap)
	// fmt.Println("--------------")
	// for i := 0; i <= 5; i++ {
	// 	maxHeap.Extract()
	// 	fmt.Println(maxHeap)
	// }

	// minHeap := &heep.MinHeap{}
	// values := []int{10, 23, 45, 73, 86, 19, 36, 5}
	// for _, val := range values {
	// 	minHeap.Insert(val)
	// }
	// fmt.Println(minHeap)
	// fmt.Println(minHeap.Extract())
	// fmt.Println(minHeap)

	// ---------------------------------------------------Graph-------------------------------------------------------------------
	graph := graph.Graph{}
	for i := 0; i <= 5; i++ {
		graph.AddVertex(i)
	}

	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(3, 4)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)
	graph.AddEdge(5, 0)
	graph.Print()
}
