package ds

type Queue struct {
	items []int
}

// Enqueue adds a value to the queue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue removes the element from the queue
func (q *Queue) Dequeue() int {
	removed_value := q.items[0]
	q.items = q.items[1:]
	return removed_value
}
