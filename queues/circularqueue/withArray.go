package circularqueue

import "fmt"

type Queue struct {
	Queue       []int
	front, rear int
}

func Init(cap int) *Queue {
	return &Queue{
		Queue: make([]int, cap),
		front: -1,
		rear:  -1,
	}
}

func (q *Queue) Enqueue(val int) {
	//check the queue is empty.if empty adding the value
	if q.front == -1 && q.rear == -1 {
		q.front = 0
		q.rear = 0
		q.Queue[q.rear] = val
		return
	}
	//check the queue is full
	if (q.rear+1)%len(q.Queue) == q.front {
		fmt.Println("full")
		return
	}
	//adding the value
	q.rear = (q.rear + 1) % len(q.Queue)
	q.Queue[q.rear] = val
}

func (q *Queue) Dequeue() {
	if q.front == -1 && q.rear == -1 {
		fmt.Println("queue is empty")
	} else if q.front == q.rear {
		fmt.Println("the deque element", q.Queue[q.front])
		q.front = -1
		q.rear = -1
	} else {
		fmt.Println("The dequeued element is", q.Queue[q.front])
		q.front = (q.front + 1) % len(q.Queue)
	}

}
