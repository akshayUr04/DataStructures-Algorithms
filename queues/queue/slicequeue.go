package queue

type Queue struct {
	items []int
}

//enqueue

func (q *Queue) SliceEnque(val int) {
	q.items = append(q.items, val)
}

// dequeue
func (q *Queue) SliceDequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}
