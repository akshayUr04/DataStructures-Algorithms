package queue

import "fmt"

type ArrQueue struct {
	data              []int
	front, rear, size int
}

func Init(cap int) *ArrQueue {
	return &ArrQueue{
		data:  make([]int, cap),
		front: 0,
		rear:  0,
		size:  0,
	}
}

func (q *ArrQueue) Enqueue(val int) {
	if q.rear == len(q.data) {
		fmt.Println("queue is full")
		return
	}
	q.data[q.rear] = val
	q.rear++
	q.size++
}

func (q *ArrQueue) Dequeue() (int, bool) {
	if q.front == q.rear {
		fmt.Println("Queue is empty")
		return 0, false
	}
	item := q.data[q.front]
	q.front++
	return item, true
}

func (q *ArrQueue) Display() {
	if q.front == q.rear {
		fmt.Println("queue is empty")
		return
	}
	for i := q.front; i < q.rear; i++ {
		fmt.Println(q.data[i])
	}
}
